package grpc

import (
	"github.com/r1nb0/food-app/cart-svc/internal/grpc/cart"
	"github.com/r1nb0/food-app/cart-svc/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       string
}

func New(cartService service.CartService, port string) *App {
	gRPCServer := grpc.NewServer()

	cart.Register(gRPCServer, cartService)

	return &App{
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() {
	lis, err := net.Listen("tcp", ":"+a.port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := a.gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
