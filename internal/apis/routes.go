package apis

import (
	"auth/internal/apis/dummy"
	"auth/internal/apis/dummysecond"
	"auth/internal/apis/login"
	"auth/internal/apis/resourcegroup"
	"auth/internal/apis/resourcepermission"
	"auth/internal/apis/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	login.SetupRouter(router)
	user.SetupRouter(router)
	resourcegroup.SetupRouter(router)
	resourcepermission.SetupRouter(router)
	dummy.SetupRouter(router)
	dummysecond.SetupRouter(router)
}
