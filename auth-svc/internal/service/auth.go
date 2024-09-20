package service

import (
	"context"
	"github.com/r1nb0/food-app/auth-svc/internal/lib"
	"github.com/r1nb0/food-app/auth-svc/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	userRepository repository.UserRepository
	tokenTTL       time.Duration
}

func NewAuthService(userRepository repository.UserRepository, tokenTTL time.Duration) *AuthService {
	return &AuthService{userRepository: userRepository}
}

// TODO Logger
func (s *AuthService) Login(ctx context.Context, email string, pass string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword(user.PassHash, []byte(pass)); err != nil {
		return "", err
	}

	token, err := lib.NewToken(user, s.tokenTTL)
	if err != nil {
		return "", err
	}

	return token, nil
}

// TODO Logger
func (s *AuthService) Register(ctx context.Context, email string, pass string) (int64, error) {
	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	id, err := s.userRepository.SaveUser(ctx, email, passHash)
	if err != nil {
		return 0, err
	}

	return id, nil
}
