package resourcegroup

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.POST("/v1/resourcegroup", CreateResourceGroupController)
	router.PUT("/v1/resourcegroup/:resource_group_id", UpdateResourceGroupController)
	router.DELETE("/v1/resourcegroup", UpdateResourceGroupController)
	router.GET("/v1/resourcegroup", UpdateResourceGroupController)
}
