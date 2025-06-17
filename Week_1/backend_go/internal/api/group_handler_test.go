package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/models"
	service "github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Remove local GroupHandler definition

// MockGroupService implements GroupServiceInterface for testing
type MockGroupService struct {
	mock.Mock
}

// Ensure MockGroupService implements GroupServiceInterface
var _ service.GroupServiceInterface = (*MockGroupService)(nil)

func (m *MockGroupService) ListGroups() ([]models.Group, error) {
	args := m.Called()
	return args.Get(0).([]models.Group), args.Error(1)
}

func (m *MockGroupService) GetGroup(id uint) (*models.Group, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Group), args.Error(1)
}

func (m *MockGroupService) CreateGroup(group *models.Group) error {
	args := m.Called(group)
	return args.Error(0)
}

func (m *MockGroupService) UpdateGroup(group *models.Group) error {
	args := m.Called(group)
	return args.Error(0)
}

func (m *MockGroupService) DeleteGroup(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockGroupService) GetGroupWords(groupID uint) ([]models.Word, error) {
	args := m.Called(groupID)
	return args.Get(0).([]models.Word), args.Error(1)
}

func (m *MockGroupService) GetGroupStudySessionsWithActivityName(groupID uint) ([]service.GroupStudySessionResponse, error) {
	args := m.Called(groupID)
	return args.Get(0).([]service.GroupStudySessionResponse), args.Error(1)
}

func setupGroupTestRouter(mockService *MockGroupService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := &GroupHandler{service: mockService}

	router.GET("/api/groups", handler.GetGroups)
	router.GET("/api/groups/:id", handler.GetGroup)
	router.GET("/api/groups/:id/words", handler.GetGroupWords)
	router.GET("/api/groups/:id/study-sessions", handler.GetGroupStudySessions)

	return router
}

func TestGetGroups(t *testing.T) {
	mockService := new(MockGroupService)
	router := setupGroupTestRouter(mockService)

	// Test case 1: Successful retrieval of groups
	t.Run("Success", func(t *testing.T) {
		mockService.ExpectedCalls = nil // Reset mock
		expectedGroups := []models.Group{
			{
				ID:    1,
				Name:  "Basic Words",
				Words: []models.Word{
					{ID: 1},
					{ID: 2},
				},
			},
			{
				ID:    2,
				Name:  "Animals",
				Words: []models.Word{
					{ID: 3},
				},
			},
		}

		mockService.On("ListGroups").Return(expectedGroups, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Group
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
		assert.Equal(t, "Basic Words", response[0].Name)
		assert.Equal(t, "Animals", response[1].Name)
		assert.Equal(t, 2, len(response[0].Words))
		assert.Equal(t, 1, len(response[1].Words))
	})

	// Test case 2: Empty groups list
	t.Run("Empty List", func(t *testing.T) {
		mockService.ExpectedCalls = nil // Reset mock
		mockService.On("ListGroups").Return([]models.Group{}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Group
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})
}

func TestGetGroup(t *testing.T) {
	mockService := new(MockGroupService)
	router := setupGroupTestRouter(mockService)

	// Test case 1: Successful retrieval of a group
	t.Run("Success", func(t *testing.T) {
		expectedGroup := &models.Group{
			ID:    1,
			Name:  "Basic Words",
			Words: []models.Word{
				{ID: 1},
				{ID: 2},
			},
		}

		mockService.On("GetGroup", uint(1)).Return(expectedGroup, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response models.Group
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Basic Words", response.Name)
		assert.Equal(t, 2, len(response.Words))
	})

	// Test case 2: Group not found
	t.Run("Not Found", func(t *testing.T) {
		mockService.On("GetGroup", uint(999)).Return(nil, gorm.ErrRecordNotFound)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups/999", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestGetGroupWords(t *testing.T) {
	mockService := new(MockGroupService)
	router := setupGroupTestRouter(mockService)

	// Test case 1: Successful retrieval of group words
	t.Run("Success", func(t *testing.T) {
		expectedWords := []models.Word{
			{
				ID:          1,
				ArabicWord:  "مرحبا",
				EnglishWord: "Hello",
			},
			{
				ID:          2,
				ArabicWord:  "شكرا",
				EnglishWord: "Thank you",
			},
		}

		mockService.On("GetGroupWords", uint(1)).Return(expectedWords, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups/1/words", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Word
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
		assert.Equal(t, "مرحبا", response[0].ArabicWord)
		assert.Equal(t, "شكرا", response[1].ArabicWord)
	})

	// Test case 2: Empty words list
	t.Run("Empty List", func(t *testing.T) {
		mockService.On("GetGroupWords", uint(2)).Return([]models.Word{}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups/2/words", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []models.Word
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})
}

func TestGetGroupStudySessions(t *testing.T) {
	mockService := new(MockGroupService)
	router := setupGroupTestRouter(mockService)

	t.Run("Success", func(t *testing.T) {
		expectedSessions := []service.GroupStudySessionResponse{
			{ID: 1, GroupID: 1, ActivityID: 1, ActivityName: "Flashcards"},
			{ID: 2, GroupID: 1, ActivityID: 2, ActivityName: "Quiz"},
		}
		mockService.On("GetGroupStudySessionsWithActivityName", uint(1)).Return(expectedSessions, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups/1/study-sessions", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []service.GroupStudySessionResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2)
		assert.Equal(t, "Flashcards", response[0].ActivityName)
		assert.Equal(t, "Quiz", response[1].ActivityName)
	})

	t.Run("Empty List", func(t *testing.T) {
		mockService.On("GetGroupStudySessionsWithActivityName", uint(2)).Return([]service.GroupStudySessionResponse{}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/groups/2/study-sessions", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []service.GroupStudySessionResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)
	})
} 