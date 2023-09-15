package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kopoze/kpz/pkg/app"
)

func ListApps(c *gin.Context) {
	var apps []app.App
	app.DB.Find(&apps)

	c.JSON(http.StatusOK, gin.H{"data": apps})
}

type CreateAppInput struct {
	Name      string `json:"name" binding:"required"`
	Subdomain string `json:"subdomain" binding:"required"`
	Port      string `json:"port" binding:"required"`
}

func CreateApp(c *gin.Context) {
	// Validate input
	var input CreateAppInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := app.App{Name: input.Name, Subdomain: input.Subdomain, Port: input.Port}
	app.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindApp(c *gin.Context) {
	// Get model if exist
	var currApp app.App

	if err := app.DB.Where("id = ?", c.Param("id")).First(&currApp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": currApp})
}

type UpdateAppInput struct {
	Name      string `json:"name"`
	Subdomain string `json:"subdomain"`
	Port      string `json:"port"`
}

func UpdateApp(c *gin.Context) {
	// Get model if exist
	var currApp app.App
	if err := app.DB.Where("id = ?", c.Param("id")).First(&currApp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateAppInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app.DB.Model(&currApp).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": currApp})
}

func DeleteApp(c *gin.Context) {
	// Get model if exist
	var currApp app.App
	if err := app.DB.Where("id = ?", c.Param("id")).First(&currApp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	app.DB.Delete(&currApp)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
