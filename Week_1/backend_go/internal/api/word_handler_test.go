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
	"errors"
	"github.com/stretchr/testify/mock"
)

var ErrWordNotFound = errors.New("word not found")

type MockWordService struct {
	mock.Mock
}

func (m *MockWordService) ListWords(page, pageSize int) ([]models.Word, int64, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]models.Word), args.Get(1).(int64), args.Error(2)
}

func (m *MockWordService) GetWord(id uint) (*models.Word, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Word), args.Error(1)
}

func (m *MockWordService) AddWordToGroup(wordID uint, groupID uint) error {
	args := m.Called(wordID, groupID)
	return args.Error(0)
}

func (m *MockWordService) RemoveWordFromGroup(wordID uint, groupID uint) error {
	args := m.Called(wordID, groupID)
	return args.Error(0)
}

func (m *MockWordService) CreateWord(word *models.Word) error {
	args := m.Called(word)
	return args.Error(0)
}

func (m *MockWordService) UpdateWord(word *models.Word) error {
	args := m.Called(word)
	return args.Error(0)
}

func (m *MockWordService) DeleteWord(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

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

func TestWordsAPI_JSONFormatAndEdgeCases(t *testing.T) {
	mockService := new(MockWordService)
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewWordHandler(mockService)
	router.GET("/api/words", handler.GetWords)
	router.GET("/api/words/:id", handler.GetWord)

	t.Run("GET /api/words returns JSON with required fields", func(t *testing.T) {
		mockService.On("ListWords", 1, 100).Return([]models.Word{}, int64(0), nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp, "success")
		assert.Contains(t, resp, "message")
		assert.Contains(t, resp, "data")
		data := resp["data"].(map[string]interface{})
		assert.Contains(t, data, "items")
	})

	t.Run("GET /api/words/:id invalid returns 404 with JSON error", func(t *testing.T) {
		mockService.On("GetWord", uint(999)).Return(nil, service.ErrWordNotFound)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words/999", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		var resp map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp, "success")
		assert.Contains(t, resp, "message")
	})

	t.Run("GET /api/words?page=9999 returns empty list", func(t *testing.T) {
		mockService.On("ListWords", 9999, 100).Return([]models.Word{}, int64(0), nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words?page=9999", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var resp map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp, "data")
		data := resp["data"].(map[string]interface{})
		assert.Contains(t, data, "items")
		items := data["items"].([]interface{})
		assert.Empty(t, items)
	})

	t.Run("Malformed request returns 400 with JSON error", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/words/abc", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		var resp map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Contains(t, resp, "success")
		assert.Contains(t, resp, "message")
	})
} 