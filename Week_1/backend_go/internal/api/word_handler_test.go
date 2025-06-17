package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/service"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/repository"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupWordRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// In-memory SQLite DB for testing
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate models
	db.AutoMigrate(&models.Word{}, &models.Group{}, &models.WordReviewItem{}, &models.StudySession{}, &models.StudyActivity{})

	// Seed with a sample group and word
	group := models.Group{Name: "Sample Group"}
	db.Create(&group)
	word := models.Word{ArabicWord: "كلمة", EnglishWord: "word", CorrectCount: 5, WrongCount: 2, Groups: []models.Group{group}}
	db.Create(&word)

	wordRepo := repository.NewWordRepository(db)
	groupRepo := repository.NewGroupRepository(db)
	wordService := service.NewWordService(wordRepo, groupRepo)
	wordHandler := NewWordHandler(wordService)
	
	// Routes
	router.GET("/api/words", wordHandler.GetWords)
	router.GET("/api/words/:id", wordHandler.GetWord)
	
	return router
}

func TestGetWords(t *testing.T) {
	router := setupWordRouter()
	
	tests := []struct {
		name           string
		page          string
		expectedCode  int
		expectedItems int
	}{
		{
			name:           "Get first page",
			page:          "1",
			expectedCode:  http.StatusOK,
			expectedItems: 1, // Only 1 word is seeded
		},
		{
			name:           "Empty page",
			page:          "999",
			expectedCode:  http.StatusOK,
			expectedItems: 0,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/words?page="+tt.page, nil)
			router.ServeHTTP(w, req)
			
			assert.Equal(t, tt.expectedCode, w.Code)
			
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			
			// Check response structure
			assert.True(t, response["success"].(bool))
			assert.Contains(t, response, "data")
			
			dataMap := response["data"].(map[string]interface{})
			items := dataMap["items"].([]interface{})
			assert.Len(t, items, tt.expectedItems)
			
			// Check word structure
			if len(items) > 0 {
				word := items[0].(map[string]interface{})
				assert.Contains(t, word, "id")
				assert.Contains(t, word, "arabic_word")
				assert.Contains(t, word, "english_word")
				assert.Contains(t, word, "correct_count")
				assert.Contains(t, word, "wrong_count")
				assert.Contains(t, word, "success_rate")
				assert.Contains(t, word, "groups")
			}
		})
	}
}

func TestGetWord(t *testing.T) {
	router := setupWordRouter()
	
	tests := []struct {
		name         string
		id           string
		expectedCode int
	}{
		{
			name:         "Get existing word",
			id:           "1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Get non-existent word",
			id:           "999",
			expectedCode: http.StatusNotFound,
		},
		{
			name:         "Invalid ID format",
			id:           "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/api/words/"+tt.id, nil)
			router.ServeHTTP(w, req)
			
			assert.Equal(t, tt.expectedCode, w.Code)
			
			if tt.expectedCode == http.StatusOK {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				
				// Check response structure
				assert.True(t, response["success"].(bool))
				assert.Contains(t, response, "data")
				
				word := response["data"].(map[string]interface{})
				assert.Contains(t, word, "id")
				assert.Contains(t, word, "arabic_word")
				assert.Contains(t, word, "english_word")
				assert.Contains(t, word, "correct_count")
				assert.Contains(t, word, "wrong_count")
				assert.Contains(t, word, "success_rate")
				assert.Contains(t, word, "groups")
			}
		})
	}
} 