package client

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
	categoryv1 "github.com/r1nb0/protos/gen/go/category"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"strconv"
)

type CategoryClient struct {
	Client categoryv1.CategoryServiceClient
}

func NewCategoryClient(cfg *config.Config) *CategoryClient {
	conn, err := grpc.NewClient(
		cfg.ProductServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("error creating gRPC client: %s", err.Error())
	}

	client := categoryv1.NewCategoryServiceClient(conn)

	return &CategoryClient{
		Client: client,
	}
}

func (c *CategoryClient) CreateCategory(ctx *gin.Context) {
	var req categoryv1.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	out, err := c.Client.Create(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, out)
}

func (c *CategoryClient) UpdateCategory(ctx *gin.Context) {
	var req categoryv1.Category
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	out, err := c.Client.Update(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	ctx.JSON(http.StatusOK, out)
}

func (c *CategoryClient) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	out, err := c.Client.GetByID(ctx, &categoryv1.GetByIDRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, out)
}

func (c *CategoryClient) GetAll(ctx *gin.Context) {
	stream, err := c.Client.GetAll(ctx, &categoryv1.GetAllRequest{})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	var categories []*categoryv1.Category
	for {
		category, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			ctx.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()},
			)
			return
		}
		categories = append(categories, category)
	}

	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryClient) DeleteCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	out, err := c.Client.Delete(ctx, &categoryv1.DeleteRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, out)
}
