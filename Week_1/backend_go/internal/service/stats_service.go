package service

import (
	"sort"
	"time"
	
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/repository"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

// StatsService handles statistics-related business logic
type StatsService struct {
	sessionRepo repository.StudySessionRepository
	reviewRepo repository.WordReviewRepository
}

// NewStatsService creates a new StatsService instance
func NewStatsService(sessionRepo repository.StudySessionRepository, reviewRepo repository.WordReviewRepository) *StatsService {
	return &StatsService{
		sessionRepo: sessionRepo,
		reviewRepo: reviewRepo,
	}
}

// CalculateSuccessRate calculates the success rate for a study session
func (s *StatsService) CalculateSuccessRate(sessionID uint) (float64, error) {
	reviews, err := s.reviewRepo.FindWordReviews(sessionID)
	if err != nil {
		return 0, err
	}

	if len(reviews) == 0 {
		return 0, nil
	}

	correct := 0
	for _, review := range reviews {
		if review.IsCorrect {
			correct++
		}
	}

	return float64(correct) / float64(len(reviews)), nil
}

// CalculateStudyStreak calculates the current study streak in days
func (s *StatsService) CalculateStudyStreak() (int, error) {
	sessions, err := s.sessionRepo.FindStudySessions()
	if err != nil {
		return 0, err
	}

	if len(sessions) == 0 {
		return 0, nil
	}

	// Sort sessions by date
	type sessionWithDate struct {
		Session *models.StudySession
		Date    time.Time
	}

	var sessionsWithDate []sessionWithDate
	for _, session := range sessions {
		sessionsWithDate = append(sessionsWithDate, sessionWithDate{
			Session: &session,
			Date:    session.CreatedAt,
		})
	}

	// Sort by date
	sort.Slice(sessionsWithDate, func(i, j int) bool {
		return sessionsWithDate[i].Date.After(sessionsWithDate[j].Date)
	})

	// Calculate streak
	streak := 0
	lastDate := time.Time{}
	for _, session := range sessionsWithDate {
		if lastDate.IsZero() || session.Date.Sub(lastDate) <= 24*time.Hour {
			streak++
		} else {
			streak = 1
		}
		lastDate = session.Date
	}

	return streak, nil
}

// GetStudyProgress returns study progress statistics
func (s *StatsService) GetStudyProgress() (*models.StudyProgress, error) {
	progress := &models.StudyProgress{}

	// Get total sessions
	sessions, err := s.sessionRepo.FindStudySessions()
	if err != nil {
		return nil, err
	}

	progress.TotalSessions = len(sessions)

	// Calculate average success rate
	var totalRate float64
	for _, session := range sessions {
		rate, err := s.CalculateSuccessRate(session.ID)
		if err != nil {
			return nil, err
		}
		totalRate += rate
	}

	if len(sessions) > 0 {
		progress.AverageSuccessRate = totalRate / float64(len(sessions))
	}

	// Get streak
	streak, err := s.CalculateStudyStreak()
	if err != nil {
		return nil, err
	}

	progress.StreakDays = streak

	return progress, nil
}

// GetQuickStats returns quick statistics about learning progress
func (s *StatsService) GetQuickStats() (*models.QuickStats, error) {
	stats := &models.QuickStats{}

	// Get total words
	reviews, err := s.reviewRepo.FindWordReviews(0) // Get all reviews
	if err != nil {
		return nil, err
	}

	stats.TotalWords = len(reviews)

	// Count correct and incorrect reviews
	correct := 0
	for _, review := range reviews {
		if review.IsCorrect {
			correct++
		}
	}

	stats.CorrectWords = correct
	stats.IncorrectWords = len(reviews) - correct

	// Calculate success rate
	if len(reviews) > 0 {
		stats.SuccessRate = float64(correct) / float64(len(reviews))
	}

	return stats, nil
}

// GetLastStudySession returns the most recent study session
func (s *StatsService) GetLastStudySession() (*models.StudySession, error) {
	sessions, err := s.sessionRepo.FindStudySessions()
	if err != nil {
		return nil, err
	}
	if len(sessions) == 0 {
		return nil, nil
	}
	return &sessions[len(sessions)-1], nil
}