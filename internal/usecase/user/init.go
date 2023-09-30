package user

import "github.com/asyauqi1511/go-workshop-service/internal/repo/user"

type Usecase struct {
	userModule user.Module
}

func New(userModule user.Module) Usecase {
	return Usecase{
		userModule: userModule,
	}
}
