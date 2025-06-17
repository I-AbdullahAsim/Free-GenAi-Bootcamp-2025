package repository

import (
	"errors"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("record not found")
)

// WordRepositoryImpl implements the WordRepository interface
type WordRepositoryImpl struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) *WordRepositoryImpl {
	return &WordRepositoryImpl{db: db}
}

// FindWords retrieves a paginated list of words with their groups
func (r *WordRepositoryImpl) FindWords(page, pageSize int) ([]models.Word, int64, error) {
	var words []models.Word
	var total int64

	offset := (page - 1) * pageSize

	// Get total count
	if err := r.db.Model(&models.Word{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated words with their groups
	if err := r.db.
		Preload("Groups").
		Offset(offset).
		Limit(pageSize).
		Find(&words).Error; err != nil {
		return nil, 0, err
	}

	return words, total, nil
}

// FindWordByID retrieves a word by ID with its groups
func (r *WordRepositoryImpl) FindWordByID(id uint) (*models.Word, error) {
	var word models.Word
	if err := r.db.
		Preload("Groups").
		First(&word, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &word, nil
}

// CreateWord creates a new word
func (r *WordRepositoryImpl) CreateWord(word *models.Word) error {
	return r.db.Create(word).Error
}

// UpdateWord updates an existing word
func (r *WordRepositoryImpl) UpdateWord(word *models.Word) error {
	return r.db.Save(word).Error
}

// DeleteWord deletes a word by ID
func (r *WordRepositoryImpl) DeleteWord(id uint) error {
	return r.db.Delete(&models.Word{}, id).Error
}

func (r *WordRepositoryImpl) FindWordsByGroupID(groupID uint, page, pageSize int) ([]models.Word, int64, error) {
	var words []models.Word
	var total int64
	offset := (page - 1) * pageSize

	if err := r.db.Model(&models.Word{}).Joins("JOIN word_groups ON word_groups.word_id = words.id").Where("word_groups.group_id = ?", groupID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Model(&models.Word{}).Joins("JOIN word_groups ON word_groups.word_id = words.id").Where("word_groups.group_id = ?", groupID).Offset(offset).Limit(pageSize).Find(&words).Error; err != nil {
		return nil, 0, err
	}

	return words, total, nil
}

func (r *WordRepositoryImpl) FindWordsWithFilters(filters map[string]interface{}, page, pageSize int) ([]models.Word, int64, error) {
	var words []models.Word
	var total int64
	dbQuery := r.db.Model(&models.Word{})

	for key, value := range filters {
		dbQuery = dbQuery.Where(key+" = ?", value)
	}
	dbQuery.Count(&total)
	offset := (page - 1) * pageSize
	if err := dbQuery.Offset(offset).Limit(pageSize).Find(&words).Error; err != nil {
		return nil, 0, err
	}
	return words, total, nil
}

func (r *WordRepositoryImpl) FindWordsWithJoins(page, pageSize int) ([]models.Word, int64, error) {
	var words []models.Word
	var total int64
	dbQuery := r.db.Model(&models.Word{}).Preload("Groups").Preload("WordReviewItems")
	dbQuery.Count(&total)
	offset := (page - 1) * pageSize
	if err := dbQuery.Offset(offset).Limit(pageSize).Find(&words).Error; err != nil {
		return nil, 0, err
	}
	return words, total, nil
}

func init() {}