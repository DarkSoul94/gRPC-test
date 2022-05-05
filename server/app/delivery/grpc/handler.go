package grpc

import (
	"context"
	"fmt"

	pb "github/DarkSoul94/gRPC-test/test-api"
)

type TestServerHandler struct {
	pb.UnimplementedTestServer
}

func NewTestServerHandler() *TestServerHandler {
	return &TestServerHandler{}
}

func (t *TestServerHandler) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	var out *pb.HelloResponse

	out = &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", in.Name),
	}

	return out, nil
}
