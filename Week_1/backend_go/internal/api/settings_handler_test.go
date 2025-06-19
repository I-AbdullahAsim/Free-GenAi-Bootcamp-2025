package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupSettingsTestRouter(mockService *MockSettingsService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewSettingsHandler(mockService)

	router.POST("/api/settings/reset-history", handler.ResetHistory)
	router.POST("/api/settings/full-reset", handler.FullReset)
	return router
}

func TestSettingsEndpoints(t *testing.T) {
	mockService := new(MockSettingsService)
	router := setupSettingsTestRouter(mockService)

	t.Run("ResetHistory Success", func(t *testing.T) {
		mockService.On("ResetHistory").Return(nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/settings/reset-history", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("ResetHistory Error", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.On("ResetHistory").Return(assert.AnError)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/settings/reset-history", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("FullReset Success", func(t *testing.T) {
		mockService.On("FullReset").Return(nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/settings/full-reset", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})

	t.Run("FullReset Error", func(t *testing.T) {
		mockService.ExpectedCalls = nil
		mockService.On("FullReset").Return(assert.AnError)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/settings/full-reset", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
} 