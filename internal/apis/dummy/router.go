package dummy

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	router.POST("/v1/dummy", DummyController)
	router.PUT("/v1/dummy", DummyController)
	router.GET("/v1/dummy", DummyController)
	router.DELETE("/v1/dummy", DummyController)
}
