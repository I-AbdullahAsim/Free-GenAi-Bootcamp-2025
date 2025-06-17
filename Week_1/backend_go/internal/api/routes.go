package api

import (
	"github.com/gin-gonic/gin"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/service"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/db"
	"github.com/i-AbdullahAsim/free-genai-bootcamp-2025/Week_1/backend_go/internal/repository"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine) {
	// Initialize repositories
	db := db.GetDB()
	wordRepo := repository.NewGORMWordRepository(db)
	groupRepo := repository.NewGORMGroupRepository(db)
	sessionRepo := repository.NewGORMStudySessionRepository(db)
	activityRepo := repository.NewGORMStudyActivityRepository(db)
	reviewRepo := repository.NewGORMWordReviewRepository(db)
	settingsRepo := repository.NewGORMSettingsRepository(db)

	// Ensure all repository interfaces are implemented
	var _ repository.WordRepository = wordRepo
	var _ repository.GroupRepository = groupRepo
	var _ repository.StudySessionRepository = sessionRepo
	var _ repository.StudyActivityRepository = activityRepo
	var _ repository.WordReviewRepository = reviewRepo
	var _ repository.SettingsRepository = settingsRepo

	// Initialize services
	wordService := service.NewWordService(wordRepo, groupRepo)
	studyService := service.NewStudyService(sessionRepo, activityRepo, reviewRepo)
	groupService := service.NewGroupService(db)
	settingsService := service.NewSettingsService(settingsRepo)

	// Initialize handlers
	wordHandler := NewWordHandler(wordService)
	groupHandler := NewGroupHandler(groupService)
	dashboardHandler := NewDashboardHandler(studyService)
	settingsHandler := NewSettingsHandler(settingsService)
	studyActivityHandler := NewStudyActivityHandler(studyService)
	studySessionHandler := NewStudySessionHandler(studyService)

	api := router.Group("/api")
	{
		// Word endpoints
		api.GET("/words", wordHandler.GetWords)
		api.POST("/words", wordHandler.CreateWord)
		api.GET("/words/:id", wordHandler.GetWord)
		api.PUT("/words/:id", wordHandler.UpdateWord)
		api.DELETE("/words/:id", wordHandler.DeleteWord)

		// Group endpoints
		api.GET("/groups", groupHandler.GetGroups)
		api.POST("/groups", groupHandler.CreateGroup)
		api.GET("/groups/:id", groupHandler.GetGroup)
		api.PUT("/groups/:id", groupHandler.UpdateGroup)
		api.DELETE("/groups/:id", groupHandler.DeleteGroup)
		api.GET("/groups/:id/words", groupHandler.GetGroupWords)
		api.GET("/groups/:id/study-sessions", groupHandler.GetGroupStudySessions)

		// Study Activity endpoints
		api.GET("/study-activities/:id", studyActivityHandler.GetStudyActivity)
		api.GET("/study-activities/:id/study-sessions", studyActivityHandler.GetStudyActivitySessions)
		api.POST("/study-activities", studyActivityHandler.CreateStudyActivity)

		// Study Session endpoints
		api.GET("/study-sessions", studySessionHandler.GetStudySessions)
		api.GET("/study-sessions/:id", studySessionHandler.GetStudySession)
		api.GET("/study-sessions/:id/words", studySessionHandler.GetStudySessionWords)
		api.POST("/study-sessions/:id/words/:word_id/review", studySessionHandler.CreateWordReview)
		api.DELETE("/study-sessions/:id", studySessionHandler.DeleteStudySession)

		// Settings endpoints
		api.POST("/settings/reset-history", settingsHandler.ResetHistory)
		api.POST("/settings/full-reset", settingsHandler.FullReset)
		api.POST("/study-sessions", studySessionHandler.CreateStudySession)

		// Dashboard endpoints
		api.GET("/dashboard/last_study_session", dashboardHandler.GetLastStudySession)
		api.GET("/dashboard/study_progress", dashboardHandler.GetStudyProgress)
		api.GET("/dashboard/quick_stats", dashboardHandler.GetQuickStats)
	}
}
