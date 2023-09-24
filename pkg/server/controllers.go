package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kopoze/kpz/pkg/app"
	"github.com/kopoze/kpz/pkg/config"
	"github.com/kopoze/kpz/pkg/hosts"
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
	conf := config.LoadConfig()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create currApp
	currApp := app.App{Name: input.Name, Subdomain: input.Subdomain, Port: input.Port}
	app.DB.Create(&currApp)
	if conf.Kopoze.Mode == "local" {
		hosts.AddSubdomain(currApp.Subdomain)
	}
	c.JSON(http.StatusOK, gin.H{"data": currApp})
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
	conf := config.LoadConfig()

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

	oldApp := currApp
	app.DB.Model(&currApp).Updates(input)
	if conf.Kopoze.Mode == "local" && oldApp.Subdomain != currApp.Subdomain {
		hosts.RemoveSubdomain(oldApp.Subdomain)
		hosts.AddSubdomain(currApp.Subdomain)
	}
	c.JSON(http.StatusOK, gin.H{"data": currApp})
}

func DeleteApp(c *gin.Context) {
	// Get model if exist
	var currApp app.App
	conf := config.LoadConfig()

	if err := app.DB.Where("id = ?", c.Param("id")).First(&currApp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if conf.Kopoze.Mode == "local" {
		hosts.RemoveSubdomain(currApp.Subdomain)
	}
	app.DB.Delete(&currApp)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
