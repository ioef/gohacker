package handlers

import (
	"gohacker/models"
	"gohacker/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProgressHandler handles progress-related requests
type ProgressHandler struct {
	store *storage.Store
}

// NewProgressHandler creates a new progress handler
func NewProgressHandler(store *storage.Store) *ProgressHandler {
	return &ProgressHandler{
		store: store,
	}
}

// GetProgress returns the user's progress
func (h *ProgressHandler) GetProgress(c *gin.Context) {
	userID := c.GetString("userID")

	progress, err := h.store.GetProgress(userID)
	if err != nil {
		// Create new progress if not found
		progress = models.NewProgress(userID)
		if err := h.store.CreateProgress(progress); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create progress"})
			return
		}
	}

	c.JSON(http.StatusOK, progress)
}

// GetAchievements returns all achievements and user's unlocked ones
func (h *ProgressHandler) GetAchievements(c *gin.Context) {
	userID := c.GetString("userID")

	progress, err := h.store.GetProgress(userID)
	if err != nil {
		progress = models.NewProgress(userID)
	}

	allAchievements := models.GetAllAchievements()

	// Mark which ones are unlocked
	type AchievementResponse struct {
		models.Achievement
		Unlocked bool `json:"unlocked"`
	}

	var response []AchievementResponse
	for _, achievement := range allAchievements {
		unlocked := false
		for _, id := range progress.UnlockedAchievements {
			if id == achievement.ID {
				unlocked = true
				break
			}
		}

		// Hide hidden achievements if not unlocked
		if achievement.Hidden && !unlocked {
			continue
		}

		response = append(response, AchievementResponse{
			Achievement: achievement,
			Unlocked:    unlocked,
		})
	}

	c.JSON(http.StatusOK, gin.H{"achievements": response})
}

// GetLeaderboard returns the top users
func (h *ProgressHandler) GetLeaderboard(c *gin.Context) {
	users, err := h.store.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	// Sort by TotalXP (simple bubble sort for now)
	for i := 0; i < len(users); i++ {
		for j := i + 1; j < len(users); j++ {
			if users[j].TotalXP > users[i].TotalXP {
				users[i], users[j] = users[j], users[i]
			}
		}
	}

	// Take top 10
	if len(users) > 10 {
		users = users[:10]
	}

	// Convert to responses
	var leaderboard []models.UserResponse
	for _, user := range users {
		leaderboard = append(leaderboard, user.ToResponse())
	}

	c.JSON(http.StatusOK, gin.H{"leaderboard": leaderboard})
}
