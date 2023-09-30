package handler

import "github.com/asyauqi1511/go-workshop-service/internal/usecase/user"

type Handler struct {
	userUsecase user.Usecase
}

func New(userUsecase user.Usecase) Handler {
	return Handler{
		userUsecase: userUsecase,
	}
}
