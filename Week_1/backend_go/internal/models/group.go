package models

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name          string         `gorm:"not null;uniqueIndex" json:"name"`
	Words         []Word         `gorm:"many2many:word_groups;" json:"words,omitempty"`
	StudySessions []StudySession `gorm:"foreignKey:GroupID;references:ID" json:"study_sessions,omitempty"`
}
