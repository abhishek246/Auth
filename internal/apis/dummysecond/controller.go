package dummysecond

import (
	"auth/pkg/authenticator"
	"net/http"

	"github.com/gin-gonic/gin"
)

const Resource = "dummy2"

var ActionMap = map[string]string{
	"POST":   "CREATE",
	"PUT":    "UPDATE",
	"GET":    "READ",
	"DELETE": "DELETE",
}

func DummyControllerSecond(c *gin.Context) {
	action := ActionMap[c.Request.Method]
	if !authenticator.Authenticate(c.GetHeader("UserId"), c.GetHeader("Token"), Resource, action) {
		c.JSON(http.StatusForbidden, gin.H{"message": "You currently don't have access to the API please request admin provide access"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "you have access"})
}
