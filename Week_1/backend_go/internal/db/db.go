package db

import (
    "database/sql"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "time"

    "github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

// InitializeDB initializes the database connection and performs migrations
func InitializeDB() (*gorm.DB, error) {
    // Set up logger
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
        logger.Config{
            SlowThreshold: time.Second,   // Slow SQL threshold
            LogLevel:      logger.Info,   // Log level
            Colorful:      true,          // Disable color
        },
    )

    // Connect to SQLite database
    db, err := gorm.Open(sqlite.Open("words.db"), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        return nil, err
    }

    // Get raw SQL database connection
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    // Apply SQL migrations
    if err := db.ApplyMigrations(sqlDB); err != nil {
        return nil, err
    }

    // Auto-migrate models as a fallback
    if err := db.AutoMigrate(
        &models.Word{},
        &models.Group{},
        &models.WordGroup{},
        &models.StudySession{},
        &models.StudyActivity{},
        &models.WordReviewItem{},
    ); err != nil {
        return nil, err
    }

    return db, nil
}