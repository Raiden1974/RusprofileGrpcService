package main

import (
	gw "RusprofileGrpcService/proto"
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "127.0.0.1:8080", "gRPC server endpoint")
	// listener port
	endpointPort = flag.Int("endpoint-port",  8082, "endpoint port")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	fmt.Println("Test direct request to GRPC...")
	firms := []string {"5017096885","6154010465","5190400349","2900000134","2508001431","2531001535","2322002888","2315007476","5017096885","6162002919"}
	i := 0

	for _, inn := range firms {
		go GetInfoFromService(inn)
		if i == 2 {
			time.Sleep(2 * time.Second)
		}
		i++
	}

	time.Sleep(1 * time.Second)

	if err := RunEndpoint(); err != nil {
		glog.Fatal(err)
	}
}

func GetInfoFromService(inn string) {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(*grpcServerEndpoint, opts...)

	if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := gw.NewRusprofileParserServiceClient(conn)
	request := &gw.FirmByINNRequest{
		Inn: inn,
	}

	var response, errResp = client.FirmInfoGet(context.Background(), request)
	if errResp != nil {
		glog.Warning("Error %v\n", errResp)
	} else {
		if response.Name != "" {
			fmt.Printf("\n%v\n%v\n%v\n%v\n", response.Name, response.Inn, response.Kpp, response.Boss)
		} else {
			fmt.Printf("\nNot found info for value = %v\n", inn)
		}
	}
}

func RunEndpoint() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	fmt.Printf("\nRegister gRPC server endpoint. Make sure the gRPC server is running properly and accessible.\n\n")
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterRusprofileParserServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	addr := fmt.Sprintf(":%d", *endpointPort)
	return http.ListenAndServe(addr, mux)
}