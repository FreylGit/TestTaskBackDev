package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	usernameEnvName  = "USERNAME"
	passwordEnvName  = "PASSWORD"
	portMongoEnvName = "MONGO_PORT"
	hostMongoEnvName = "MONGO_HOST"
)

type MongoConfig interface {
	DSN() string
}
type mongoConfig struct {
	username string
	password string
	port     string
	host     string
}

func NewMongoConfig() (*mongoConfig, error) {
	username := os.Getenv(usernameEnvName)
	if len(username) == 0 {
		return nil, errors.New("mongo username not found")
	}

	password := os.Getenv(passwordEnvName)
	if len(password) == 0 {
		return nil, errors.New("mongo password not found")
	}

	port := os.Getenv(portMongoEnvName)

	if len(port) == 0 {
		return nil, errors.New("mongo port not found")
	}
	host := os.Getenv(hostMongoEnvName)

	if len(port) == 0 {
		return nil, errors.New("mongo host not found")
	}
	return &mongoConfig{
		username: username,
		password: password,
		port:     port,
		host:     host,
	}, nil
}

func (cfg *mongoConfig) DSN() string {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		cfg.username,
		cfg.password,
		cfg.host,
		cfg.port)

	return mongoURI
}
