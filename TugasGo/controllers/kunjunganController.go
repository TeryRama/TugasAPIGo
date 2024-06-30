package controllers

import (
	"net/http"
	"project-root/config"
	"project-root/models"

	"github.com/gin-gonic/gin"
)

func GetKunjungan(c *gin.Context) {
	var kunjungan []models.Kunjungan
	if err := config.DB.Preload("Pasien").Preload("Dokter").Find(&kunjungan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kunjungan)
}

func GetKunjunganByID(c *gin.Context) {
	var kunjungan models.Kunjungan
	id := c.Param("id")

	if err := config.DB.Preload("Pasien").Preload("Dokter").First(&kunjungan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kunjungan not found"})
		return
	}
	c.JSON(http.StatusOK, kunjungan)
}

func CreateKunjungan(c *gin.Context) {
	var kunjungan models.Kunjungan
	if err := c.ShouldBindJSON(&kunjungan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&kunjungan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kunjungan)
}

func UpdateKunjungan(c *gin.Context) {
	var kunjungan models.Kunjungan
	id := c.Param("id")

	if err := config.DB.First(&kunjungan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kunjungan not found"})
		return
	}

	if err := c.ShouldBindJSON(&kunjungan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&kunjungan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, kunjungan)
}

func DeleteKunjungan(c *gin.Context) {
	var kunjungan models.Kunjungan
	id := c.Param("id")

	if err := config.DB.First(&kunjungan, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kunjungan not found"})
		return
	}

	if err := config.DB.Delete(&kunjungan).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Kunjungan deleted"})
}
