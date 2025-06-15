package models

import (
    "gorm.io/gorm"
)

type Word struct {
    gorm.Model
    ArabicWord    string         `gorm:"not null"`
    EnglishWord   string         `gorm:"not null"`
    Parts         []byte         `gorm:"type:json"`
    Groups        []Group        `gorm:"many2many:word_groups;"`
    ReviewItems   []WordReviewItem
}