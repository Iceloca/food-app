package auth

import (
	"github.com/r1nb0/food-app/api-gateway/internal/config"
	authv1 "github.com/r1nb0/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ServiceClient struct {
	Client authv1.AuthClient
}

func InitServiceClient(cfg *config.Config) authv1.AuthClient {
	conn, err := grpc.NewClient(
		cfg.AuthServiceURL,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		))
	if err != nil {
		log.Fatal(err)
	}

	return authv1.NewAuthClient(conn)
}
