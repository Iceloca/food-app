package auth

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/auth-svc/internal/service"
	"github.com/r1nb0/food-app/pkg/database"
	authv1 "github.com/r1nb0/protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServer struct {
	authv1.UnimplementedAuthServer
	auth service.Auth
}

func Register(gRPCServer *grpc.Server, auth service.Auth) {
	authv1.RegisterAuthServer(gRPCServer, &authServer{auth: auth})
}

func (s *authServer) Login(
	ctx context.Context,
	req *authv1.LoginRequest,
) (*authv1.LoginResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}

	token, err := s.auth.Login(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid password or email")
		}
		return nil, status.Error(codes.Internal, "failed to login")
	}
	return &authv1.LoginResponse{Token: token}, nil
}

func (s *authServer) Register(
	ctx context.Context,
	req *authv1.RegisterRequest,
) (*authv1.RegisterResponse, error) {
	if req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "email is required")
	}
	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "password is required")
	}
	uid, err := s.auth.Register(ctx, req.Email, req.Password)
	if err != nil {
		if errors.Is(err, database.ErrAlreadyExists) {
			return nil, status.Error(codes.InvalidArgument, "user already exists")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authv1.RegisterResponse{UserId: uid}, nil
}
