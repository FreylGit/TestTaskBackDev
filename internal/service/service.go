package service

import (
	"context"
	"errors"
)

var (
	ErrGenerateToken = errors.New("error when generating token")
	ErrVerifyToken   = errors.New("error validating token")
	ErrTokenPairs    = errors.New("error access token is not suitable for refresh token")
)

type AuthService interface {
	Create(ctx context.Context, userId string) (string, string, error)
	Update(ctx context.Context, accessToken string, refreshToken string) (string, string, error)
}
