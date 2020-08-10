package implements

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"io/ioutil"
	"net/http"
	"regexp"
	"runtime"
	ps "rusprofilegrpcservice/proto"
	"rusprofilegrpcservice/server/cache"
	"strings"
)

type FirmInfoType int

const (
	Inn FirmInfoType = iota
	Kpp
	NameAndBoss
)

const URL  = "https://www.rusprofile.ru/search?query=%v"

type RusProfileParserServiceServer struct {
	ResponseCache  *cache.Cache
}

func (s *RusProfileParserServiceServer) FirmInfoGet(ctx context.Context,
	req *ps.FirmByINNRequest) (*ps.FirmInfoResponse, error) {
	defer recovery()

	var response *ps.FirmInfoResponse

	if resp, ok := s.ResponseCache.Get(req.Inn); ok {
		fmt.Printf("From cache info by key %v\n", req.Inn)
		response = resp.(*ps.FirmInfoResponse)
	} else {
		fmt.Printf("Online info for key %v\n", req.Inn)

		response = new(ps.FirmInfoResponse)

		resp, err := http.Get(fmt.Sprintf(URL, req.Inn))
		if err != nil {
			return response, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return response, fmt.Errorf(resp.Status)
		}

		var body []byte
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return response, err
		}

		var bodyToParse = string(body)

		maxWorkers := runtime.GOMAXPROCS(0)
		sem := semaphore.NewWeighted(int64(maxWorkers))
		errs, ctxGroup := errgroup.WithContext(ctx)

		if err := sem.Acquire(ctxGroup, 3); err != nil {
			return response, err
		}

		errs.Go(func() error {
			defer sem.Release(1)
			defer recovery()

			return ParseFirmInfo(response, bodyToParse, Inn)
		})

		errs.Go(func() error {
			defer sem.Release(1)
			defer recovery()

			return ParseFirmInfo(response, bodyToParse, Kpp)
		})

		errs.Go(func() error {
			defer sem.Release(1)
			defer recovery()

			return ParseFirmInfo(response, bodyToParse, NameAndBoss)
		})

		if err := errs.Wait(); err != nil {
			return response, err
		}

		s.ResponseCache.Set(req.Inn, response)
	}

	return response, nil
}

func ParseFirmInfo(response *ps.FirmInfoResponse, body string, firmInfo FirmInfoType) error {
	switch firmInfo {
	case Inn:
		response.Inn = FindByPattern(".*clip_inn.?>(.*)<", body)
	case Kpp:
		response.Kpp = FindByPattern(".*clip_kpp.?>(.*)<", body)
	case NameAndBoss:
		var result = FindByPattern("legalName.?>(.*)</", body)

		if result != "" {
			response.Name = strings.Replace(result, "&quot;", "\"", -1)
			response.Boss = FindByPattern(fmt.Sprintf("<meta name=\"keywords\" content=.*%v, (.*), ИНН ", result), body)
		}
	}

	return nil
}

func FindByPattern(pattern string, body string) string {
	re := regexp.MustCompile(pattern)
	sub := re.FindAllStringSubmatch(body, -1)
	for _, element := range sub {
		return element[1]
	}

	return ""
}

func recovery()  {
	if err := recover(); err != nil {
		fmt.Println("Recovered from err: ", err)
	}
}