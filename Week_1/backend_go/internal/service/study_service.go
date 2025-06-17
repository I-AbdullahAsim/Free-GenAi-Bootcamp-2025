package service

import (
	"time"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/repository"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/utils"
)

type StudyService struct {
	sessionRepo repository.StudySessionRepository
	activityRepo repository.StudyActivityRepository
	reviewRepo repository.WordReviewRepository
}

func NewStudyService(sessionRepo repository.StudySessionRepository, activityRepo repository.StudyActivityRepository, reviewRepo repository.WordReviewRepository) *StudyService {
	return &StudyService{
		sessionRepo: sessionRepo,
		activityRepo: activityRepo,
		reviewRepo: reviewRepo,
	}
}

// Study session management
func (s *StudyService) GetAllStudySessions() ([]models.StudySession, error) {
	sessions, err := s.sessionRepo.FindStudySessions()
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *StudyService) CreateStudySession(session *models.StudySession) error {
	return s.sessionRepo.CreateStudySession(session)
}

func (s *StudyService) GetStudySession(id uint) (*models.StudySession, error) {
	return s.sessionRepo.FindStudySessionByID(id)
}

func (s *StudyService) GetStudySessionWords(id uint) ([]models.Word, error) {
	return s.sessionRepo.FindStudySessionWords(id)
}

func (s *StudyService) UpdateStudySession(session *models.StudySession) error {
	return s.sessionRepo.UpdateStudySession(session)
}

func (s *StudyService) DeleteStudySession(id uint) error {
	return s.sessionRepo.DeleteStudySession(id)
}

// Study activity management
func (s *StudyService) CreateStudyActivity(activity *models.StudyActivity) error {
	return s.activityRepo.CreateStudyActivity(activity)
}

func (s *StudyService) GetStudyActivity(id uint) (*models.StudyActivity, error) {
	return s.activityRepo.FindStudyActivityByID(id)
}

func (s *StudyService) UpdateStudyActivity(activity *models.StudyActivity) error {
	return s.activityRepo.UpdateStudyActivity(activity)
}

func (s *StudyService) DeleteStudyActivity(id uint) error {
	return s.activityRepo.DeleteStudyActivity(id)
}

// GetStudyActivitySessions retrieves study sessions associated with an activity
func (s *StudyService) GetStudyActivitySessions(activityID uint) ([]models.StudySession, error) {
	return s.activityRepo.FindStudyActivitySessions(activityID)
}

// Word review item management
func (s *StudyService) CreateWordReview(review *models.WordReviewItem) error {
	return s.reviewRepo.CreateWordReview(review)
}

func (s *StudyService) GetWordReviewItem(id uint) (*models.WordReviewItem, error) {
	return s.reviewRepo.FindWordReviewByID(id)
}

func (s *StudyService) UpdateWordReviewItem(reviewItem *models.WordReviewItem) error {
	return s.reviewRepo.UpdateWordReview(reviewItem)
}

func (s *StudyService) DeleteWordReviewItem(id uint) error {
	return s.reviewRepo.DeleteWordReview(id)
}

// Dashboard statistics
func (s *StudyService) GetLastStudySession() (*models.StudySession, error) {
	lastSession, err := s.sessionRepo.FindLastStudySession()
	if err != nil {
		return nil, err
	}
	if lastSession != nil {
		// Include group data
		group, err := s.activityRepo.FindGroupByID(lastSession.GroupID)
		if err != nil {
			return nil, err
		}
		if group != nil {
			lastSession.Group = *group
		}
	}
	return lastSession, nil
}

func (s *StudyService) GetStudyProgress() (*models.StudyProgress, error) {
	progress := &models.StudyProgress{}
	
	// Get total sessions
	sessions, err := s.sessionRepo.FindStudySessions()
	if err != nil {
		return nil, err
	}
	progress.TotalSessions = len(sessions)
	
	// Get total words studied
	wordsStudied := make(map[uint]bool)
	for _, session := range sessions {
		sessionWords, err := s.sessionRepo.FindStudySessionWords(session.ID)
		if err != nil {
			return nil, err
		}
		for _, word := range sessionWords {
			wordsStudied[word.ID] = true
		}
	}
	progress.TotalWordsStudied = len(wordsStudied)
	
	// Get total available words
	allWords, err := s.sessionRepo.FindStudySessionWords(0) // 0 means all words
	if err != nil {
		return nil, err
	}
	progress.TotalAvailableWords = len(allWords)
	
	return progress, nil
}

func (s *StudyService) GetQuickStats() (*models.QuickStats, error) {
	stats := &models.QuickStats{}

	// Get total study sessions
	sessions, err := s.sessionRepo.FindStudySessions()
	if err != nil {
		return nil, err
	}
	stats.TotalStudySessions = len(sessions)

	// Get total active groups
	activeGroups := make(map[uint]bool)
	for _, session := range sessions {
		activeGroups[session.GroupID] = true
	}
	stats.TotalActiveGroups = len(activeGroups)

	// Calculate success rate
	reviews, err := s.reviewRepo.FindAllWordReviews()
	if err != nil {
		return nil, err
	}
	totalAttempts := len(reviews)
	totalCorrect := 0
	for _, review := range reviews {
		if review.IsCorrect {
			totalCorrect++
		}
	}
	if totalAttempts > 0 {
		stats.SuccessRate = float64(totalCorrect) / float64(totalAttempts) * 100
	} else {
		stats.SuccessRate = 0
	}

	// Calculate study streak days
	lastSession, err := s.sessionRepo.FindLastStudySession()
	if err != nil {
		return nil, err
	}
	if lastSession != nil {
		stats.StudyStreakDays = s.calculateStreakDays(lastSession.CreatedAt)
	}

	return stats, nil
}

func (s *StudyService) calculateStreakDays(lastStudyTime time.Time) int {
	current := utils.GetStartOfDay(time.Now())
	lastStudyDay := utils.GetStartOfDay(lastStudyTime)
	streak := 0
	for utils.IsSameDay(current, lastStudyDay) || utils.CalculateDateDifference(lastStudyDay, current) <= 1 {
		streak++
		lastStudyDay = lastStudyDay.AddDate(0, 0, -1)
	}
	return streak
}

