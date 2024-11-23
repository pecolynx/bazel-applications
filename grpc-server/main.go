package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	v1 "github.com/pecolynx/bazel-applications/proto/hello/v1"
	"google.golang.org/grpc"
)

// GreeterServer is the server API for Greeter service.
type GreeterServer struct {
	v1.UnimplementedGreeterServer
}

// NewGreeterServer creates a new GreeterServer.
func NewGreeterServer() *GreeterServer {
	return &GreeterServer{}
}

// SayHello implements GreeterServer
func (s *GreeterServer) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{
		Message: "Hello, " + in.Name,
	}, nil
}
func main() {
	request := v1.HelloRequest{}
	fmt.Println(request)

	server := grpc.NewServer()
	v1.RegisterGreeterServer(server, NewGreeterServer())

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen: %v\n", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	go func() {
		if err := server.Serve(listener); err != nil {
			fmt.Fprintf(os.Stderr, "failed to serve: %v\n", err)
		}
	}()
	<-ctx.Done()
	server.GracefulStop()

}
