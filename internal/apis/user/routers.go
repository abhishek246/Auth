package user

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	router.POST("/v1/user/register", RegisterUserController)
	router.POST("/v1/user/changepassword", ChangePasswordController)
	router.POST("/v1/user/:userId/permission", AddPermissionController)
	router.PUT("/v1/user/:userId/permission", RevokePermissionController)
}
