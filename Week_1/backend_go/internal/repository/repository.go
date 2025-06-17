package repository

import "github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"


// WordRepository defines methods for word operations
type WordRepository interface {
	FindWords(page, pageSize int) ([]models.Word, int64, error) // Returns words, total count, and error
	FindWordByID(id uint) (*models.Word, error)
	CreateWord(word *models.Word) error
	UpdateWord(word *models.Word) error
	DeleteWord(id uint) error
	FindWordsByGroupID(groupID uint, page, pageSize int) ([]models.Word, int64, error)
	FindWordsWithFilters(filters map[string]interface{}, page, pageSize int) ([]models.Word, int64, error)
	FindWordsWithJoins(page, pageSize int) ([]models.Word, int64, error)
}

// GroupRepository defines methods for group operations
type GroupRepository interface {
	FindGroups() ([]models.Group, error)
	FindGroupByID(id uint) (*models.Group, error)
	CreateGroup(group *models.Group) error
	UpdateGroup(group *models.Group) error
	DeleteGroup(id uint) error
	FindGroupWords(groupID uint) ([]models.Word, error)
	FindGroupStudySessions(groupID uint) ([]models.StudySession, error)
}

// StudySessionRepository defines methods for study session operations
type StudySessionRepository interface {
	FindStudySessions() ([]models.StudySession, error)
	FindStudySessionByID(id uint) (*models.StudySession, error)
	CreateStudySession(session *models.StudySession) error
	UpdateStudySession(session *models.StudySession) error
	DeleteStudySession(id uint) error
	FindStudySessionWords(sessionID uint) ([]models.Word, error)
	FindStudySessionsByGroupID(groupID uint) ([]models.StudySession, error)
	FindStudySessionsByActivityID(activityID uint) ([]models.StudySession, error)
	FindLastStudySession() (*models.StudySession, error)
	FindAllWords() ([]models.Word, error)
}

// StudyActivityRepository defines methods for study activity operations
type StudyActivityRepository interface {
	FindStudyActivities() ([]models.StudyActivity, error)
	FindStudyActivityByID(id uint) (*models.StudyActivity, error)
	CreateStudyActivity(activity *models.StudyActivity) error
	UpdateStudyActivity(activity *models.StudyActivity) error
	DeleteStudyActivity(id uint) error
	FindStudyActivitySessions(activityID uint) ([]models.StudySession, error)
	FindGroupByID(id uint) (*models.Group, error)
}

// WordReviewRepository defines methods for word review operations
type WordReviewRepository interface {
	CreateWordReview(review *models.WordReviewItem) error
	FindWordReviews(sessionID uint) ([]models.WordReviewItem, error)
	FindWordReviewByID(id uint) (*models.WordReviewItem, error)
	UpdateWordReview(review *models.WordReviewItem) error
	DeleteWordReview(id uint) error
	FindWordReviewsByWordID(wordID uint) ([]models.WordReviewItem, error)
	FindAllWordReviews() ([]models.WordReviewItem, error)
}

// SettingsRepository defines methods for settings operations
type SettingsRepository interface {
	FindSettings() (*models.Settings, error)
	UpdateSettings(settings *models.Settings) error
	ResetHistory() error
	FullReset() error
}
