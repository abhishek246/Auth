package resourcegroup

import (
	"net/http"

	"auth/pkg/contracts"

	"github.com/gin-gonic/gin"
)

func CreateResourceGroupController(c *gin.Context) {
	var resourceGroupCreateRequest contracts.ResourceGroupCreateRequest

	if err := c.ShouldBindJSON(&resourceGroupCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Only Admin access
	err := InsertResourceGroup(resourceGroupCreateRequest.GroupName, resourceGroupCreateRequest.GroupDescription)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resource Group created successfully"})
}

func UpdateResourceGroupController(c *gin.Context) {
	var resourceGroupUpdateRequest contracts.ResourceGroupUpdateRequest
	groupId := c.Param("resource_group_id")
	if err := c.ShouldBindJSON(&resourceGroupUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Only Admin access

	UpdateResourceGroup(groupId, resourceGroupUpdateRequest.GroupName, resourceGroupUpdateRequest.GroupDescription)

	c.JSON(http.StatusOK, gin.H{"message": "Resource Group updated successfully"})
}
