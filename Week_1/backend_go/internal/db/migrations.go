package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Migration struct {
	Filename string
	SQL      string
}

// RunMigrations uses the global GORM DB to run raw SQL migrations
func RunMigrations() error {
	if DB == nil {
		return fmt.Errorf("DB is not initialized")
	}

	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %v", err)
	}

	return ApplyMigrations(sqlDB)
}

// ApplyMigrations reads and applies all SQL migrations in order
func ApplyMigrations(db *sql.DB) error {
	migrations, err := readMigrations(MigrationsPath)
	if err != nil {
		return fmt.Errorf("failed to read migrations: %v", err)
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Filename < migrations[j].Filename
	})

	for _, migration := range migrations {
		log.Printf("Applying migration: %s", migration.Filename)

		queries := strings.Split(migration.SQL, ";")
		for _, q := range queries {
			clean := strings.TrimSpace(q)
			if clean != "" {
				if _, err := db.Exec(clean); err != nil {
					return fmt.Errorf("failed to execute migration %s: %v", migration.Filename, err)
				}
			}
		}
	}

	return nil
}

func readMigrations(dir string) ([]Migration, error) {
	var migrations []Migration

	absDir, err := filepath.Abs(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path: %v", err)
	}

	files, err := os.ReadDir(absDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() &&
			strings.HasSuffix(file.Name(), ".sql") &&
			!strings.Contains(file.Name(), "_rollback") { // ✅ skip rollback files
			content, err := os.ReadFile(filepath.Join(absDir, file.Name()))
			if err != nil {
				return nil, fmt.Errorf("failed to read file %s: %v", file.Name(), err)
			}

			migrations = append(migrations, Migration{
				Filename: file.Name(),
				SQL:      string(content),
			})
		}
	}

	return migrations, nil
}
