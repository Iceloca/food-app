package app

import (
	"github.com/r1nb0/food-app/product-svc/internal/grpc/product"
	"github.com/r1nb0/food-app/product-svc/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
	port       string
}

func New(productService service.ProductService, port string) *App {
	gRPCServer := grpc.NewServer()

	product.Register(gRPCServer, productService)

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
