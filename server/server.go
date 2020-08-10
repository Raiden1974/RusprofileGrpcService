package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"log"
	"net"
	ps "rusprofilegrpcservice/proto"
	"rusprofilegrpcservice/server/cache"
	imp "rusprofilegrpcservice/server/implements"
	"time"
)

const cacheExpire = "10s"

var (
	// command-line options:
	// gRPC server endpoint
	grpcServer = flag.String("grpc-server",  "127.0.0.1:8080", "gRPC server")

	// cache expire duration
	cacheExpireDuration = flag.String("cache-expire",  cacheExpire, "cache expire duration (example 10s)")
)

func main() {
	duration, err := time.ParseDuration(*cacheExpireDuration)
	if err != nil {
		glog.Fatalf("Error parse respCache-expire param value=%v", *cacheExpireDuration)
		return
	}
	server := grpc.NewServer()

	service := &imp.RusProfileParserServiceServer{}

	respCache := cache.New(duration, duration/2.0)
	service.ResponseCache = respCache
	defer respCache.Close()

	ps.RegisterRusprofileParserServiceServer(server, service)

	fmt.Printf("%v %v",  *grpcServer, duration)
	listener, err := net.Listen("tcp", *grpcServer)
	if err != nil {
		log.Fatal("Unable to create GRPC listener:", err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
