package service

import (
	"gorm.io/gorm"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

type GroupService struct {
	db *gorm.DB
}

func NewGroupService(db *gorm.DB) *GroupService {
	return &GroupService{
		db: db,
	}
}

// ListGroups retrieves all groups
func (s *GroupService) ListGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := s.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// GetGroup retrieves a group by ID
func (s *GroupService) GetGroup(id uint) (*models.Group, error) {
	var group models.Group
	if err := s.db.Preload("Words").First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

// CreateGroup creates a new group
func (s *GroupService) CreateGroup(group *models.Group) error {
	return s.db.Create(group).Error
}

// UpdateGroup updates an existing group
func (s *GroupService) UpdateGroup(group *models.Group) error {
	return s.db.Save(group).Error
}

// DeleteGroup deletes a group by ID
func (s *GroupService) DeleteGroup(id uint) error {
	return s.db.Delete(&models.Group{}, id).Error
}

// GetGroupStudySessions retrieves study sessions associated with a group
func (s *GroupService) GetGroupStudySessions(groupID uint) ([]models.StudySession, error) {
	var sessions []models.StudySession
	if err := s.db.Where("group_id = ?", groupID).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// GetGroupWords retrieves all words in a group
func (s *GroupService) GetGroupWords(groupID uint) ([]models.Word, error) {
	var words []models.Word
	if err := s.db.Model(&models.Word{}).
		Joins("INNER JOIN word_groups ON words.id = word_groups.word_id").
		Where("word_groups.group_id = ?", groupID).
		Find(&words).Error; err != nil {
		return nil, err
	}
	return words, nil
}

// AddWordToGroup adds a word to a group
func (s *GroupService) AddWordToGroup(wordID uint, groupID uint) error {
	wordGroup := models.WordGroup{
		WordID:  wordID,
		GroupID: groupID,
	}
	return s.db.Create(&wordGroup).Error
}

// RemoveWordFromGroup removes a word from a group
func (s *GroupService) RemoveWordFromGroup(wordID uint, groupID uint) error {
	return s.db.Where("word_id = ? AND group_id = ?", wordID, groupID).
		Delete(&models.WordGroup{}).Error
}

type GroupStudySessionResponse struct {
	ID           uint   `json:"id"`
	GroupID      uint   `json:"group_id"`
	ActivityID   uint   `json:"activity_id"`
	ActivityName string `json:"activity_name"`
}

func (s *GroupService) GetGroupStudySessionsWithActivityName(groupID uint) ([]GroupStudySessionResponse, error) {
	var results []GroupStudySessionResponse
	err := s.db.Table("study_sessions").
		Select("study_sessions.id, study_sessions.group_id, study_sessions.activity_id, study_activities.name as activity_name").
		Joins("JOIN study_activities ON study_sessions.activity_id = study_activities.id").
		Where("study_sessions.group_id = ?", groupID).
		Scan(&results).Error
	return results, err
}

// GroupServiceInterface defines the methods required for a group service
type GroupServiceInterface interface {
	ListGroups() ([]models.Group, error)
	GetGroup(id uint) (*models.Group, error)
	CreateGroup(group *models.Group) error
	UpdateGroup(group *models.Group) error
	DeleteGroup(id uint) error
	GetGroupWords(groupID uint) ([]models.Word, error)
	GetGroupStudySessionsWithActivityName(groupID uint) ([]GroupStudySessionResponse, error)
}