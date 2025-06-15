package models

import (
    "gorm.io/gorm"
    "time"
)

type StudySession struct {
    gorm.Model
    GroupID          uint           `gorm:"not null"`
    CreatedAt        time.Time      `gorm:"not null"`
    StudyActivityID  uint           `gorm:"not null"`
    StudyActivity    StudyActivity
    ReviewItems      []WordReviewItem
    Group            Group
}