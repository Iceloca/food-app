package grpc

import (
	"github.com/r1nb0/food-app/auth-svc/internal/grpc/auth"
	"github.com/r1nb0/food-app/auth-svc/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       string
}

func New(authService service.Auth, port string) *App {
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
		log.Fatalf("failed to listen: %v", err)
	}

	if err = a.gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
