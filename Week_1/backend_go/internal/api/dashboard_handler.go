package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/service"
)

// DashboardHandler handles dashboard related endpoints
type DashboardHandler struct {
	service *service.StudyService
}

// NewDashboardHandler creates a new dashboard handler
func NewDashboardHandler(service *service.StudyService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

// GetDashboardStats returns dashboard statistics
func (h *DashboardHandler) GetDashboardStats(c *gin.Context) {
	stats, err := h.service.GetQuickStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

// GetLastStudySession returns the last study session
func (h *DashboardHandler) GetLastStudySession(c *gin.Context) {
	session, err := h.service.GetLastStudySession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

// GetStudyProgress returns study progress statistics
func (h *DashboardHandler) GetStudyProgress(c *gin.Context) {
	progress, err := h.service.GetStudyProgress()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, progress)
}

// GetQuickStats returns quick statistics about learning progress
func (h *DashboardHandler) GetQuickStats(c *gin.Context) {
	stats, err := h.service.GetQuickStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}
