// internal/login/controller.go
package login

import (
	"context"
	"net/http"
	"time"

	"auth/internal/apis/user"
	"auth/internal/redisclient"
	"auth/pkg/contracts"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LoginController(c *gin.Context) {
	var loginData contracts.LoginRequest

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userId, status := AuthenticateLogin(loginData.EmailID, loginData.Password, loginData.ClientId)
	if !status {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Generate Token and store it in redis.
	// Get all the permission used by the customer and cache it redis

	token := generateUUID()
	ctx := context.Background()

	err := redisclient.RedisClient.Set(ctx, userId, token, 600*time.Second).Err()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication Successfully but permission failed"})
		return
	}
	err = user.CacheUserPermission(token, userId)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Error while handling error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "user_id": userId})
}

func generateUUID() string {
	newUUID := uuid.New()
	return newUUID.String()
}
