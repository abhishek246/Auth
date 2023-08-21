// internal/login/routers.go
package login

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.POST("/v1/login", LoginController)
}
