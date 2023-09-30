package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) GetRegistrantDetail(c *gin.Context) (interface{}, error) {

	registrantId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("id not valid")
	}

	data, err := h.userUsecase.GetRegistrantData(c, registrantId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
