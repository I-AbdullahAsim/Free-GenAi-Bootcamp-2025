package db_test

import (
	"os"
	"path/filepath"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	appdb "github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/db"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

// Setup a temporary test database
func setupTestDB(t *testing.T) *gorm.DB {
	testDBFile := "test_words.db"
	_ = os.Remove(testDBFile)

	db, err := gorm.Open(sqlite.Open(testDBFile), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// override app DB
	appdb.DB = db

	// Set the correct migrations path
	appdb.MigrationsPath = filepath.Join("..", "..", "migrations")
	appdb.SeedsPath = filepath.Join("..", "..", "seeds")

	return db
}

// Utility to check if a table exists
func tableExists(db *gorm.DB, tableName string) bool {
	var count int64
	db.Raw("SELECT count(*) FROM sqlite_master WHERE type='table' AND name = ?", tableName).Scan(&count)
	return count > 0
}

// Test that all migrations run successfully and tables exist
func Test_DB_Migration_And_Schema_Creation(t *testing.T) {
	db := setupTestDB(t)

	err := appdb.RunMigrations()
	if err != nil {
		t.Fatalf("Migration failed: %v", err)
	}

	expectedTables := []string{
		"words", "groups", "word_groups",
		"study_sessions", "study_activities", "word_review_items",
	}

	for _, table := range expectedTables {
		if !tableExists(db, table) {
			t.Errorf("Expected table %s to exist but it does not", table)
		}
	}
}

// Test that seed data populates the words and groups
func Test_Seed_Data_Import(t *testing.T) {
	db := setupTestDB(t)

	err := appdb.RunMigrations()
	if err != nil {
		t.Fatalf("Migration failed: %v", err)
	}

	err = appdb.SeedData()
	if err != nil {
		t.Fatalf("Seeding failed: %v", err)
	}

	var wordCount, groupCount int64
	db.Model(&models.Word{}).Count(&wordCount)
	db.Model(&models.Group{}).Count(&groupCount)

	if wordCount == 0 {
		t.Error("Expected some words to be seeded, but found none")
	}

	if groupCount == 0 {
		t.Error("Expected some groups to be seeded, but found none")
	}
}

// Test that word_group mappings are created
func Test_Seed_Group_Mapping(t *testing.T) {
	db := setupTestDB(t)

	err := appdb.RunMigrations()
	if err != nil {
		t.Fatalf("Migration failed: %v", err)
	}

	err = appdb.SeedData()
	if err != nil {
		t.Fatalf("Seeding failed: %v", err)
	}

	var count int64
	db.Model(&models.WordGroup{}).Count(&count)

	if count == 0 {
		t.Error("Expected word_groups mappings, but none found")
	}
}
