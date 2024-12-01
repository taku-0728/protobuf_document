package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	hellopb "protobuf_document/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type myServer struct {
	hellopb.UnimplementedHelloServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	port := 8080
	listner, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error occured err: %v", err)
	}

	s := grpc.NewServer()

	hellopb.RegisterHelloServiceServer(s, NewMyServer())

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(listner)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

func (m *myServer) GetHello(ctx context.Context, _ *emptypb.Empty) (*hellopb.GetHelloResponse, error) {
	return &hellopb.GetHelloResponse{
		Message: "Hello World",
	}, nil
}
