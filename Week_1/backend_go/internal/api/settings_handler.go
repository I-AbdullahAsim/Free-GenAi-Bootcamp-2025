package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/service"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

// SettingsHandler handles settings related endpoints
type SettingsHandler struct {
	service service.SettingsServiceInterface
}

// NewSettingsHandler creates a new settings handler
func NewSettingsHandler(service service.SettingsServiceInterface) *SettingsHandler {
	return &SettingsHandler{service: service}
}

// GetSettings retrieves user settings
func (h *SettingsHandler) GetSettings(c *gin.Context) {
	settings, err := h.service.GetSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}

// UpdateSettings updates user settings
func (h *SettingsHandler) UpdateSettings(c *gin.Context) {
	var settings models.Settings
	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateSettings(&settings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, settings)
}

// ResetHistory resets the user's learning history
func (h *SettingsHandler) ResetHistory(c *gin.Context) {
	if err := h.service.ResetHistory(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// FullReset performs a complete system reset
func (h *SettingsHandler) FullReset(c *gin.Context) {
	if err := h.service.FullReset(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
