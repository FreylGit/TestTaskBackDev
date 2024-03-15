package auth

import (
	"context"
	"github.com/FreylGit/TestTaskBackDev/internal/service"
	"github.com/FreylGit/TestTaskBackDev/internal/utils"
	"time"
)

func (s *serv) Update(ctx context.Context, accessToken string, refreshToken string) (string, string, error) {
	//Верефицируем оба токена
	rtokenModel, err := s.tokenRepository.Get(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}
	claims, err := utils.VerifyToken(accessToken, s.config.SecretKey())
	if err != nil {
		return "", "", service.ErrVerifyToken
	}
	if rtokenModel.Id != claims.RefreshId {
		return "", "", service.ErrTokenPairs
	}

	//Выдаем новые
	exp := time.Now().Add(s.config.Expired())
	rtokenNew, err := utils.GenerateRefreshToken(claims.Id, exp.Add(s.config.Expired()*24), s.config.SecretKey())
	rtokenModel.Token = rtokenNew
	err = s.tokenRepository.Update(ctx, rtokenModel)
	if err != nil {
		return "", "", err
	}
	atokenNew, err := utils.GenerateAccessToken(claims.Id, claims.RefreshId, exp, s.config.SecretKey())
	if err != nil {
		return "", "", service.ErrGenerateToken
	}

	return atokenNew, rtokenNew, nil
}
