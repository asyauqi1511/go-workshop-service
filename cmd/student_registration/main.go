package main

import (
	"github.com/asyauqi1511/go-workshop-service/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	// Load config.

	// Initialize DB Connection.
	db, err := database.ConnectDB()

	r := gin.Default()
	r.Run()
}
