package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	RefreshId string `json:"refreshId"`
}

func GenerateAccessToken(userId string, refreshId string, exp time.Time, secretKey []byte) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Id:        userId,
			ExpiresAt: exp.Unix(),
		},
		RefreshId: refreshId,
	}
	//SHA512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(secretKey)
}

func GenerateRefreshToken(userId string, exp time.Time, secretKey []byte) (string, error) {
	claims := jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: exp.Unix(),
	}

	//SHA512
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(secretKey)
}

func VerifyToken(tokenStr string, secretKey []byte) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}

			return secretKey, nil
		},
	)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
