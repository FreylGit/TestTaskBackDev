package auth

import (
	"github.com/FreylGit/TestTaskBackDev/internal/repository"
	"github.com/FreylGit/TestTaskBackDev/internal/service"
	"net/http"
)

type AuthAPI struct {
	authService service.AuthService
}

func NewAuthAPI(authService service.AuthService) *AuthAPI {
	return &AuthAPI{authService: authService}
}

type response struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type request struct {
	Id string `json:"id"`
}

func handleError(w http.ResponseWriter, err error) {
	var statusCode int
	var errorMessage string

	switch err {
	case service.ErrGenerateToken, repository.ErrCreate, repository.ErrUpdate:
		statusCode = http.StatusInternalServerError
		errorMessage = "Internal server error"
	case service.ErrVerifyToken, repository.ErrConvert, repository.ErrNotFound:
		statusCode = http.StatusBadRequest
		errorMessage = "Invalid request"
	case service.ErrTokenPairs:
		statusCode = http.StatusUnauthorized
		errorMessage = "Unauthorized"
	default:
		statusCode = http.StatusInternalServerError
		errorMessage = "Something went wrong"
	}

	w.WriteHeader(statusCode)
	w.Write([]byte(errorMessage))
}
