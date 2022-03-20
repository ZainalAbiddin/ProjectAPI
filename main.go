package main

import (
	"github.com/ZainalAbiddin/ProjectAPI/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/user/:name/*action", handler.Nama)
	// soal 1
	r.GET("/item", handler.ItemHandler)
	// soal 2
	r.GET("/id/:id/produk", handler.ItemhandlerID)
	// soal 2, 5 query string
	r.GET("/product", handler.ProductHandler)
	// soal 3, 5 query param
	r.GET("/smartphone/:namaproduk/:merk/:jenis/:spesifikasi/:harga", handler.ProductSmartphone)

	// di Postman -> ganti post
	// Body -> -raw -> JSON
	// input data, contoh =
	/*	{
		"Pertemuan-Satu" :"Pengenalan Dasar Pemrograman",
		"Pertemuan-Dua" :"Tipe Data",
		"Pertemuan-Tiga": "Operator",
		"Pertemuan-Empat" : "Array",
		"Pertemuan-Lima" : "fungsi"
	} */
	r.POST("/materi", handler.MateriHandler)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
