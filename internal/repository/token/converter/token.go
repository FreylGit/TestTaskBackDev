package converter

import (
	"github.com/FreylGit/TestTaskBackDev/internal/model"
	modelRepo "github.com/FreylGit/TestTaskBackDev/internal/repository/token/model"
)

func ToTokenFromRepo(rtoken modelRepo.RefreshToken) *model.RefreshToken {
	return &model.RefreshToken{
		Id:    rtoken.Id,
		Token: rtoken.Token,
		Exp:   rtoken.Exp,
	}
}
