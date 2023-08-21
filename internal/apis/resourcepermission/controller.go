package resourcepermission

import (
	"auth/pkg/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateResourcePermissionController(c *gin.Context) {
	var resourcePermissionCreateRequest contracts.ResourcePermissionCreateRequest

	if err := c.ShouldBindJSON(&resourcePermissionCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := InsertResourcePermission(resourcePermissionCreateRequest.PermissionName,
		resourcePermissionCreateRequest.PermissionAction,
		resourcePermissionCreateRequest.ResourceGroupName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to create permission"})
		return
	}
	// Generate Token and store it in redis.

	c.JSON(http.StatusOK, gin.H{"message": "Permission created successfully"})
}
