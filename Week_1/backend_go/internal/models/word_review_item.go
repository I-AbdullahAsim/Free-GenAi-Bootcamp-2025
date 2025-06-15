package models

import (
    "gorm.io/gorm"
    "time"
)

type WordReviewItem struct {
    gorm.Model
    WordID          uint        `gorm:"not null"`
    StudySessionID  uint        `gorm:"not null"`
    Correct         bool        `gorm:"not null"`
    CreatedAt       time.Time   `gorm:"not null"`
    Word            Word
    StudySession    StudySession
}