package middleware

import (
	"log"
	"net/http"

	"github.com/asyauqi1511/go-workshop-service/internal/constants"
	"github.com/asyauqi1511/go-workshop-service/internal/repo/user"
	"github.com/gin-gonic/gin"
)

func Auth(userRes user.Module) func(c *gin.Context) {
	return func(c *gin.Context) {
		username := c.GetHeader(constants.HeaderUsername)
		password := c.GetHeader(constants.HeaderPassword)

		user, err := userRes.GetAuth(c, username, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		if user.UserID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		c.Set(constants.UserIDTag, user.UserID)
		c.Set(constants.UserRoleTag, user.Role)

		c.Next()
	}
}

func HandleHTTP(handle func(c *gin.Context) (any, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, err := handle(c)
		if err != nil {
			log.Printf("Error: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
			c.Abort()
			return
		}

		if data != nil {
			c.JSON(http.StatusOK, gin.H{"data": data})
		} else {
			c.JSON(http.StatusOK, nil)

		}
		c.Next()
	}
}
