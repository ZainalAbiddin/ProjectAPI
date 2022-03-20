package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MateriGolangWEB struct {
	Pertemuan1 string `json:"Pertemuan-Satu" binding:"required"`
	Pertemuan2 string `json:"Pertemuan-Dua" binding:"required"`
	Pertemuan3 string `json:"Pertemuan-Tiga" binding:"required"`
	Pertemuan4 string `json:"Pertemuan-Empat" binding:"required"`
	Pertemuan5 string `json:"Pertemuan-Lima" binding:"required"`
}

func MateriHandler(c *gin.Context) {
	var materigolang MateriGolangWEB
	err := c.ShouldBindJSON(&materigolang)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message":     "Input data berhasil",
		"Time":        time.Now(),
		"Pertemuan 1": materigolang.Pertemuan1,
		"Pertemuan 2": materigolang.Pertemuan2,
		"Pertemuan 3": materigolang.Pertemuan3,
		"Pertemuan 4": materigolang.Pertemuan4,
		"Pertemuan 5": materigolang.Pertemuan5,
	})
}
