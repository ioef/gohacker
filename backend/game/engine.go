package game

import (
	"fmt"
	"gohacker/game/challenges"
	"gohacker/models"
	"gohacker/storage"
	"gohacker/utils"
	"strings"
	"time"
)

// Engine manages the game logic
type Engine struct {
	store      *storage.Store
	challenges []models.Challenge
}

// NewEngine creates a new game engine
func NewEngine(store *storage.Store) *Engine {
	// Load all challenges
	allChallenges := []models.Challenge{}
	allChallenges = append(allChallenges, challenges.GetBeginnerChallenges()...)
	allChallenges = append(allChallenges, challenges.GetIntermediateChallenges()...)
	allChallenges = append(allChallenges, challenges.GetAdvancedChallenges()...)

	return &Engine{
		store:      store,
		challenges: allChallenges,
	}
}

// GetAllChallenges returns all challenges
func (e *Engine) GetAllChallenges() []models.Challenge {
	return e.challenges
}

// GetChallengeByID returns a specific challenge
func (e *Engine) GetChallengeByID(id string) (*models.Challenge, error) {
	for _, c := range e.challenges {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("challenge not found")
}

// GetChallengesForUser returns challenges with user progress
func (e *Engine) GetChallengesForUser(userID string) ([]models.ChallengeResponse, error) {
	user, err := e.store.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	progress, err := e.store.GetProgress(userID)
	if err != nil {
		// Create new progress if not found
		progress = models.NewProgress(userID)
		if err := e.store.CreateProgress(progress); err != nil {
			return nil, err
		}
	}

	var responses []models.ChallengeResponse
	for _, challenge := range e.challenges {
		// Check if completed
		completed := false
		for _, completedID := range progress.CompletedChallenges {
			if completedID == challenge.ID {
				completed = true
				break
			}
		}

		// Check if locked
		locked := user.Level < challenge.RequiredLevel

		responses = append(responses, models.ChallengeResponse{
			ID:            challenge.ID,
			Title:         challenge.Title,
			Description:   challenge.Description,
			Story:         challenge.Story,
			Difficulty:    challenge.Difficulty,
			Category:      challenge.Category,
			XPReward:      challenge.XPReward,
			RequiredLevel: challenge.RequiredLevel,
			Hints:         challenge.Hints,
			StarterCode:   challenge.StarterCode,
			Order:         challenge.Order,
			Completed:     completed,
			Locked:        locked,
		})
	}

	return responses, nil
}

// ExecuteChallenge executes user code and validates it
func (e *Engine) ExecuteChallenge(userID, challengeID, code string) (*ExecutionResult, error) {
	// Get challenge
	challenge, err := e.GetChallengeByID(challengeID)
	if err != nil {
		return nil, err
	}

	// Get user
	user, err := e.store.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Check if user level is sufficient
	if user.Level < challenge.RequiredLevel {
		return nil, fmt.Errorf("challenge is locked")
	}

	// Validate code
	if err := utils.ValidateGoCode(code); err != nil {
		return &ExecutionResult{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// Execute code
	execResult, err := utils.ExecuteCode(code, 5*time.Second)
	if err != nil {
		return &ExecutionResult{
			Success: false,
			Message: err.Error(),
		}, nil
	}

	// Check if execution was successful
	if !execResult.Success {
		return &ExecutionResult{
			Success:       false,
			Message:       "Code execution failed",
			Output:        execResult.Output,
			Error:         execResult.Error,
			TestsPassed:   0,
			TotalTests:    len(challenge.TestCases),
			ExecutionTime: execResult.Time,
		}, nil
	}

	// Validate against test cases
	testsPassed := 0
	for _, testCase := range challenge.TestCases {
		if strings.TrimSpace(execResult.Output) == strings.TrimSpace(testCase.ExpectedOutput) {
			testsPassed++
		}
	}

	allTestsPassed := testsPassed == len(challenge.TestCases)

	result := &ExecutionResult{
		Success:       allTestsPassed,
		Message:       getResultMessage(allTestsPassed, testsPassed, len(challenge.TestCases)),
		Output:        execResult.Output,
		TestsPassed:   testsPassed,
		TotalTests:    len(challenge.TestCases),
		ExecutionTime: execResult.Time,
		FormattedCode: execResult.FormattedCode, // Pass formatted code to frontend
	}

	// If all tests passed, update progress
	if allTestsPassed {
		if err := e.CompleteChallenge(userID, challengeID); err != nil {
			return nil, err
		}
		result.XPGained = challenge.XPReward
		result.NewLevel = user.Level
	}

	// Update challenge progress
	e.updateChallengeProgress(userID, challengeID, code, allTestsPassed, execResult.Time)

	return result, nil
}

// CompleteChallenge marks a challenge as completed and awards XP
func (e *Engine) CompleteChallenge(userID, challengeID string) error {
	// Get challenge
	challenge, err := e.GetChallengeByID(challengeID)
	if err != nil {
		return err
	}

	// Get user
	user, err := e.store.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Get progress
	progress, err := e.store.GetProgress(userID)
	if err != nil {
		progress = models.NewProgress(userID)
	}

	// Check if already completed
	for _, id := range progress.CompletedChallenges {
		if id == challengeID {
			return nil // Already completed
		}
	}

	// Add XP
	user.AddXP(challenge.XPReward)

	// Update progress
	progress.CompleteChallenge(challengeID, challenge.Difficulty, challenge.Category)

	// Check for achievements
	e.checkAchievements(user, progress)

	// Save
	if err := e.store.UpdateUser(user); err != nil {
		return err
	}
	if err := e.store.UpdateProgress(progress); err != nil {
		return err
	}

	return nil
}

// updateChallengeProgress updates or creates challenge progress
func (e *Engine) updateChallengeProgress(userID, challengeID, code string, completed bool, execTime int) {
	cp, err := e.store.GetChallengeProgress(userID, challengeID)
	if err != nil {
		// Create new
		cp = &models.ChallengeProgress{
			UserID:      userID,
			ChallengeID: challengeID,
			Completed:   completed,
			Attempts:    1,
			LastCode:    code,
			BestTime:    execTime,
		}
		if completed {
			cp.CompletedAt = time.Now().Format(time.RFC3339)
		}
		e.store.CreateChallengeProgress(cp)
	} else {
		// Update existing
		cp.Attempts++
		cp.LastCode = code
		if completed && !cp.Completed {
			cp.Completed = true
			cp.CompletedAt = time.Now().Format(time.RFC3339)
		}
		if execTime < cp.BestTime || cp.BestTime == 0 {
			cp.BestTime = execTime
		}
		e.store.UpdateChallengeProgress(cp)
	}
}

// checkAchievements checks and unlocks achievements
func (e *Engine) checkAchievements(user *models.User, progress *models.Progress) {
	achievements := models.GetAllAchievements()

	for _, achievement := range achievements {
		// Check if already unlocked
		alreadyUnlocked := false
		for _, id := range progress.UnlockedAchievements {
			if id == achievement.ID {
				alreadyUnlocked = true
				break
			}
		}
		if alreadyUnlocked {
			continue
		}

		// Check achievement conditions
		unlocked := false
		switch achievement.ID {
		case "first_steps":
			unlocked = progress.Statistics.TotalChallengesCompleted >= 1
		case "level_10":
			unlocked = user.Level >= 10
		case "level_25":
			unlocked = user.Level >= 25
		case "level_50":
			unlocked = user.Level >= 50
		case "quick_learner":
			// Would need to track challenges completed today
			unlocked = false
		case "streak_master":
			unlocked = progress.CurrentStreak >= 7
		case "go_guru":
			unlocked = progress.Statistics.TotalChallengesCompleted >= len(e.challenges)
		}

		if unlocked {
			progress.UnlockAchievement(achievement.ID)
			user.AddXP(achievement.XPBonus)
		}
	}
}

// ExecutionResult represents the result of challenge execution
type ExecutionResult struct {
	Success       bool   `json:"success"`
	Message       string `json:"message"`
	Output        string `json:"output"`
	Error         string `json:"error,omitempty"`
	TestsPassed   int    `json:"tests_passed"`
	TotalTests    int    `json:"total_tests"`
	XPGained      int    `json:"xp_gained,omitempty"`
	NewLevel      int    `json:"new_level,omitempty"`
	ExecutionTime int    `json:"execution_time"` // milliseconds
	FormattedCode string `json:"formatted_code"` // go fmt formatted code
}

func getResultMessage(success bool, passed, total int) string {
	if success {
		messages := []string{
			"🎉 Excellent work! Challenge completed!",
			"💯 Perfect! You're a natural!",
			"🚀 Amazing! Keep up the great work!",
			"⭐ Brilliant! You've mastered this!",
			"🏆 Outstanding! You're on fire!",
		}
		return messages[passed%len(messages)]
	}
	return fmt.Sprintf("❌ Tests failed: %d/%d passed. Keep trying!", passed, total)
}
