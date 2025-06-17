package models

import (
	"gorm.io/gorm"
)

// StudyProgress represents study progress statistics
type StudyProgress struct {
	TotalSessions      int     `json:"total_sessions"`
	AverageSuccessRate float64 `json:"average_success_rate"`
	StreakDays         int     `json:"streak_days"`
	CorrectWords       int     `json:"correct_words"`
	IncorrectWords     int     `json:"incorrect_words"`
	SuccessRate        float64 `json:"success_rate"`
	TotalWordsStudied  int     `json:"total_words_studied"`
	TotalAvailableWords int    `json:"total_available_words"`
}

// QuickStats represents quick statistics about learning progress
type QuickStats struct {
	TotalWords         int     `json:"total_words"`
	CorrectWords       int     `json:"correct_words"`
	IncorrectWords     int     `json:"incorrect_words"`
	SuccessRate        float64 `json:"success_rate"`
	TotalStudySessions int     `json:"total_study_sessions"`
	TotalActiveGroups  int     `json:"total_active_groups"`
	StudyStreakDays    int     `json:"study_streak_days"`
}

// StudySession represents a study session
type StudySession struct {
	gorm.Model
	GroupID         uint   `gorm:"not null"`
	ActivityID      uint   `gorm:"not null"`
	StudyActivities []StudyActivity `gorm:"foreignKey:StudySessionID;references:ID"`
	Group           Group  `gorm:"foreignKey:GroupID;references:ID"`
}

// WordReviewItem represents a word review record
type WordReviewItem struct {
	gorm.Model
	SessionID uint   `gorm:"not null"`
	WordID    uint   `gorm:"not null"`
	IsCorrect bool   `gorm:"not null"`
	StudySessionID uint   `gorm:"not null"`
	StudySession StudySession `gorm:"foreignKey:StudySessionID;references:ID"`
	Word       Word `gorm:"foreignKey:WordID;references:ID"`
}

// Settings represents application settings
type Settings struct {
	gorm.Model
	ResetHistory bool `json:"reset_history"`
	FullReset    bool `json:"full_reset"`
}
