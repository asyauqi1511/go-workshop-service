package user

import (
	"context"

	"github.com/asyauqi1511/go-workshop-service/internal/entity"
)

func (uc Usecase) RegisterUser(ctx context.Context, param entity.UserRegisterParam) error {
	return uc.userModule.RegisterUser(ctx, param)
}
