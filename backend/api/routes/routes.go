package routes

import (
	"gohacker/api/handlers"
	"gohacker/api/middleware"
	"gohacker/config"
	"gohacker/game"
	"gohacker/storage"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, store *storage.Store, engine *game.Engine, cfg *config.Config) {
	// Apply CORS middleware
	router.Use(middleware.CORSMiddleware())

	// Create handlers
	authHandler := handlers.NewAuthHandler(store, cfg)
	challengeHandler := handlers.NewChallengeHandler(engine)
	progressHandler := handlers.NewProgressHandler(store)

	// Public routes
	api := router.Group("/api")
	{
		// Health check (support both GET and HEAD)
		healthHandler := func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "GoHacker API is running! 🚀",
			})
		}
		api.GET("/health", healthHandler)
		api.HEAD("/health", healthHandler)

		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
	}

	// Protected routes (require authentication)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		// User routes
		protected.GET("/auth/me", authHandler.GetMe)

		// Challenge routes
		challenges := protected.Group("/challenges")
		{
			challenges.GET("", challengeHandler.GetChallenges)
			challenges.GET("/:id", challengeHandler.GetChallenge)
			challenges.POST("/:id/execute", challengeHandler.ExecuteChallenge)
		}

		// Progress routes
		progress := protected.Group("/progress")
		{
			progress.GET("", progressHandler.GetProgress)
			progress.GET("/achievements", progressHandler.GetAchievements)
		}

		// Leaderboard
		protected.GET("/leaderboard", progressHandler.GetLeaderboard)
	}
}
