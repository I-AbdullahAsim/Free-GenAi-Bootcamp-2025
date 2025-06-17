package models

import (
    "gorm.io/gorm"
    "time"
)

type StudyActivity struct {
    gorm.Model
    StudySessionID uint   `gorm:"not null"`
    GroupID        uint   `gorm:"not null"`
    CreatedAt      time.Time
    StudySession   StudySession `gorm:"foreignKey:StudySessionID;references:ID"`
    Group          Group
}