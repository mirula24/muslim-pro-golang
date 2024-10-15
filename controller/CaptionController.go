package controller

import (
	"muslim_pro/config"
	"muslim_pro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCaptionInput struct {
	Email  string `json:"email" binding:"required"`
	Baris1 string `json:"baris1"`
	Baris2 string `json:"baris2"`
	Baris3 string `json:"baris3"`
	Baris4 string `json:"baris4"`
}
// CreateCaption godoc
// @Summary Create a new caption
// @Description Create a new caption with the input payload
// @Tags captions
// @Accept json
// @Produce json
// @Param caption body CreateCaptionInput true "Create caption"
// @Success 201 {object} models.Caption
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /captions [post]
func CreateCaption(c *gin.Context) {

	var input CreateCaptionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var caption models.Caption

	if err := config.DB.Raw("SELECT * FROM captions WHERE status = false AND email = ?", input.Email).Scan(&caption).Error; err != nil {
		c.JSON(http.StatusNotImplemented, gin.H{"error": "the caption by email already submit"})
		return
	}

	caption.Email = input.Email
	caption.Baris1 = input.Baris1
	caption.Baris2 = input.Baris2
	caption.Baris3 = input.Baris3
	caption.Baris4 = input.Baris4
	caption.Status = false

	if err := config.DB.Create(&caption).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, caption)

}

// GetAllCaption godoc
// @Summary Get all captions
// @Description Get a list of all captions
// @Tags captions
// @Produce json
// @Success 200 {array} models.Caption
// @Failure 500 {object} map[string]string
// @Router /captions [get]
func GetAllCaption(c *gin.Context) {
	var captions []models.Caption
	if err := config.DB.Find(&captions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, captions)

}

// GetOneCaption godoc
// @Summary Get a caption
// @Description Get a caption by its ID
// @Tags captions
// @Produce json
// @Param id path int true "Caption ID"
// @Success 200 {object} models.Caption
// @Failure 404 {object} map[string]string
// @Router /captions/{id} [get]
func GetOneCaption(c *gin.Context) {
	id := c.Param("id")
	var caption models.Caption
	if err := config.DB.First(&caption, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Caption not found"})
		return
	}

	c.JSON(http.StatusOK, caption)
}

// GetRandomCaption godoc
// @Summary Get a random caption
// @Description Get a random caption with status true
// @Tags captions
// @Produce json
// @Success 200 {object} models.Caption
// @Failure 404 {object} map[string]string
// @Router /captions/random [get]
func GetRandomCaption(c *gin.Context) {
	var caption models.Caption

	if err := config.DB.Raw("SELECT * FROM captions WHERE status = true ORDER BY RANDOM() LIMIT 1").Scan(&caption).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Caption not found"})
		return
	}

	c.JSON(http.StatusOK, caption)

}

// UpdateStatusCaption godoc
// @Summary Update caption status
// @Description Update the status of a caption to true
// @Tags captions
// @Produce json
// @Param id path int true "Caption ID"
// @Success 202 {object} models.Caption
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /captions/{id}/status [put]
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

// DeleteCaption godoc
// @Summary Delete a caption
// @Description Delete a caption by its ID
// @Tags captions
// @Produce json
// @Param id path int true "Caption ID"
// @Success 202 {string} string "Caption deleted"
// @Failure 404 {object} map[string]string
// @Router /captions/{id} [delete]
func DeleteCaption(c *gin.Context) {
	id := c.Param("id")
	var caption models.Caption
	if err := config.DB.Delete(&caption, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Caption not found"})
		return
	}
	c.JSON(http.StatusAccepted, "Caption deleted")
}
