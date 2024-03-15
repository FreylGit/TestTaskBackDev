package auth

import (
	"github.com/FreylGit/TestTaskBackDev/internal/config"
	"github.com/FreylGit/TestTaskBackDev/internal/repository"
	"github.com/FreylGit/TestTaskBackDev/internal/service"
)

var _ service.AuthService = (*serv)(nil)

type serv struct {
	tokenRepository repository.TokenRepository
	config          config.ServerConfig
}

func NewService(tokenRepository repository.TokenRepository, config config.ServerConfig) *serv {
	return &serv{tokenRepository: tokenRepository, config: config}
}
