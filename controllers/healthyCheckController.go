package controllers

import (
	models "../applicationModels"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @tags HealthyCheck
// @Summary healthy check for this service
// @Produce json
// @Success 200 {object} applicationModels.HealthyCheckResponse "Get current api status"
// @Router /healthyCheck [get]
func Healthy(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, models.HealthyCheckResponse{
		Message: "Welcome to simple api demo",
		Version: "1.0.0",
	})
}
