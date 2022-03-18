package main

import (
	"github.com/ZainalAbiddin/ProjectAPI/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", handler.Nama)
	// soal 2, 5 query string
	r.GET("/product", handler.ProductHandler)
	// soal 3, 5 query param
	r.GET("/smartphone/:namaproduk/:merk/:jenis/:spesifikasi/:harga", handler.ProductSmartphone)
	r.Run()
}
