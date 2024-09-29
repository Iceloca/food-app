package category

import (
	"context"
	"github.com/r1nb0/food-app/product-svc/internal/domain/models"
	"github.com/r1nb0/food-app/product-svc/internal/service"
	categoryv1 "github.com/r1nb0/protos/gen/go/category"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type categoryServer struct {
	categoryService service.CategoryService
	categoryv1.UnimplementedCategoryServiceServer
}

func Register(gRPCServer *grpc.Server, categoryService service.CategoryService) {
	categoryv1.RegisterCategoryServiceServer(gRPCServer, &categoryServer{
		categoryService: categoryService,
	})
}

func (s *categoryServer) Create(
	ctx context.Context, req *categoryv1.CreateRequest,
) (*categoryv1.CreateResponse, error) {
	id, err := s.categoryService.Create(ctx, models.CategoryCreate{
		Name:     req.GetName(),
		ImageURL: req.GetImageURL(),
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &categoryv1.CreateResponse{
		Id: id,
	}, nil
}

func (s *categoryServer) Update(
	ctx context.Context, req *categoryv1.Category,
) (*categoryv1.UpdateResponse, error) {
	if err := s.categoryService.Update(ctx, models.Category{
		ID:       req.GetId(),
		Name:     req.GetName(),
		ImageURL: req.GetImageURL(),
	}); err != nil {
		return &categoryv1.UpdateResponse{
			Success: false,
		}, status.Error(codes.Internal, err.Error())
	}

	return &categoryv1.UpdateResponse{
		Success: true,
	}, nil
}

func (s *categoryServer) Delete(
	ctx context.Context,
	req *categoryv1.DeleteRequest,
) (*categoryv1.DeleteResponse, error) {
	if err := s.categoryService.Delete(ctx, req.GetId()); err != nil {
		return &categoryv1.DeleteResponse{
			Success: false,
		}, status.Error(codes.Internal, err.Error())
	}

	return &categoryv1.DeleteResponse{
		Success: true,
	}, nil
}

func (s *categoryServer) GetAll(
	_ *categoryv1.GetAllRequest,
	stream grpc.ServerStreamingServer[categoryv1.Category],
) error {
	categories, err := s.categoryService.GetAll(stream.Context())
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	for _, category := range categories {
		if err := stream.Send(&categoryv1.Category{
			Id:       category.ID,
			Name:     category.Name,
			ImageURL: category.ImageURL,
		}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
	}

	return nil
}

func (s *categoryServer) GetByID(
	ctx context.Context,
	req *categoryv1.GetByIDRequest,
) (*categoryv1.Category, error) {
	category, err := s.categoryService.GetByID(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &categoryv1.Category{
		Id:       category.ID,
		Name:     category.Name,
		ImageURL: category.ImageURL,
	}, nil
}
