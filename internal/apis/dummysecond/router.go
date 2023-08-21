package dummysecond

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	router.POST("/v1/dummy2", DummyControllerSecond)
	router.PUT("/v1/dummy2", DummyControllerSecond)
	router.GET("/v1/dummy2", DummyControllerSecond)
	router.DELETE("/v1/dummy2", DummyControllerSecond)
}
