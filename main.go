package main

import (
	"github.com/ZainalAbiddin/ProjectAPI/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", handler.Nama)
	r.Run()
}
