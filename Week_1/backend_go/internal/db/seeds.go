package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

    "github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models" // ✅ correct

)

// SeedData loads all JSON files from /seeds and adds to DB
func SeedData() error {
	if DB == nil {
		return fmt.Errorf("DB not initialized")
	}

	files, err := os.ReadDir(SeedsPath)
	if err != nil {
		return fmt.Errorf("failed to read seeds directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			groupName := strings.TrimSuffix(file.Name(), ".json")
			err := seedGroup(filepath.Join(SeedsPath, file.Name()), groupName)
			if err != nil {
				return fmt.Errorf("failed to seed %s: %v", file.Name(), err)
			}
		}
	}

	return nil
}

// seedGroup loads words from a file and associates them with a group
func seedGroup(filepath string, groupName string) error {
	var words []models.Word

	content, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read %s: %v", filepath, err)
	}

	err = json.Unmarshal(content, &words)
	if err != nil {
		return fmt.Errorf("failed to unmarshal %s: %v", filepath, err)
	}

	// Create or find the group
	group := models.Group{Name: groupName}
	err = DB.FirstOrCreate(&group, models.Group{Name: groupName}).Error
	if err != nil {
		return fmt.Errorf("failed to create/find group %s: %v", groupName, err)
	}

	for _, word := range words {
		// Check if the word exists already
		existing := models.Word{}
		err = DB.Where("arabic_word = ? AND english_word = ?", word.ArabicWord, word.EnglishWord).FirstOrCreate(&existing, word).Error
		if err != nil {
			return fmt.Errorf("failed to create/find word: %v", err)
		}

		// Create association in word_groups
		link := models.WordGroup{
			WordID:  existing.ID,
			GroupID: group.ID,
		}
		DB.FirstOrCreate(&link, link)
	}

	return nil
}
