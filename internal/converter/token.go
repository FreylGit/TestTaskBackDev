package converter

import (
	"github.com/FreylGit/TestTaskBackDev/internal/model"
	modelRepo "github.com/FreylGit/TestTaskBackDev/internal/repository/token/model"
)

func ToTokenFromService(rtoken model.RefreshToken) *modelRepo.RefreshToken {

	return &modelRepo.RefreshToken{
		Id:    rtoken.Id,
		Token: rtoken.Token,
		Exp:   rtoken.Exp,
	}
}

func ToTokenCreateFromService(rtoken model.RefreshToken) *modelRepo.RefreshTokenCreate {
	return &modelRepo.RefreshTokenCreate{
		Token: rtoken.Token,
		Exp:   rtoken.Exp,
	}
}
