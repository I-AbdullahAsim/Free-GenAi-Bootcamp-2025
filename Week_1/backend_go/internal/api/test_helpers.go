package api

import (
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"github.com/stretchr/testify/mock"
)

// StudyServiceInterface defines all study-related operations
type StudyServiceInterface interface {
	// Study Session operations
	GetAllStudySessions() ([]models.StudySession, error)
	GetStudySession(id uint) (*models.StudySession, error)
	GetStudySessionWords(sessionID uint) ([]models.Word, error)
	CreateStudySession(session *models.StudySession) error
	CreateWordReview(review *models.WordReviewItem) error
	DeleteStudySession(id uint) error

	// Study Activity operations
	CreateStudyActivity(activity *models.StudyActivity) error
	GetStudyActivity(id uint) (*models.StudyActivity, error)
	DeleteStudyActivity(id uint) error
	GetStudyActivitySessions(activityID uint) ([]models.StudySession, error)
}

// MockStudyService implements StudyServiceInterface for testing
type MockStudyService struct {
	mock.Mock
}

// Study Session methods
func (m *MockStudyService) GetAllStudySessions() ([]models.StudySession, error) {
	args := m.Called()
	return args.Get(0).([]models.StudySession), args.Error(1)
}

func (m *MockStudyService) GetStudySession(id uint) (*models.StudySession, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.StudySession), args.Error(1)
}

func (m *MockStudyService) GetStudySessionWords(sessionID uint) ([]models.Word, error) {
	args := m.Called(sessionID)
	return args.Get(0).([]models.Word), args.Error(1)
}

func (m *MockStudyService) CreateStudySession(session *models.StudySession) error {
	args := m.Called(session)
	return args.Error(0)
}

func (m *MockStudyService) CreateWordReview(review *models.WordReviewItem) error {
	args := m.Called(review)
	return args.Error(0)
}

func (m *MockStudyService) DeleteStudySession(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// Study Activity methods
func (m *MockStudyService) CreateStudyActivity(activity *models.StudyActivity) error {
	args := m.Called(activity)
	return args.Error(0)
}

func (m *MockStudyService) GetStudyActivity(id uint) (*models.StudyActivity, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.StudyActivity), args.Error(1)
}

func (m *MockStudyService) DeleteStudyActivity(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockStudyService) GetStudyActivitySessions(activityID uint) ([]models.StudySession, error) {
	args := m.Called(activityID)
	return args.Get(0).([]models.StudySession), args.Error(1)
} 