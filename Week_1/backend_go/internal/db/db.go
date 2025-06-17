package db

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
)

var (
	dbInstance    *gorm.DB
	once          sync.Once
	DB            *gorm.DB
	MigrationsPath = filepath.Join("..", "..", "migrations")
	SeedsPath     = filepath.Join("..", "..", "seeds")
)

// GetDB returns the global database instance
func GetDB() *gorm.DB {
	once.Do(func() {
		db, err := InitializeDB()
		if err != nil {
			log.Fatal("Failed to initialize database:", err)
		}
		dbInstance = db
	})
	return dbInstance
}

// InitializeDB sets up SQLite DB, applies migrations, and runs AutoMigrate
func InitializeDB() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(sqlite.Open("words.db?_journal=DELETE&_timeout=5000&_busy_timeout=5000&_foreign_keys=1&_locking=NORMAL"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Set global DB instance for migration file
	DB = db

	// Run migrations
	if err := RunMigrations(); err != nil {
		return nil, err
	}

	// Auto migrate models
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
