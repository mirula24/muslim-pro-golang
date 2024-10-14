package controller

import (
	"muslim_pro/config"
	"muslim_pro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCaption(c *gin.Context) {

	input := struct {
		Email  string `json:"email" binding:"required"`
		Baris1 string `json:"baris1"`
		Baris2 string `json:"baris2"`
		Baris3 string `json:"baris3"`
		Baris4 string `json:"baris4"`
	}{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var caption models.Caption

	caption.Email = input.Email
	caption.Baris1 = input.Baris1
	caption.Baris2 = input.Baris2
	caption.Baris3 = input.Baris3
	caption.Baris4 = input.Baris4

	if err := config.DB.Create(&caption).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, caption)

}

func GetAllCaption(c *gin.Context) {
	var captions []models.Caption
	if err := config.DB.Find(&captions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, captions)

}

func GetOneCaption(c *gin.Context) {
	id := c.Param("id")
	var caption models.Caption
	if err := config.DB.First(&caption, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Caption not found"})
		return
	}

	c.JSON(http.StatusOK, caption)
}
func GetRandomCaption(c *gin.Context) {
	var caption models.Caption

	if err := config.DB.Raw("SELECT * FROM captions WHERE status = true ORDER BY RANDOM() LIMIT 1").Scan(&caption).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Caption not found"})
		return
	}

	c.JSON(http.StatusOK, caption)

}

func UpdateStatusCaption(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hero ID is required"})
		return
	}

	var caption models.Caption

	if err := config.DB.First(&caption, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Caption not found"})
		return
	}

	caption.Status = true

	// Simpan perubahan ke database
	if err := config.DB.Save(&caption).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Kembalikan response sukses
	c.JSON(http.StatusAccepted, caption)

}

func DeleteCaption(c *gin.Context) {
	id := c.Param("id")
	var caption models.Caption
	if err := config.DB.Delete(&caption, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Caption not found"})
		return
	}
	c.JSON(http.StatusAccepted, "Caption deleted")
}
