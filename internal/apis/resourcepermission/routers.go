package resourcepermission

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	router.POST("/v1/resourcepermission", CreateResourcePermissionController)
}
