package routes

import (
	"github.com/gin-gonic/gin"
	authv1 "github.com/r1nb0/protos/gen/go/auth"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c authv1.AuthClient) {
	var req AuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
	}

	token, err := c.Login(ctx, &authv1.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(ctx *gin.Context, c authv1.AuthClient) {
	var req AuthRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
	}

	uid, err := c.Register(ctx, &authv1.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	ctx.JSON(http.StatusOK, gin.H{"uid": uid})
}
