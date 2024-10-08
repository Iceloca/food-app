package cart

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/cart-svc/internal/domain/models"
	"github.com/r1nb0/food-app/cart-svc/internal/service"
	"github.com/r1nb0/food-app/pkg/database"
	cartv1 "github.com/r1nb0/protos/gen/go/cart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type cartServer struct {
	cartService service.CartService
	cartv1.UnimplementedCartServiceServer
}

func Register(gRPCServer *grpc.Server, cartService service.CartService) {
	cartv1.RegisterCartServiceServer(
		gRPCServer,
		&cartServer{cartService: cartService},
	)
}

func (s *cartServer) Create(
	ctx context.Context,
	req *cartv1.CreateRequest,
) (*cartv1.CreateResponse, error) {
	cart := models.NewCartCreateFromGRPC(req)

	id, err := s.cartService.Create(ctx, cart)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &cartv1.CreateResponse{
		Id: id,
	}, nil
}

func (s *cartServer) Delete(
	ctx context.Context,
	req *cartv1.DeleteRequest,
) (*cartv1.DeleteResponse, error) {
	if err := s.cartService.Delete(ctx, req.Id); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return &cartv1.DeleteResponse{
				Success: false,
			}, status.Error(codes.NotFound, "cart not found")
		}
		return &cartv1.DeleteResponse{
			Success: false,
		}, status.Error(codes.Internal, err.Error())
	}

	return &cartv1.DeleteResponse{
		Success: true,
	}, nil
}

func (s *cartServer) GetByID(
	ctx context.Context,
	req *cartv1.GetByIDRequest,
) (*cartv1.Cart, error) {
	cart, err := s.cartService.GetByID(ctx, req.Id)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "cart not found")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return cart.MapCartToGRPC(), nil
}

func (s *cartServer) AddItem(
	ctx context.Context,
	req *cartv1.AddItemRequest,
) (*cartv1.AddItemResponse, error) {
	item := models.NewItemFromGRPC(req.Item)
	if err := s.cartService.AddItem(ctx, req.BasketId, item); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return &cartv1.AddItemResponse{
				Success: false,
			}, status.Error(codes.NotFound, "cart not found")
		}
		return &cartv1.AddItemResponse{
			Success: false,
		}, status.Error(codes.Internal, err.Error())
	}

	return &cartv1.AddItemResponse{
		Success: true,
	}, nil
}

func (s *cartServer) DeleteItem(
	ctx context.Context,
	req *cartv1.DeleteItemRequest,
) (*cartv1.DeleteItemResponse, error) {
	if err := s.cartService.DeleteItem(ctx, req.BasketId, req.ItemId); err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return &cartv1.DeleteItemResponse{
				Success: false,
			}, status.Error(codes.NotFound, "cart not found")
		}
		return &cartv1.DeleteItemResponse{
			Success: false,
		}, status.Error(codes.Internal, err.Error())
	}

	return &cartv1.DeleteItemResponse{
		Success: true,
	}, nil
}
