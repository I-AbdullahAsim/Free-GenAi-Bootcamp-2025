package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"github.com/stretchr/testify/mock"
)

func setupStudySessionTestRouter(mockService StudyServiceInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewStudySessionHandler(mockService)

	router.GET("/api/study-sessions", handler.GetStudySessions)
	router.GET("/api/study-sessions/:id", handler.GetStudySession)
	router.GET("/api/study-sessions/:id/words", handler.GetStudySessionWords)

	return router
}

func setupStudySessionTestRouterWithReview(mockService StudyServiceInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewStudySessionHandler(mockService)

	router.GET("/api/study-sessions", handler.GetStudySessions)
	router.GET("/api/study-sessions/:id", handler.GetStudySession)
	router.GET("/api/study-sessions/:id/words", handler.GetStudySessionWords)
	router.POST("/api/study-sessions/:id/words/:word_id/review", handler.CreateWordReview)

	return router
}

func TestGetStudySessions(t *testing.T) {
	mockService := new(MockStudyService)
	router := setupStudySessionTestRouter(mockService)

	// Test case 1: Successful retrieval of study sessions
	t.Run("Success", func(t *testing.T) {
		now := time.Now()
		expectedSessions := []models.StudySession{
			{
				Model: gorm.Model{
					ID:        1,
					CreatedAt: now,
					UpdatedAt: now,
				},
				GroupID:    1,
				ActivityID: 1,
			},
			{
				Model: gorm.Model{
					ID:        2,
					CreatedAt: now,
					UpdatedAt: now,
				},
				GroupID:    2,
				ActivityID: 2,
			},
		}

		mockService.On("GetAllStudySessions").Return(expectedSessions, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-sessions", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.StudySession
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
		assert.Equal(t, uint(1), response[0].GroupID)
		assert.Equal(t, uint(2), response[1].GroupID)
	})

	// Test case 2: Empty sessions list
	t.Run("Empty List", func(t *testing.T) {
		mockService.ExpectedCalls = nil // Reset mock
		mockService.On("GetAllStudySessions").Return([]models.StudySession{}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-sessions", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.StudySession
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})
}

func TestGetStudySession(t *testing.T) {
	mockService := new(MockStudyService)
	router := setupStudySessionTestRouter(mockService)

	// Test case 1: Successful retrieval of a study session
	t.Run("Success", func(t *testing.T) {
		now := time.Now()
		expectedSession := &models.StudySession{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
			GroupID:    1,
			ActivityID: 1,
		}

		mockService.On("GetStudySession", uint(1)).Return(expectedSession, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-sessions/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.StudySession
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.GroupID)
		assert.Equal(t, uint(1), response.ActivityID)
	})

	// Test case 2: Session not found
	t.Run("Not Found", func(t *testing.T) {
		mockService.ExpectedCalls = nil // Reset mock
		mockService.On("GetStudySession", uint(999)).Return(nil, gorm.ErrRecordNotFound)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-sessions/999", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetStudySessionWords(t *testing.T) {
	mockService := new(MockStudyService)
	router := setupStudySessionTestRouter(mockService)

	// Test case 1: Successful retrieval of session words
	t.Run("Success", func(t *testing.T) {
		expectedWords := []models.Word{
			{
				ID:           1,
				ArabicWord:   "مرحبا",
				EnglishWord:  "Hello",
				CorrectCount: 5,
				WrongCount:   2,
			},
			{
				ID:           2,
				ArabicWord:   "شكرا",
				EnglishWord:  "Thank you",
				CorrectCount: 3,
				WrongCount:   1,
			},
		}

		mockService.On("GetStudySessionWords", uint(1)).Return(expectedWords, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-sessions/1/words", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Word
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
		assert.Equal(t, "مرحبا", response[0].ArabicWord)
		assert.Equal(t, "شكرا", response[1].ArabicWord)
		assert.Equal(t, 5, response[0].CorrectCount)
		assert.Equal(t, 2, response[0].WrongCount)
		assert.Equal(t, 3, response[1].CorrectCount)
		assert.Equal(t, 1, response[1].WrongCount)
	})

	// Test case 2: Empty words list
	t.Run("Empty List", func(t *testing.T) {
		mockService.ExpectedCalls = nil // Reset mock
		mockService.On("GetStudySessionWords", uint(2)).Return([]models.Word{}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-sessions/2/words", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Word
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})
}

func TestCreateWordReview(t *testing.T) {
	mockService := new(MockStudyService)
	router := setupStudySessionTestRouterWithReview(mockService)

	t.Run("Success", func(t *testing.T) {
		mockService.On("CreateWordReview", mock.AnythingOfType("*models.WordReviewItem")).Return(nil)

		payload := `{"is_correct": true}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/study-sessions/2/words/1/review", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var response struct {
			Success bool `json:"success"`
			Message string `json:"message"`
			Data    models.WordReviewItem `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.True(t, response.Success)
		assert.Equal(t, "Word review updated successfully", response.Message)
		assert.Equal(t, uint(1), response.Data.WordID)
		assert.Equal(t, uint(2), response.Data.SessionID)
		assert.True(t, response.Data.IsCorrect)
	})

	t.Run("Invalid Payload", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/study-sessions/2/words/1/review", strings.NewReader("{}"))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Service Error", func(t *testing.T) {
		mockService.ExpectedCalls = nil // Reset mock
		mockService.On("CreateWordReview", mock.AnythingOfType("*models.WordReviewItem")).Return(gorm.ErrInvalidData)

		payload := `{"is_correct": false}`
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/study-sessions/2/words/1/review", strings.NewReader(payload))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
} 