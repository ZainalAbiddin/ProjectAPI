package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// localhost:8080/user/suddin/
func Nama(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

// localhost:8080/product?judul=Daun-Jatuh&penulis=tere-liye&harga=10000&penerbit=gramedia&cetakan=5
func ProductHandler(c *gin.Context) {
	judul := c.Query("judul")
	penulis := c.Query("penulis")
	penerbit := c.Query("penerbit")
	harga := c.Query("harga")
	cetakan := c.Query("cetakan")

	c.JSON(http.StatusOK, gin.H{
		"message":   "ini buku",
		"item name": judul,
		"penulis":   penulis,
		"penerbit":  penerbit,
		"harga":     harga,
		"cetakan":   cetakan,
		"time":      time.Now(),
	})
}

// localhost:8080/smartphone/smartphone/xiaomi/redmi-note-5/ram-4-gb/3000000
func ProductSmartphone(c *gin.Context) {
	merk := c.Param("merk")
	jenis := c.Param("jenis")
	spesifikasi := c.Param("spesifikasi")
	harga := c.Param("harga")
	namaproduk := c.Param("namaproduk")
	c.JSON(http.StatusOK, gin.H{
		"message":     "ini smartphone",
		"item name":   namaproduk,
		"merk":        merk,
		"jenis":       jenis,
		"spesifikasi": spesifikasi,
		"harga":       harga,
		"time":        time.Now(),
	})
}

// contoh soal 1
// localhost:8080/item?judul=filosofi-teras&harga=76000
func ItemHandler(c *gin.Context) {
	judul := c.Query("judul")
	harga := c.Query("harga")

	c.JSON(http.StatusOK, gin.H{
		"message":   "ini buku",
		"item name": judul,
		"price":     harga,
		"time":      time.Now(),
	})
}

// contoh Soal 2
// localhost:8080/id/2/product?judul="tv"&harga=400000
// beda parametr vs query
// parameter langsung dituluis dalam product/param/
// query harus ditulus ?=query&query
func ItemhandlerID(c *gin.Context) {
	id := c.Param("id")
	title := c.Query("judul")
	price := c.Query("harga")
	c.JSON(http.StatusOK, gin.H{
		"id":           id,
		"message":      "item name",
		"product name": title,
		"price":        price,
	})
}

func Welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
