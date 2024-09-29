package product

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/repository"
	"github.com/r1nb0/food-app/product-svc/internal/service"
	productv1 "github.com/r1nb0/protos/gen/go/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productServer struct {
	productv1.UnimplementedProductServiceServer
	productService service.ProductService
}

func Register(gRPCServer *grpc.Server, productService service.ProductService) {
	productv1.RegisterProductServiceServer(gRPCServer, &productServer{
		productService: productService,
	})
}

func (s *productServer) CreateProduct(
	ctx context.Context,
	req *productv1.CreateRequest,
) (*productv1.CreateResponse, error) {
	id, err := s.productService.CreateProduct(ctx, models.NewCreateProductDTOFromGRPC(req))
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.CreateResponse{Id: id}, nil
}

func (s *productServer) DeleteProduct(
	ctx context.Context,
	req *productv1.DeleteRequest,
) (*productv1.DeleteResponse, error) {
	success, err := s.productService.DeleteProduct(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.DeleteResponse{Success: success}, nil
}

func (s *productServer) GetProductByID(
	ctx context.Context,
	req *productv1.GetByIDRequest,
) (*productv1.Product, error) {
	product, err := s.productService.GetProductByID(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return product.MapToGRPCProduct(), nil
}

func (s *productServer) GetAllProducts(
	_ *productv1.GetAllRequest,
	stream grpc.ServerStreamingServer[productv1.Product],
) error {
	products, err := s.productService.GetAllProducts(stream.Context())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}

	for _, product := range products {
		if err := stream.Send(product.MapToGRPCProduct()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (s *productServer) GetProductsByCategory(
	req *productv1.GetByCategoryRequest,
	stream grpc.ServerStreamingServer[productv1.Product],
) error {
	products, err := s.productService.GetProductsByCategory(stream.Context(), req.GetCategoryName())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}

	for _, product := range products {
		if err := stream.Send(product.MapToGRPCProduct()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (s *productServer) GetAllCategories(
	_ *productv1.GetAllCategoriesRequest,
	stream grpc.ServerStreamingServer[productv1.Category],
) error {
	categories, err := s.productService.GetAllCategories(stream.Context())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}

	for _, category := range categories {
		if err := stream.Send(category.MapToGRPCCategory()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (s *productServer) UpdateProduct(
	ctx context.Context,
	req *productv1.UpdateRequest,
) (*productv1.UpdateResponse, error) {
	success, err := s.productService.UpdateProduct(ctx, models.NewUpdateProductDTOFromGRPC(req))
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &productv1.UpdateResponse{Success: success}, nil
}

func (s *productServer) GetBestProducts(
	_ *productv1.GetBestProductsRequest,
	stream grpc.ServerStreamingServer[productv1.Product],
) error {
	products, err := s.productService.GetBestProducts(stream.Context())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return status.Error(codes.NotFound, err.Error())
		}
		return status.Error(codes.Internal, err.Error())
	}

	for _, product := range products {
		if err := stream.Send(product.MapToGRPCProduct()); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}
