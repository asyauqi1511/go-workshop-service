package main

import (
	"fmt"

	"github.com/asyauqi1511/go-workshop-service/internal/config"
	"github.com/asyauqi1511/go-workshop-service/internal/handler"
	"github.com/asyauqi1511/go-workshop-service/internal/middleware"
	"github.com/asyauqi1511/go-workshop-service/internal/pkg/database"
	userRepo "github.com/asyauqi1511/go-workshop-service/internal/repo/user"
	userUc "github.com/asyauqi1511/go-workshop-service/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load config.
	appConfig, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed load config %v", err))
	}
	fmt.Println(appConfig)

	// Initialize DB Connection.
	db, err := database.ConnectDB(appConfig.DB)
	if err != nil {
		panic(fmt.Sprintf("Failed connect database %v", err))
	}

	// Initialize Redis.
	// redis, err := redis.ConnectRedis(appConfig.Redis)
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed connect redis %v", err))
	// }

	// Iniatialize Repo.
	userRes := userRepo.New(db, nil)

	// Initialize Usecase.
	userUsecase := userUc.New(userRes)

	// Initialize Handler.
	h := handler.New(userUsecase)

	r := gin.Default()

	// User Auth Group
	userGroup := r.Group("/user", middleware.Auth(userRes))
	userGroup.GET("/:id", middleware.HandleHTTP(h.GetRegistrantDetail))

	// Without Auth
	r.POST("/register", middleware.HandleHTTP(h.RegisterUser))

	r.Run(":8080")
}
