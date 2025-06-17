package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/service"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/utils"
	"strconv"
)

// WordHandler handles word related endpoints
type WordHandler struct {
	service *service.WordService
}

// NewWordHandler creates a new word handler
func NewWordHandler(service *service.WordService) *WordHandler {
	return &WordHandler{service: service}
}

// GetWords returns a paginated list of words
func (h *WordHandler) GetWords(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize := 100 // Fixed page size as per spec
	
	words, total, err := h.service.ListWords(page, pageSize)
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	totalPages := (int(total) + pageSize - 1) / pageSize
	
	utils.Success(c, gin.H{
		"total_pages": totalPages,
		"current_page": page,
		"items_per_page": pageSize,
		"total_items": total,
		"items": words,
	}, "Words retrieved successfully")
}

// CreateWord creates a new word
func (h *WordHandler) CreateWord(c *gin.Context) {
	var word models.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		utils.Error(c, err.Error())
		return
	}

	if err := h.service.CreateWord(&word); err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, word, "Word created successfully")
}

// GetWord retrieves a word by ID
func (h *WordHandler) GetWord(c *gin.Context) {
	id := c.Param("id")
	wordID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusBadRequest, "Invalid ID format")
		return
	}

	word, err := h.service.GetWord(uint(wordID))
	if err != nil {
		if err == service.ErrWordNotFound {
			utils.ErrorWithStatus(c, http.StatusNotFound, "Word not found")
			return
		}
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	utils.Success(c, word, "Word retrieved successfully")
}

// UpdateWord updates an existing word
func (h *WordHandler) UpdateWord(c *gin.Context) {
	id := c.Param("id")
	wordID, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(c, "invalid ID format")
		return
	}

	var word models.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		utils.Error(c, err.Error())
		return
	}

	word.ID = uint(wordID)
	if err := h.service.UpdateWord(&word); err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, word, "Word updated successfully")
}

// DeleteWord deletes a word
func (h *WordHandler) DeleteWord(c *gin.Context) {
	id := c.Param("id")
	wordID, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(c, "invalid ID format")
		return
	}

	if err := h.service.DeleteWord(uint(wordID)); err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, nil, "Word deleted successfully")
}

// AddWordToGroup adds a word to a group
func (h *WordHandler) AddWordToGroup(c *gin.Context) {
	wordID := c.Param("id")
	groupID := c.Param("group_id")

	wordIDInt, err := strconv.Atoi(wordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid word ID format"})
		return
	}

	groupIDInt, err := strconv.Atoi(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID format"})
		return
	}

	if err := h.service.AddWordToGroup(uint(wordIDInt), uint(groupIDInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// RemoveWordFromGroup removes a word from a group
func (h *WordHandler) RemoveWordFromGroup(c *gin.Context) {
	wordID := c.Param("id")
	groupID := c.Param("group_id")

	wordIDInt, err := strconv.Atoi(wordID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid word ID format"})
		return
	}

	groupIDInt, err := strconv.Atoi(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group ID format"})
		return
	}

	if err := h.service.RemoveWordFromGroup(uint(wordIDInt), uint(groupIDInt)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
