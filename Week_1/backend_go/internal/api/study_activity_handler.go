package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"strconv"
)

// StudyActivityHandler handles study activity related endpoints
type StudyActivityHandler struct {
	service StudyServiceInterface
}

// NewStudyActivityHandler creates a new study activity handler
func NewStudyActivityHandler(service StudyServiceInterface) *StudyActivityHandler {
	return &StudyActivityHandler{service: service}
}

// CreateStudyActivity creates a new study activity
func (h *StudyActivityHandler) CreateStudyActivity(c *gin.Context) {
	var activity models.StudyActivity
	if err := c.ShouldBindJSON(&activity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateStudyActivity(&activity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, activity)
}

// GetStudyActivity retrieves a study activity by ID
func (h *StudyActivityHandler) GetStudyActivity(c *gin.Context) {
	id := c.Param("id")
	activityID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	activity, err := h.service.GetStudyActivity(uint(activityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activity)
}

// DeleteStudyActivity deletes a study activity
func (h *StudyActivityHandler) DeleteStudyActivity(c *gin.Context) {
	id := c.Param("id")
	activityID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	if err := h.service.DeleteStudyActivity(uint(activityID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetStudyActivitySessions retrieves study sessions associated with an activity
func (h *StudyActivityHandler) GetStudyActivitySessions(c *gin.Context) {
	id := c.Param("id")
	activityID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	sessions, err := h.service.GetStudyActivitySessions(uint(activityID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessions)
}
