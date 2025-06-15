package models

import (
    "gorm.io/gorm"
    "time"
)

type Group struct {
    gorm.Model
    Name         string         `gorm:"not null"`
    Words        []Word         `gorm:"many2many:word_groups;"`
    StudySessions []StudySession
}