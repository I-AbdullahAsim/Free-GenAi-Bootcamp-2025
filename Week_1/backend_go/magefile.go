package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/db"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

// InitDB initializes the database connection
func InitDB() error {
	conn, err := db.InitializeDB()
	if err != nil {
		return err
	}
	models.DB = conn
	return nil
}

// Migrate runs SQL migrations
func Migrate() error {
	mg.Deps(InitDB)

	sqlDB, err := models.DB.DB()
	if err != nil {
		log.Println(err)
		return err
	}
	defer sqlDB.Close()

	if err := models.DB.AutoMigrate(
		&models.Word{},
		&models.Group{},
		&models.StudySession{},
		&models.StudyActivity{},
		&models.WordReviewItem{},
	).Error(); err != "" {
		return fmt.Errorf("migration failed: %s", err)
	}

	log.Println("Migrations applied successfully")
	return nil
}

// Seed loads seed data into the database from JSON files
func Seed() error {
	mg.Deps(Migrate)

	sqlDB, err := models.DB.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	// Begin transaction from GORM DB
	tx := models.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Load and insert basic words
	basicWords := []models.Word{}
	if err := loadSeedData("seeds/basic_words.json", &basicWords); err != nil {
		tx.Rollback()
		return err
	}
	for _, word := range basicWords {
		if err := tx.Create(&word).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Load and insert animals
	animalWords := []models.Word{}
	if err := loadSeedData("seeds/animals.json", &animalWords); err != nil {
		tx.Rollback()
		return err
	}
	for _, word := range animalWords {
		if err := tx.Create(&word).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Insert sample groups
	groups := []models.Group{
		{Name: "basics"},
		{Name: "animals"},
	}
	for _, group := range groups {
		if err := tx.Create(&group).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	log.Println("Seed data added successfully")
	return nil
}

// loadSeedData loads JSON data from a file
func loadSeedData(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		return err
	}

	return nil
}
