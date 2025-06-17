package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type StringSlice []string

func (s StringSlice) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	return json.Marshal(s)
}

func (s *StringSlice) Scan(value interface{}) error {
	if value == nil {
		*s = StringSlice{}
		return nil
	}

	var bytes []byte
	switch v := value.(type) {
	case []byte:
		bytes = v
	case string:
		bytes = []byte(v)
	default:
		return errors.New("type assertion to []byte or string failed")
	}

	return json.Unmarshal(bytes, s)
}

type Word struct {
	ID              uint             `gorm:"primaryKey" json:"id"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DeletedAt       gorm.DeletedAt   `gorm:"index" json:"deleted_at"`
	ArabicWord      string           `gorm:"not null" json:"arabic_word"`
	EnglishWord     string           `gorm:"not null" json:"english_word"`
	Parts           StringSlice      `gorm:"type:text" json:"parts"`
	Groups          []Group          `gorm:"many2many:word_groups;" json:"groups,omitempty"`
	WordReviewItems []WordReviewItem `gorm:"foreignKey:WordID" json:"word_review_items,omitempty"`
	CorrectCount    int              `gorm:"default:0" json:"correct_count"`
	WrongCount      int              `gorm:"default:0" json:"wrong_count"`
	SuccessRate     float64          `gorm:"-" json:"success_rate"` // Virtual field, calculated on the fly
}
