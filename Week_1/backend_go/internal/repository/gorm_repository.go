package repository

import (
	"fmt"
	"gorm.io/gorm"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

// GORMWordRepository implements WordRepository using GORM
type GORMWordRepository struct {
	db *gorm.DB
}

func NewGORMWordRepository(db *gorm.DB) *GORMWordRepository {
	return &GORMWordRepository{db: db}
}

func (r *GORMWordRepository) FindWords(page, pageSize int) ([]models.Word, int64, error) {
	var words []models.Word
	var total int64

	// Get total count
	if err := r.db.Model(&models.Word{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if err := r.db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&words).Error; err != nil {
		return nil, 0, err
	}

	return words, total, nil
}

func (r *GORMWordRepository) FindWordsByGroupID(groupID uint, page, pageSize int) ([]models.Word, int64, error) {
	var words []models.Word
	var total int64

	// Get total count
	if err := r.db.Model(&models.Word{}).Where("group_id = ?", groupID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination and filter
	if err := r.db.Where("group_id = ?", groupID).Offset((page - 1) * pageSize).Limit(pageSize).Find(&words).Error; err != nil {
		return nil, 0, err
	}

	return words, total, nil
}

func (r *GORMWordRepository) FindWordsWithFilters(filters map[string]interface{}, page, pageSize int) ([]models.Word, int64, error) {
	var words []models.Word
	var total int64
	query := r.db.Model(&models.Word{})

	// Apply filters
	for key, value := range filters {
		query = query.Where(key, value)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if err := query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&words).Error; err != nil {
		return nil, 0, err
	}

	return words, total, nil
}

func (r *GORMWordRepository) FindWordsWithJoins(page, pageSize int) ([]models.Word, int64, error) {
	var words []models.Word
	var total int64

	// Get total count
	if err := r.db.Model(&models.Word{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply joins and pagination
	if err := r.db.Preload("Groups").Offset((page - 1) * pageSize).Limit(pageSize).Find(&words).Error; err != nil {
		return nil, 0, err
	}

	return words, total, nil
}

func (r *GORMWordRepository) FindWordByID(id uint) (*models.Word, error) {
	var word models.Word
	if err := r.db.First(&word, id).Error; err != nil {
		return nil, err
	}
	return &word, nil
}

func (r *GORMWordRepository) CreateWord(word *models.Word) error {
	return r.db.Create(word).Error
}

func (r *GORMWordRepository) UpdateWord(word *models.Word) error {
	return r.db.Save(word).Error
}

func (r *GORMWordRepository) DeleteWord(id uint) error {
	return r.db.Delete(&models.Word{}, id).Error
}



// GORMGroupRepository implements GroupRepository using GORM
type GORMGroupRepository struct {
	db *gorm.DB
}

func NewGORMGroupRepository(db *gorm.DB) *GORMGroupRepository {
	return &GORMGroupRepository{db: db}
}

func (r *GORMGroupRepository) FindGroups() ([]models.Group, error) {
	var groups []models.Group
	if err := r.db.Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (r *GORMGroupRepository) FindGroupByID(id uint) (*models.Group, error) {
	var group models.Group
	if err := r.db.First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GORMGroupRepository) CreateGroup(group *models.Group) error {
	return r.db.Create(group).Error
}

func (r *GORMGroupRepository) UpdateGroup(group *models.Group) error {
	return r.db.Save(group).Error
}

func (r *GORMGroupRepository) DeleteGroup(id uint) error {
	return r.db.Delete(&models.Group{}, id).Error
}

func (r *GORMGroupRepository) FindGroupWords(groupID uint) ([]models.Word, error) {
	var words []models.Word
	if err := r.db.Where("group_id = ?", groupID).Find(&words).Error; err != nil {
		return nil, err
	}
	return words, nil
}

func (r *GORMGroupRepository) FindGroupStudySessions(groupID uint) ([]models.StudySession, error) {
	var sessions []models.StudySession
	if err := r.db.Where("group_id = ?", groupID).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// GORMStudySessionRepository implements StudySessionRepository using GORM
type GORMStudySessionRepository struct {
	db *gorm.DB
}

func NewGORMStudySessionRepository(db *gorm.DB) *GORMStudySessionRepository {
	return &GORMStudySessionRepository{db: db}
}

func (r *GORMStudySessionRepository) FindStudySessionsByActivityID(activityID uint) ([]models.StudySession, error) {
	var sessions []models.StudySession
	if err := r.db.Where("activity_id = ?", activityID).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *GORMStudySessionRepository) FindStudySessions() ([]models.StudySession, error) {
	var sessions []models.StudySession
	if err := r.db.Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

func (r *GORMStudySessionRepository) FindLastStudySession() (*models.StudySession, error) {
	var session models.StudySession
	if err := r.db.Order("created_at DESC").First(&session).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *GORMStudySessionRepository) FindAllWords() ([]models.Word, error) {
	var words []models.Word
	if err := r.db.Find(&words).Error; err != nil {
		return nil, err
	}
	return words, nil
}

func (r *GORMStudySessionRepository) FindStudySessionByID(id uint) (*models.StudySession, error) {
	var session models.StudySession
	if err := r.db.Preload("Words").First(&session, id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *GORMStudySessionRepository) CreateStudySession(session *models.StudySession) error {
	return r.db.Create(session).Error
}

func (r *GORMStudySessionRepository) UpdateStudySession(session *models.StudySession) error {
	return r.db.Save(session).Error
}

func (r *GORMStudySessionRepository) DeleteStudySession(id uint) error {
	return r.db.Delete(&models.StudySession{}, id).Error
}

func (r *GORMStudySessionRepository) FindStudySessionWords(sessionID uint) ([]models.Word, error) {
	var words []models.Word
	if result := r.db.Model(&models.StudySession{}).Where("id = ?", sessionID).Association("Words").Find(&words); result.Error != nil {
		return nil, fmt.Errorf("failed to find words for session %d: %v", sessionID, result.Error)
	}
	return words, nil
}

func (r *GORMStudySessionRepository) FindStudySessionsByGroupID(groupID uint) ([]models.StudySession, error) {
	var sessions []models.StudySession
	if err := r.db.Where("group_id = ?", groupID).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// GORMStudyActivityRepository implements StudyActivityRepository using GORM
type GORMStudyActivityRepository struct {
	db *gorm.DB
}

func NewGORMStudyActivityRepository(db *gorm.DB) *GORMStudyActivityRepository {
	return &GORMStudyActivityRepository{db: db}
}

func (r *GORMStudyActivityRepository) FindStudyActivities() ([]models.StudyActivity, error) {
	var activities []models.StudyActivity
	if err := r.db.Find(&activities).Error; err != nil {
		return nil, err
	}
	return activities, nil
}

func (r *GORMStudyActivityRepository) FindStudyActivityByID(id uint) (*models.StudyActivity, error) {
	var activity models.StudyActivity
	if err := r.db.First(&activity, id).Error; err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *GORMStudyActivityRepository) CreateStudyActivity(activity *models.StudyActivity) error {
	return r.db.Create(activity).Error
}

func (r *GORMStudyActivityRepository) UpdateStudyActivity(activity *models.StudyActivity) error {
	return r.db.Save(activity).Error
}

func (r *GORMStudyActivityRepository) DeleteStudyActivity(id uint) error {
	return r.db.Delete(&models.StudyActivity{}, id).Error
}

func (r *GORMStudyActivityRepository) FindStudyActivitySessions(activityID uint) ([]models.StudySession, error) {
	var sessions []models.StudySession
	if err := r.db.Where("study_activity_id = ?", activityID).Find(&sessions).Error; err != nil {
		return nil, err
	}
	return sessions, nil
}

// GORMWordReviewRepository implements WordReviewRepository using GORM
type GORMWordReviewRepository struct {
	db *gorm.DB
}

func NewGORMWordReviewRepository(db *gorm.DB) *GORMWordReviewRepository {
	return &GORMWordReviewRepository{db: db}
}

func (r *GORMWordReviewRepository) CreateWordReview(review *models.WordReviewItem) error {
	return r.db.Create(review).Error
}

func (r *GORMWordReviewRepository) FindWordReviews(sessionID uint) ([]models.WordReviewItem, error) {
	var reviews []models.WordReviewItem
	if err := r.db.Where("study_session_id = ?", sessionID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *GORMWordReviewRepository) FindAllWordReviews() ([]models.WordReviewItem, error) {
	var reviews []models.WordReviewItem
	if err := r.db.Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *GORMWordReviewRepository) FindWordReviewByID(id uint) (*models.WordReviewItem, error) {
	var review models.WordReviewItem
	if err := r.db.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *GORMWordReviewRepository) UpdateWordReview(review *models.WordReviewItem) error {
	return r.db.Save(review).Error
}

func (r *GORMWordReviewRepository) DeleteWordReview(id uint) error {
	return r.db.Delete(&models.WordReviewItem{}, id).Error
}

func (r *GORMWordReviewRepository) FindWordReviewsByWordID(wordID uint) ([]models.WordReviewItem, error) {
	var reviews []models.WordReviewItem
	if err := r.db.Where("word_id = ?", wordID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

// GORMSettingsRepository implements SettingsRepository using GORM
type GORMSettingsRepository struct {
	db *gorm.DB
}

func NewGORMSettingsRepository(db *gorm.DB) *GORMSettingsRepository {
	return &GORMSettingsRepository{db: db}
}

func (r *GORMSettingsRepository) FindSettings() (*models.Settings, error) {
	var settings models.Settings
	if err := r.db.First(&settings).Error; err != nil {
		return nil, err
	}
	return &settings, nil
}

func (r *GORMSettingsRepository) UpdateSettings(settings *models.Settings) error {
	return r.db.Save(settings).Error
}

func (r *GORMSettingsRepository) ResetHistory() error {
	if err := r.db.Exec("DELETE FROM word_review_items").Error; err != nil {
		return err
	}
	if err := r.db.Exec("DELETE FROM study_sessions").Error; err != nil {
		return err
	}
	return nil
}

func (r *GORMSettingsRepository) FullReset() error {
	if err := r.db.Exec("DELETE FROM word_review_items").Error; err != nil {
		return err
	}
	if err := r.db.Exec("DELETE FROM study_sessions").Error; err != nil {
		return err
	}
	if err := r.db.Exec("DELETE FROM study_activities").Error; err != nil {
		return err
	}
	return nil
}
