package client

import "C"
import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/r1nb0/food-app/api-gateway/internal/config"
	productv1 "github.com/r1nb0/protos/gen/go/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ProductClient struct {
	Client productv1.ProductServiceClient
}

func NewProductClient(cfg *config.Config) *ProductClient {
	conn, err := grpc.NewClient(
		cfg.ProductServiceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatalf("error creating gRPC client: %s", err.Error())
	}

	client := productv1.NewProductServiceClient(conn)

	return &ProductClient{
		Client: client,
	}
}

func (c *ProductClient) CreateProduct(ctx *gin.Context) {
	var req productv1.CreateRequest
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

func (c *ProductClient) UpdateProduct(ctx *gin.Context) {
	var req productv1.Product
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	out, err := c.Client.Update(ctx, &req)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	ctx.JSON(http.StatusOK, out)
}

func (c *ProductClient) DeleteProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	out, err := c.Client.Delete(ctx, &productv1.DeleteRequest{
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

func (c *ProductClient) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
	}

	out, err := c.Client.GetByID(ctx, &productv1.GetByIDRequest{
		Id: int64(id),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	ctx.JSON(http.StatusOK, out)
}

func (c *ProductClient) GetByCategory(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	stream, err := c.Client.GetByCategory(ctx, &productv1.GetByCategoryRequest{
		CategoryId: int64(id),
	})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	products, err := getProductsFromStream(stream)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductClient) GetAll(ctx *gin.Context) {
	stream, err := c.Client.GetAll(ctx, &productv1.GetAllRequest{})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	products, err := getProductsFromStream(stream)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (c *ProductClient) GetDailyRecs(ctx *gin.Context) {
	stream, err := c.Client.GetDailyRecs(ctx, &productv1.GetDailyRecsRequest{})
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	products, err := getProductsFromStream(stream)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func getProductsFromStream(
	stream grpc.ServerStreamingClient[productv1.Product],
) ([]*productv1.Product, error) {
	var products []*productv1.Product
	for {
		product, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}
