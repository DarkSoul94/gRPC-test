package main

import (
	"context"
	"fmt"
	pb "github/DarkSoul94/gRPC-test/test-api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewTestClient(conn)

	res, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Slava"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Message)
}
