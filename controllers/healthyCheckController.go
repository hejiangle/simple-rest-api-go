package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @tag.name HealthyCheck
// @Summary healthy check for this service
// @Produce json
// @Success 200 {string} string "Get current api status"
// @Router /healthyCheck [get]
func Healthy(c *gin.Context) {
	w := c.Writer
	w.Header().Set("Content-Type", "application/type")

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message": "Welcome to simple api demo", "version": "1.0"}`))
}
