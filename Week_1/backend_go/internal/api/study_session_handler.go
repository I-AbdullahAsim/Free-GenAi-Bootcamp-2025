package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

// StudySessionHandler handles study session related endpoints
type StudySessionHandler struct {
	service StudyServiceInterface
}

// NewStudySessionHandler creates a new study session handler
func NewStudySessionHandler(service StudyServiceInterface) *StudySessionHandler {
	return &StudySessionHandler{service: service}
}

// CreateStudySession creates a new study session
func (h *StudySessionHandler) CreateStudySession(c *gin.Context) {
	var session models.StudySession
	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateStudySession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, session)
}

// GetStudySessions returns all study sessions
func (h *StudySessionHandler) GetStudySessions(c *gin.Context) {
	sessions, err := h.service.GetAllStudySessions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

// GetStudySession retrieves a study session by ID
func (h *StudySessionHandler) GetStudySession(c *gin.Context) {
	id := c.Param("id")
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	session, err := h.service.GetStudySession(uint(sessionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, session)
}

// GetStudySessionWords returns words associated with a study session
func (h *StudySessionHandler) GetStudySessionWords(c *gin.Context) {
	id := c.Param("id")
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	words, err := h.service.GetStudySessionWords(uint(sessionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, words)
}

// CreateWordReview creates a new word review for a study session
func (h *StudySessionHandler) CreateWordReview(c *gin.Context) {
	sessionID := c.Param("id")
	wordID := c.Param("word_id")

	// Convert both IDs to uint
	sessionIDUint, err := strconv.Atoi(sessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session ID format"})
		return
	}

	wordIDUint, err := strconv.Atoi(wordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid word ID format"})
		return
	}

	// Use a dedicated struct for payload validation
	type reviewPayload struct {
		IsCorrect *bool `json:"is_correct"`
	}
	var payload reviewPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON payload"})
		return
	}
	if payload.IsCorrect == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required field: is_correct"})
		return
	}

	review := models.WordReviewItem{
		SessionID: uint(sessionIDUint),
		WordID:    uint(wordIDUint),
		IsCorrect: *payload.IsCorrect,
	}

	if err := h.service.CreateWordReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Word review updated successfully",
		"data":    review,
	})
}

// DeleteStudySession deletes a study session
func (h *StudySessionHandler) DeleteStudySession(c *gin.Context) {
	id := c.Param("id")
	sessionID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	err = h.service.DeleteStudySession(uint(sessionID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
