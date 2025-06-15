package models

import "gorm.io/gorm"

type WordGroup struct {
    gorm.Model
    WordID   uint `gorm:"not null"`
    GroupID  uint `gorm:"not null"`
}