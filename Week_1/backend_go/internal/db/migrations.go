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

// Migration represents a single database migration
type Migration struct {
    Filename string
    SQL      string
}

// ApplyMigrations reads and applies all SQL migrations in order
func ApplyMigrations(db *sql.DB) error {
    migrationsDir := "../migrations"
    migrations, err := readMigrations(migrationsDir)
    if err != nil {
        return fmt.Errorf("failed to read migrations: %v", err)
    }

    // Sort migrations by filename (0001, 0002, etc.)
    sort.Slice(migrations, func(i, j int) bool {
        return migrations[i].Filename < migrations[j].Filename
    })

    for _, migration := range migrations {
        log.Printf("Applying migration: %s", migration.Filename)
        
        // Split SQL into individual statements
        statements := strings.Split(migration.SQL, "--")
        for _, stmt := range statements {
            stmt = strings.TrimSpace(stmt)
            if stmt == "" {
                continue
            }

            _, err := db.Exec(stmt)
            if err != nil {
                return fmt.Errorf("failed to execute migration %s: %v", migration.Filename, err)
            }
        }
    }

    return nil
}

// readMigrations reads all .sql files from the migrations directory
func readMigrations(dir string) ([]Migration, error) {
    var migrations []Migration

    // Get absolute path to migrations directory
    absDir, err := filepath.Abs(dir)
    if err != nil {
        return nil, fmt.Errorf("failed to get absolute path: %v", err)
    }

    // Read all files in directory
    files, err := os.ReadDir(absDir)
    if err != nil {
        return nil, fmt.Errorf("failed to read directory: %v", err)
    }

    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
            // Read file content
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