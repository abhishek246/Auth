package user

import (
	"auth/pkg/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserController(c *gin.Context) {
	var registerUserData contracts.RegisterUserRequest

	if err := c.ShouldBindJSON(&registerUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := RegisterUser(registerUserData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error while registering data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func ChangePasswordController(c *gin.Context) {
	var changePasswordData contracts.ChangePassword

	if err := c.ShouldBindJSON(&changePasswordData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := ChangePassword(changePasswordData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error while registering data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func AddPermissionController(c *gin.Context) {
	userId := c.Param("userId")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Admin can.
	var addPermissionRequest contracts.UserPermissionRequest
	if err := c.ShouldBindJSON(&addPermissionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := AddPermission(userId, addPermissionRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to add permission"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Permission added successfully"})
}

func RevokePermissionController(c *gin.Context) {
	userId := c.Param("userId")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var addPermissionRequest contracts.UserPermissionRequest
	if err := c.ShouldBindJSON(&addPermissionRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := RevokePermission(userId, addPermissionRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to add permission"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Permission revoked successfully"})
}
