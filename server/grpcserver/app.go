package grpcserver

import (
	"net"

	grpcAPP "github/DarkSoul94/gRPC-test/server/app/delivery/grpc"
	pb "github/DarkSoul94/gRPC-test/test-api"

	"google.golang.org/grpc"
)

type Handlers struct {
	TestHandler pb.TestServer
}

type App struct {
	Handlers

	grpcServer *grpc.Server
}

func NewApp() *App {
	return &App{
		grpcServer: grpc.NewServer(),
		Handlers: Handlers{
			TestHandler: grpcAPP.NewTestServerHandler(),
		},
	}
}

func (a *App) Run(port string) {
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	pb.RegisterTestServer(a.grpcServer, a.TestHandler)
	a.grpcServer.Serve(l)
}

func (a *App) Stop() {
	a.grpcServer.Stop()
}
