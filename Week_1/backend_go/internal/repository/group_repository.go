package repository

import (
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"gorm.io/gorm"
)

type GroupRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{db: db}
}

func (r *GroupRepositoryImpl) FindGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := r.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *GroupRepositoryImpl) FindGroupByID(id uint) (*models.Group, error) {
	var group models.Group
	if err := r.db.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRepositoryImpl) CreateGroup(group *models.Group) error {
	return r.db.Create(group).Error
}

func (r *GroupRepositoryImpl) UpdateGroup(group *models.Group) error {
	return r.db.Save(group).Error
}

func (r *GroupRepositoryImpl) DeleteGroup(id uint) error {
	return r.db.Delete(&models.Group{}, id).Error
}

func (r *GroupRepositoryImpl) FindGroupWords(groupID uint) ([]models.Word, error) {
	var words []models.Word
	if err := r.db.Model(&models.Group{ID: groupID}).Association("Words").Find(&words); err != nil {
		return nil, err
	}
	return words, nil
}

func (r *GroupRepositoryImpl) FindGroupStudySessions(groupID uint) ([]models.StudySession, error) {
	var sessions []models.StudySession
	if err := r.db.Model(&models.Group{ID: groupID}).Association("StudySessions").Find(&sessions); err != nil {
		return nil, err
	}
	return sessions, nil
}

// Temporary placeholder to satisfy Go compiler
func init() {}