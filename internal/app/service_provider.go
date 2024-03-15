package app

import (
	"context"
	"github.com/FreylGit/TestTaskBackDev/internal/api/auth"
	"github.com/FreylGit/TestTaskBackDev/internal/client/db"
	"github.com/FreylGit/TestTaskBackDev/internal/config"
	"github.com/FreylGit/TestTaskBackDev/internal/repository"
	"github.com/FreylGit/TestTaskBackDev/internal/repository/token"
	"github.com/FreylGit/TestTaskBackDev/internal/service"
	aServ "github.com/FreylGit/TestTaskBackDev/internal/service/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type serviceProvider struct {
	serverConfig    config.ServerConfig
	mongoConfig     config.MongoConfig
	client          *db.Client
	tokenRepository repository.TokenRepository
	authService     service.AuthService
	authAPI         *auth.AuthAPI
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) ServerConfig() config.ServerConfig {
	if s.serverConfig == nil {
		cfg, err := config.NewServerConfig()
		if err != nil {
			log.Fatalf("Failed to get server config: %w", err.Error())
		}
		s.serverConfig = cfg
	}

	return s.serverConfig
}

func (s *serviceProvider) MongoConfig() config.MongoConfig {
	if s.serverConfig == nil {
		cfg, err := config.NewMongoConfig()
		if err != nil {
			log.Fatalf("Failed to get mongo config: %w", err.Error())
		}
		s.mongoConfig = cfg
	}

	return s.mongoConfig
}

func (s *serviceProvider) AuthAPI(ctx context.Context) *auth.AuthAPI {
	if s.authAPI == nil {
		s.authAPI = auth.NewAuthAPI(s.AuthService(ctx))
	}

	return s.authAPI
}

func (s *serviceProvider) Client(ctx context.Context) *db.Client {
	if s.client == nil {
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(s.MongoConfig().DSN()))
		if err != nil {
			log.Fatal("Failed connect to mongodb")
		}
		s.client = db.NewClient(client)
	}
	return s.client
}

func (s *serviceProvider) TokenRepository(ctx context.Context) repository.TokenRepository {
	if s.tokenRepository == nil {
		s.tokenRepository = token.NewRepository(s.Client(ctx))
	}

	return s.tokenRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = aServ.NewService(s.TokenRepository(ctx), s.ServerConfig())
	}

	return s.authService
}
