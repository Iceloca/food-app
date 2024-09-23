package service

import "context"

type Auth interface {
	Login(ctx context.Context, email string, pass string) (string, error)
	Register(ctx context.Context, email string, pass string) (int64, error)
}
