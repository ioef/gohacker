package handlers

import (
	"gohacker/game"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ChallengeHandler handles challenge-related requests
type ChallengeHandler struct {
	engine *game.Engine
}

// NewChallengeHandler creates a new challenge handler
func NewChallengeHandler(engine *game.Engine) *ChallengeHandler {
	return &ChallengeHandler{
		engine: engine,
	}
}

// GetChallenges returns all challenges for the current user
func (h *ChallengeHandler) GetChallenges(c *gin.Context) {
	userID := c.GetString("userID")

	challenges, err := h.engine.GetChallengesForUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"challenges": challenges})
}

// GetChallenge returns a specific challenge
func (h *ChallengeHandler) GetChallenge(c *gin.Context) {
	challengeID := c.Param("id")
	userID := c.GetString("userID")

	challenge, err := h.engine.GetChallengeByID(challengeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
		return
	}

	// Get user's challenges to check if locked/completed
	challenges, err := h.engine.GetChallengesForUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Find the challenge response
	for _, ch := range challenges {
		if ch.ID == challengeID {
			c.JSON(http.StatusOK, ch)
			return
		}
	}

	// Fallback (shouldn't happen)
	c.JSON(http.StatusOK, challenge)
}

// ExecuteRequest represents a code execution request
type ExecuteRequest struct {
	Code string `json:"code" binding:"required"`
}

// ExecuteChallenge executes user code for a challenge
func (h *ChallengeHandler) ExecuteChallenge(c *gin.Context) {
	challengeID := c.Param("id")
	userID := c.GetString("userID")

	var req ExecuteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.engine.ExecuteChallenge(userID, challengeID, req.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
