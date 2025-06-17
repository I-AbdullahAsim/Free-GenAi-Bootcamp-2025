package models

import "gorm.io/gorm"

type WordGroup struct {
	gorm.Model
	WordID  uint `gorm:"not null;uniqueIndex:idx_word_group"`
	GroupID uint `gorm:"not null;uniqueIndex:idx_word_group"`
}

func (WordGroup) TableName() string {
	return "word_groups"
}
