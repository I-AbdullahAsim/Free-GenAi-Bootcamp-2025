package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/gorm"

	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

// SeedDatabase populates the database with initial data from JSON files
func SeedDatabase(db *gorm.DB) error {
	// Get the absolute path to the seeds directory
	seedsDir := filepath.Join("..", "..", "seeds")

	// Read all JSON files in the seeds directory
	files, err := os.ReadDir(seedsDir)
	if err != nil {
		return fmt.Errorf("failed to read seeds directory: %v", err)
	}

	var seededGroups []models.Group

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(seedsDir, file.Name())

			// Create a group with the same name as the file (without extension)
			groupName := file.Name()
			groupName = groupName[:len(groupName)-len(filepath.Ext(groupName))]

			var group models.Group
			group.Name = groupName

			// Create the group if it doesn't exist
			if err := db.FirstOrCreate(&group, models.Group{Name: groupName}).Error; err != nil {
				return fmt.Errorf("failed to create group %s: %v", groupName, err)
			}
			seededGroups = append(seededGroups, group)

			// Read JSON file
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %v", filePath, err)
			}

			// Parse JSON into slice of word maps
			var words []map[string]interface{}
			if err := json.Unmarshal(fileData, &words); err != nil {
				return fmt.Errorf("failed to parse JSON %s: %v", filePath, err)
			}

			// Create words and associate with group
			for _, wordData := range words {
				arabicWord, ok := wordData["arabic_word"].(string)
				if !ok {
					return fmt.Errorf("missing arabic_word in %s", filePath)
				}
				enWord, ok := wordData["english_word"].(string)
				if !ok {
					return fmt.Errorf("missing english_word in %s", filePath)
				}

				// Handle parts as string array
				partsInterface, ok := wordData["parts"].([]interface{})
				if !ok {
					return fmt.Errorf("missing or invalid parts in %s", filePath)
				}

				// Convert interface slice to string slice
				parts := make(models.StringSlice, len(partsInterface))
				for i, part := range partsInterface {
					partStr, ok := part.(string)
					if !ok {
						return fmt.Errorf("invalid part type in %s", filePath)
					}
					parts[i] = partStr
				}

				word := models.Word{
					ArabicWord:  arabicWord,
					EnglishWord: enWord,
					Parts:       parts,
				}

				// Create word if it doesn't exist
				if err := db.FirstOrCreate(&word, models.Word{
					ArabicWord:  arabicWord,
					EnglishWord: enWord,
				}).Error; err != nil {
					return fmt.Errorf("failed to create word %s: %v", arabicWord, err)
				}

				// Associate word with group
				if err := db.Model(&group).Association("Words").Append(&word); err != nil {
					return fmt.Errorf("failed to associate word %s with group %s: %v",
						arabicWord, groupName, err)
				}
			}
		}
	}

	// Log seeded groups
	log.Printf("Successfully seeded %d groups: %v", len(seededGroups),
		func() []string {
			var names []string
			for _, g := range seededGroups {
				names = append(names, g.Name)
			}
			return names
		}())

	return nil
}
