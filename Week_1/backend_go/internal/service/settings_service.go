package service

import (
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/repository"
)

// SettingsServiceInterface allows for mocking in tests and decoupling handler from implementation
// Only the methods needed for handler are included

type SettingsServiceInterface interface {
	GetSettings() (*models.Settings, error)
	UpdateSettings(settings *models.Settings) error
	ResetHistory() error
	FullReset() error
}

type SettingsService struct {
	settingsRepo repository.SettingsRepository
}

func NewSettingsService(settingsRepo repository.SettingsRepository) *SettingsService {
	return &SettingsService{
		settingsRepo: settingsRepo,
	}
}

// GetSettings retrieves user settings
func (s *SettingsService) GetSettings() (*models.Settings, error) {
	return s.settingsRepo.FindSettings()
}

// UpdateSettings updates user settings
func (s *SettingsService) UpdateSettings(settings *models.Settings) error {
	return s.settingsRepo.UpdateSettings(settings)
}

// ResetHistory resets the user's learning history
func (s *SettingsService) ResetHistory() error {
	return s.settingsRepo.ResetHistory()
}

// FullReset performs a complete system reset
func (s *SettingsService) FullReset() error {
	return s.settingsRepo.FullReset()
}
