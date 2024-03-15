package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

const (
	hostEnvName       = "HTTP_HOST"
	portServerEnvName = "HTTP_PORT"
	expiredEnvName    = "EXPIRED"
	secretKeyEnvName  = "SECRET_KEY"
)

type ServerConfig interface {
	HttpPort() string
	SecretKey() []byte
	Expired() time.Duration
}

type serverConfig struct {
	host      string
	port      string
	secretKey string
	expired   int64
}

func NewServerConfig() (*serverConfig, error) {
	host := os.Getenv(hostEnvName)
	if len(host) == 0 {
		return nil, errors.New("mongo username not found")
	}

	port := os.Getenv(portServerEnvName)

	if len(host) == 0 {
		return nil, errors.New("server port not found")
	}

	secretKey := os.Getenv(secretKeyEnvName)
	if len(secretKey) == 0 {
		return nil, errors.New("server secret key not found")
	}

	expired, err := strconv.ParseInt(os.Getenv(expiredEnvName), 10, 64)
	if err != nil {
		return nil, errors.New("failed parse expired")
	}
	return &serverConfig{
		host:      host,
		port:      port,
		secretKey: secretKey,
		expired:   expired,
	}, nil
}

func (cfg *serverConfig) HttpPort() string {
	return cfg.port
}

func (cfg *serverConfig) SecretKey() []byte {
	secretKeyByte := []byte(cfg.secretKey)
	return secretKeyByte
}

func (cfg *serverConfig) Expired() time.Duration {
	return time.Hour * time.Duration(cfg.expired)
}
