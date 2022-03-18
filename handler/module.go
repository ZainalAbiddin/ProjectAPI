package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// localhost:8080/user/suddin/
func Nama(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}
