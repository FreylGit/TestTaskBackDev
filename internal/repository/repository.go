package repository

import (
	"context"
	"errors"
	"github.com/FreylGit/TestTaskBackDev/internal/model"
)

var (
	ErrUpdate   = errors.New("failed to update token")
	ErrCreate   = errors.New("failed to save token")
	ErrConvert  = errors.New("failed to convert data")
	ErrNotFound = errors.New("could not find entry")
)

type TokenRepository interface {
	Create(ctx context.Context, token *model.RefreshToken) error
	Get(ctx context.Context, token string) (*model.RefreshToken, error)
	Update(ctx context.Context, token *model.RefreshToken) error
	Delete(ctx context.Context, token string) error
}
