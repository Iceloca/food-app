package auth

import (
	"context"
	"errors"
	"github.com/r1nb0/food-app/auth-svc/internal/lib/jwt"
	"github.com/r1nb0/food-app/auth-svc/internal/repository"
	"github.com/r1nb0/food-app/auth-svc/internal/service"
	"github.com/r1nb0/food-app/pkg/database"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type authService struct {
	userRepository repository.UserRepository
	logger         *slog.Logger
	tokenTTL       time.Duration
}

func NewAuthService(
	userRepository repository.UserRepository,
	logger *slog.Logger,
	tokenTTL time.Duration,
) service.Auth {
	return &authService{
		userRepository: userRepository,
		logger:         logger,
		tokenTTL:       tokenTTL,
	}
}

func (s *authService) Login(ctx context.Context, email string, pass string) (string, error) {
	const op = "authService.Login"

	log := s.logger.With(
		slog.String("operation", op),
		slog.String("email", email),
	)

	log.Info("attempting to login user")

	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			log.Warn("user not found",
				slog.String("error", err.Error()),
			)
		} else {
			log.Error("failed to get user",
				slog.String("error", err.Error()),
			)
		}
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword(user.PassHash, []byte(pass)); err != nil {
		log.Info(
			"invalid credentials",
			slog.String("error", err.Error()),
		)
		return "", service.ErrInvalidCredentials
	}

	token, err := jwt.NewToken(user, s.tokenTTL)
	if err != nil {
		log.Error("failed to generate token",
			slog.String("error", err.Error()),
		)
		return "", err
	}

	return token, nil
}

func (s *authService) Register(ctx context.Context, email string, pass string) (int64, error) {
	const op = "authService.Register"

	log := s.logger.With(
		slog.String("operation", op),
		slog.String("email", email),
	)

	log.Info("attempting to register user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Error(
			"failed to generate password hash",
			slog.String("error", err.Error()),
		)
		return 0, err
	}

	id, err := s.userRepository.SaveUser(ctx, email, passHash)
	if err != nil {
		log.Error(
			"failed to register user",
			slog.String("error", err.Error()),
		)
		return 0, err
	}

	return id, nil
}
