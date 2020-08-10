package main

import (
	"context"
	"log"
	"net"
	ps "rusprofilegrpcservice/proto"
	imp "rusprofilegrpcservice/server/implements"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"rusprofilegrpcservice/server/cache"
)

func TestImplementation_GRPCRouting_FirmInfoGet_ShouldOK(t *testing.T) {
	var request = new(ps.FirmByINNRequest)
	request.Inn = "5017096885"

	ctx := context.TODO()
	srv, listener := startGRPCServer()
	defer func() {time.Sleep(10 * time.Millisecond)}()
	defer srv.Stop()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(getBufDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := ps.NewRusprofileParserServiceClient(conn)

	resp, err := client.FirmInfoGet(ctx, request)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Inn != request.Inn {
		t.Fatalf("expected inn=%v, got %v", request.Inn, resp.Inn)
	}
}

func startGRPCServer() (*grpc.Server, *bufconn.Listener) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)
	server := grpc.NewServer()
	service := &imp.RusProfileParserServiceServer{}

	duration, _ := time.ParseDuration("5s")
	respCache := cache.New(duration, duration/2.0)
	service.ResponseCache = respCache
	defer respCache.Close()

	ps.RegisterRusprofileParserServiceServer(server, service)

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()
	return server, listener
}

func getBufDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, url string) (net.Conn, error) {
		return listener.Dial()
	}
}