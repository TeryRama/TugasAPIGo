// controllers/pasienController.go
package controllers

import (
	"net/http"

	"project-root/config"
	"project-root/models"

	"github.com/gin-gonic/gin"
)

func CreatePasien(c *gin.Context) {
	var pasien models.Pasien
	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&pasien)
	c.JSON(http.StatusOK, pasien)
}

func GetPasien(c *gin.Context) {
	var pasien []models.Pasien
	config.DB.Find(&pasien)
	c.JSON(http.StatusOK, pasien)
}

func GetPasienByID(c *gin.Context) {
	var pasien models.Pasien
	id := c.Param("id")
	if err := config.DB.First(&pasien, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pasien not found"})
		return
	}
	c.JSON(http.StatusOK, pasien)
}

func UpdatePasien(c *gin.Context) {
	var pasien models.Pasien
	id := c.Param("id")
	if err := config.DB.First(&pasien, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pasien not found"})
		return
	}
	if err := c.ShouldBindJSON(&pasien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&pasien)
	c.JSON(http.StatusOK, pasien)
}

func DeletePasien(c *gin.Context) {
	var pasien models.Pasien
	id := c.Param("id")
	if err := config.DB.First(&pasien, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pasien not found"})
		return
	}
	config.DB.Delete(&pasien)
	c.JSON(http.StatusOK, gin.H{"message": "Pasien deleted"})
}
