package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/service"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/utils"
	repo "github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/repository"
	"gorm.io/gorm"
)

// GroupHandler handles group related endpoints
type GroupHandler struct {
	service service.GroupServiceInterface
}

// NewGroupHandler creates a new group handler
func NewGroupHandler(service service.GroupServiceInterface) *GroupHandler {
	return &GroupHandler{service: service}
}

// GetGroups returns a list of groups
func (h *GroupHandler) GetGroups(c *gin.Context) {
	groups, err := h.service.ListGroups()
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, groups, "Groups retrieved successfully")
}

// GetGroup retrieves a group by ID
func (h *GroupHandler) GetGroup(c *gin.Context) {
	id := c.Param("id")
	groupID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusBadRequest, "invalid ID format")
		return
	}

	group, err := h.service.GetGroup(uint(groupID))
	if err != nil {
		if err == ErrGroupNotFound || err == repo.ErrNotFound || err == gorm.ErrRecordNotFound {
			utils.ErrorWithStatus(c, http.StatusNotFound, "Group not found")
			return
		}
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, group, "Group retrieved successfully")
}

// CreateGroup creates a new group
func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		utils.ErrorWithStatus(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.CreateGroup(&group); err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, group, "Group created successfully")
}

// UpdateGroup updates an existing group
func (h *GroupHandler) UpdateGroup(c *gin.Context) {
	id := c.Param("id")
	groupID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusBadRequest, "invalid ID format")
		return
	}

	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		utils.ErrorWithStatus(c, http.StatusBadRequest, err.Error())
		return
	}

	group.ID = uint(groupID)
	if err := h.service.UpdateGroup(&group); err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, group, "Group updated successfully")
}

// DeleteGroup deletes a group
func (h *GroupHandler) DeleteGroup(c *gin.Context) {
	id := c.Param("id")
	groupID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusBadRequest, "invalid ID format")
		return
	}

	if err := h.service.DeleteGroup(uint(groupID)); err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, nil, "Group deleted successfully")
}

// GetGroupWords retrieves all words in a group
func (h *GroupHandler) GetGroupWords(c *gin.Context) {
	id := c.Param("id")
	groupID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusBadRequest, "invalid ID format")
		return
	}

	words, err := h.service.GetGroupWords(uint(groupID))
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, words, "Group words retrieved successfully")
}

// GetGroupStudySessions retrieves study sessions associated with a group
func (h *GroupHandler) GetGroupStudySessions(c *gin.Context) {
	id := c.Param("id")
	groupID, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusBadRequest, "invalid ID format")
		return
	}

	sessions, err := h.service.GetGroupStudySessionsWithActivityName(uint(groupID))
	if err != nil {
		utils.ErrorWithStatus(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Success(c, sessions, "Group study sessions retrieved successfully")
}
