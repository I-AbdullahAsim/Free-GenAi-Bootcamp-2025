package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func setupStudyActivityTestRouter(mockService StudyServiceInterface) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewStudyActivityHandler(mockService)

	router.POST("/api/study-activities", handler.CreateStudyActivity)
	router.GET("/api/study-activities/:id", handler.GetStudyActivity)
	router.DELETE("/api/study-activities/:id", handler.DeleteStudyActivity)
	router.GET("/api/study-activities/:id/sessions", handler.GetStudyActivitySessions)

	return router
}

func TestCreateStudyActivity(t *testing.T) {
	mockService := new(MockStudyService)
	router := setupStudyActivityTestRouter(mockService)

	// Test case 1: Successful creation
	t.Run("Success", func(t *testing.T) {
		activity := models.StudyActivity{
			Model: gorm.Model{
				ID: 1,
			},
			StudySessionID: 1,
			GroupID:        1,
		}
		mockService.On("CreateStudyActivity", mock.AnythingOfType("*models.StudyActivity")).Return(nil)

		body, _ := json.Marshal(activity)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/study-activities", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	// Test case 2: Invalid input
	t.Run("Invalid Input", func(t *testing.T) {
		mockService.ExpectedCalls = nil // Reset mock
		mockService.On("CreateStudyActivity", mock.AnythingOfType("*models.StudyActivity")).Return(gorm.ErrInvalidData)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/study-activities", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestGetStudyActivity(t *testing.T) {
	mockService := new(MockStudyService)
	router := setupStudyActivityTestRouter(mockService)

	// Test case 1: Successful retrieval
	t.Run("Success", func(t *testing.T) {
		now := time.Now()
		expectedActivity := &models.StudyActivity{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
			StudySessionID: 1,
			GroupID:        1,
		}

		mockService.On("GetStudyActivity", uint(1)).Return(expectedActivity, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-activities/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.StudyActivity
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, uint(1), response.StudySessionID)
		assert.Equal(t, uint(1), response.GroupID)
	})

	// Test case 2: Activity not found
	t.Run("Not Found", func(t *testing.T) {
		mockService.ExpectedCalls = nil // Reset mock
		mockService.On("GetStudyActivity", uint(999)).Return(nil, gorm.ErrRecordNotFound)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-activities/999", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetStudyActivitySessions(t *testing.T) {
	mockService := new(MockStudyService)
	router := setupStudyActivityTestRouter(mockService)

	// Test case 1: Successful retrieval of sessions
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
				ActivityID: 1,
			},
		}

		mockService.On("GetStudyActivitySessions", uint(1)).Return(expectedSessions, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-activities/1/sessions", nil)
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
		mockService.On("GetStudyActivitySessions", uint(2)).Return([]models.StudySession{}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/study-activities/2/sessions", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.StudySession
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})
} 