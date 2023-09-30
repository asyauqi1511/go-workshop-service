package handler

import (
	"github.com/asyauqi1511/go-workshop-service/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h Handler) RegisterUser(c *gin.Context) (interface{}, error) {
	var (
		request entity.UserRegisterParam
	)

	err := c.ShouldBindJSON(&request)
	if err != nil {
		return nil, err
	}

	err = h.userUsecase.RegisterUser(c, request)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
