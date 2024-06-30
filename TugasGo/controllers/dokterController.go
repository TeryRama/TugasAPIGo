package controllers

import (
	"net/http"
	"project-root/config"
	"project-root/models"

	"github.com/gin-gonic/gin"
)

func GetDokter(c *gin.Context) {
	var dokter []models.Dokter
	if err := config.DB.Find(&dokter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dokter)
}

func GetDokterByID(c *gin.Context) {
	var dokter models.Dokter
	id := c.Param("id")

	if err := config.DB.First(&dokter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dokter not found"})
		return
	}
	c.JSON(http.StatusOK, dokter)
}

func CreateDokter(c *gin.Context) {
	var dokter models.Dokter
	if err := c.ShouldBindJSON(&dokter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&dokter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dokter)
}

func UpdateDokter(c *gin.Context) {
	var dokter models.Dokter
	id := c.Param("id")

	if err := config.DB.First(&dokter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dokter not found"})
		return
	}

	if err := c.ShouldBindJSON(&dokter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&dokter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dokter)
}

func DeleteDokter(c *gin.Context) {
	var dokter models.Dokter
	id := c.Param("id")

	if err := config.DB.First(&dokter, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dokter not found"})
		return
	}

	if err := config.DB.Delete(&dokter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Dokter deleted"})
}
