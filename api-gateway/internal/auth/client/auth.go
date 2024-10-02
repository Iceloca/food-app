package client

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
	authv1 "github.com/r1nb0/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

type AuthClient struct {
	Client authv1.AuthClient
}

func NewAuthClient(cfg *config.Config) *AuthClient {
	conn, err := grpc.NewClient(
		cfg.AuthServiceURL,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		))
	if err != nil {
		log.Fatalf("error creating gRPC client: %s", err.Error())
	}

	client := authv1.NewAuthClient(conn)

	return &AuthClient{
		Client: client,
	}
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *AuthClient) Login(ctx *gin.Context) {
	var req AuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	token, err := c.Client.Login(ctx, &authv1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (c *AuthClient) Register(ctx *gin.Context) {
	var req AuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}

	fmt.Println(req)

	uid, err := c.Client.Register(ctx, &authv1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	ctx.JSON(http.StatusOK, uid)
}
