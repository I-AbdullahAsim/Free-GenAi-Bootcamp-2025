package service

import (
	"gorm.io/gorm"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week-1/backend_go/internal/models"
)

type GroupService struct {
	db *gorm.DB
}

func NewGroupService(db *gorm.DB) *GroupService {
	return &GroupService{db: db}
}

func (s *GroupService) CreateGroup(group *models.Group) error {
	return s.db.Create(group).Error
}

func (s *GroupService) GetGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := s.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (s *GroupService) GetGroupByID(id uint) (*models.Group, error) {
	var group models.Group
	if err := s.db.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (s *GroupService) UpdateGroup(group *models.Group) error {
	return s.db.Save(group).Error
}

func (s *GroupService) DeleteGroup(id uint) error {
	return s.db.Delete(&models.Group{}, id).Error
}