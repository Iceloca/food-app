package product

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/pkg/database"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/service"
	productv1 "github.com/r1nb0/protos/gen/go/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productServer struct {
	productService service.ProductService
	productv1.UnimplementedProductServiceServer
}

func Register(gRPCServer *grpc.Server, productService service.ProductService) {
	productv1.RegisterProductServiceServer(gRPCServer, &productServer{
		productService: productService,
	})
}

func (s *productServer) Create(
	ctx context.Context,
	req *productv1.CreateRequest,
) (*productv1.CreateResponse, error) {
	newProduct := models.NewProductCreateFromGRPC(req)
	id, err := s.productService.Create(ctx, newProduct)
	if err != nil {
		if errors.Is(err, database.ErrAlreadyExists) {
			return nil, status.Error(
				codes.AlreadyExists,
				"product with this name already exists",
			)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.CreateResponse{
		Id: id,
	}, nil
}

func (s *productServer) Update(
	ctx context.Context,
	req *productv1.Product,
) (*productv1.UpdateResponse, error) {
	updateProduct := models.NewProductFromGRPC(req)
	err := s.productService.Update(ctx, updateProduct)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		return &productv1.UpdateResponse{
			Success: false,
		}, status.Error(codes.Internal, err.Error())
	}

	return &productv1.UpdateResponse{
		Success: true,
	}, nil
}

func (s *productServer) GetByID(
	ctx context.Context,
	req *productv1.GetByIDRequest,
) (*productv1.Product, error) {
	product, err := s.productService.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return product.MapToGRPCProduct(), nil
}

func (s *productServer) Delete(
	ctx context.Context,
	req *productv1.DeleteRequest,
) (*productv1.DeleteResponse, error) {
	err := s.productService.Delete(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "product not found")
		}
		return &productv1.DeleteResponse{
			Success: false,
		}, status.Error(codes.Internal, err.Error())
	}

	return &productv1.DeleteResponse{
		Success: true,
	}, nil
}

func (s *productServer) GetAll(
	_ *productv1.GetAllRequest,
	stream grpc.ServerStreamingServer[productv1.Product],
) error {
	products, err := s.productService.GetAll(stream.Context())
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return status.Error(codes.NotFound, "products not found")
		}
		return status.Error(codes.Internal, err.Error())
	}

	for _, product := range products {
		if err = stream.Send(product.MapToGRPCProduct()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (s *productServer) GetByCategory(
	req *productv1.GetByCategoryRequest,
	stream grpc.ServerStreamingServer[productv1.Product],
) error {
	products, err := s.productService.GetByCategory(stream.Context(), req.CategoryId)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return status.Error(codes.NotFound, "products not found")
		}
		return status.Error(codes.Internal, err.Error())
	}

	for _, product := range products {
		if err = stream.Send(product.MapToGRPCProduct()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (s *productServer) GetDailyRecs(
	_ *productv1.GetDailyRecsRequest,
	stream grpc.ServerStreamingServer[productv1.Product],
) error {
	products, err := s.productService.GetDailyRecs(stream.Context())
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return status.Error(codes.NotFound, "daily recs not found")
		}
		return status.Error(codes.Internal, err.Error())
	}

	for _, product := range products {
		if err = stream.Send(product.MapToGRPCProduct()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}
	return nil
}
