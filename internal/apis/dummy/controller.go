package dummy

import (
	"auth/pkg/authenticator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const Resource = "dummy"

var ActionMap = map[string]string{
	"POST":   "CREATE",
	"PUT":    "UPDATE",
	"GET":    "READ",
	"DELETE": "DELETE",
}

func DummyController(c *gin.Context) {
	action := ActionMap[c.Request.Method]

	fmt.Println("UserID HEADER :", c.GetHeader("UserId"))
	fmt.Println("Token: ", c.GetHeader("Token"))
	if !authenticator.Authenticate(c.GetHeader("UserId"), c.GetHeader("Token"), Resource, action) {
		c.JSON(http.StatusForbidden, gin.H{"message": "You currently don't have access to the API please request admin provide access"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "You have access"})
}
