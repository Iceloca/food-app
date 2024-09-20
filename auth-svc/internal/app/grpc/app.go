package grpc

import (
	"github.com/r1nb0/food-app/auth-svc/internal/grpc/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       string
}

func New(authService auth.Auth, port string) *App {
	gRPCServer := grpc.NewServer()

	auth.Register(gRPCServer, authService)

	return &App{
		gRPCServer,
		port,
	}
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		log.Fatal(err)
	}

	if err = a.gRPCServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
