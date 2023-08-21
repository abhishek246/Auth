// cmd/Auth/main.go
package main

import (
	"auth/internal/apis"
	"auth/internal/db"
	redisclient "auth/internal/redisclient"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.InitDB()
	redisclient.InitRedisClient()

	apis.SetupRoutes(router)
	router.Run(":8080")
}
