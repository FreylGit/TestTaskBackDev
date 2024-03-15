package auth

import (
	"context"
	"github.com/FreylGit/TestTaskBackDev/internal/model"
	"github.com/FreylGit/TestTaskBackDev/internal/service"
	"github.com/FreylGit/TestTaskBackDev/internal/utils"
	"time"
)

func (s *serv) Create(ctx context.Context, userId string) (string, string, error) {
	exp := time.Now().Add(s.config.Expired())
	rtoken, err := utils.GenerateRefreshToken(userId, exp.Add(s.config.Expired()*24), s.config.SecretKey())
	if err != nil {
		return "", "", service.ErrGenerateToken
	}
	rtokenModel := &model.RefreshToken{
		Token: rtoken,
		Exp:   exp,
	}
	err = s.tokenRepository.Create(ctx, rtokenModel)
	if err != nil {
		return "", "", err
	}
	rtokenModel, err = s.tokenRepository.Get(ctx, rtoken)
	if err != nil {
		return "", "", err
	}

	atoken, err := utils.GenerateAccessToken(userId, rtokenModel.Id, exp, s.config.SecretKey())
	if err != nil {
		return "", "", service.ErrGenerateToken
	}

	return atoken, rtoken, nil
}
