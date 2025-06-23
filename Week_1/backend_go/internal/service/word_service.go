package service

import (
	"errors"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/repository"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

var (
	ErrWordNotFound = errors.New("word not found")
)

// WordServiceInterface allows for mocking in tests and decoupling handler from implementation
// Only the methods needed for handler are included

type WordServiceInterface interface {
	ListWords(page, pageSize int) ([]models.Word, int64, error)
	GetWord(id uint) (*models.Word, error)
	CreateWord(word *models.Word) error
	UpdateWord(word *models.Word) error
	DeleteWord(id uint) error
	AddWordToGroup(wordID uint, groupID uint) error
	RemoveWordFromGroup(wordID uint, groupID uint) error
}

type WordService struct {
	wordRepo repository.WordRepository
	groupRepo repository.GroupRepository
}

func NewWordService(wordRepo repository.WordRepository, groupRepo repository.GroupRepository) *WordService {
	return &WordService{
		wordRepo: wordRepo,
		groupRepo: groupRepo,
	}
}

// ListWords retrieves all words with pagination
func (s *WordService) ListWords(page, pageSize int) ([]models.Word, int64, error) {
	words, total, err := s.wordRepo.FindWords(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	// Calculate success rate for each word
	for i := range words {
		words[i].SuccessRate = calculateSuccessRate(words[i].CorrectCount, words[i].WrongCount)
	}

	return words, total, nil
}

// GetWord retrieves a word by ID
func (s *WordService) GetWord(id uint) (*models.Word, error) {
	word, err := s.wordRepo.FindWordByID(id)
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, ErrWordNotFound
		}
		return nil, err
	}

	// Calculate success rate
	word.SuccessRate = calculateSuccessRate(word.CorrectCount, word.WrongCount)

	return word, nil
}

// calculateSuccessRate calculates the success rate based on correct and wrong counts
func calculateSuccessRate(correct, wrong int) float64 {
	total := correct + wrong
	if total == 0 {
		return 0
	}
	return float64(correct) / float64(total)
}

// CreateWord creates a new word
func (s *WordService) CreateWord(word *models.Word) error {
	return s.wordRepo.CreateWord(word)
}

// UpdateWord updates an existing word
func (s *WordService) UpdateWord(word *models.Word) error {
	return s.wordRepo.UpdateWord(word)
}

// DeleteWord deletes a word by ID
func (s *WordService) DeleteWord(id uint) error {
	return s.wordRepo.DeleteWord(id)
}

// AddWordToGroup adds a word to a group
func (s *WordService) AddWordToGroup(wordID uint, groupID uint) error {
	word, err := s.wordRepo.FindWordByID(wordID)
	if err != nil {
		return err
	}
	
	group, err := s.groupRepo.FindGroupByID(groupID)
	if err != nil {
		return err
	}
	
	word.Groups = append(word.Groups, *group)
	return s.wordRepo.UpdateWord(word)
}

// RemoveWordFromGroup removes a word from a group
func (s *WordService) RemoveWordFromGroup(wordID uint, groupID uint) error {
	word, err := s.wordRepo.FindWordByID(wordID)
	if err != nil {
		return err
	}
	
	newGroups := make([]models.Group, 0)
	for _, g := range word.Groups {
		if g.ID != groupID {
			newGroups = append(newGroups, g)
		}
	}
	word.Groups = newGroups
	return s.wordRepo.UpdateWord(word)
}