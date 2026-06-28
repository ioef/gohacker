package models

import "time"

// Progress represents a user's overall game progress
type Progress struct {
	UserID               string     `json:"user_id"`
	CompletedChallenges  []string   `json:"completed_challenges"`
	UnlockedAchievements []string   `json:"unlocked_achievements"`
	CurrentStreak        int        `json:"current_streak"`
	LongestStreak        int        `json:"longest_streak"`
	LastPlayedDate       time.Time  `json:"last_played_date"`
	TotalPlayTime        int        `json:"total_play_time"` // in seconds
	Statistics           Statistics `json:"statistics"`
}

// Statistics holds detailed user statistics
type Statistics struct {
	TotalChallengesCompleted int            `json:"total_challenges_completed"`
	ChallengesByDifficulty   map[string]int `json:"challenges_by_difficulty"`
	ChallengesByCategory     map[string]int `json:"challenges_by_category"`
	TotalHintsUsed           int            `json:"total_hints_used"`
	AverageAttempts          float64        `json:"average_attempts"`
	FastestCompletion        int            `json:"fastest_completion"` // in seconds
}

// NewProgress creates a new Progress instance for a user
func NewProgress(userID string) *Progress {
	return &Progress{
		UserID:               userID,
		CompletedChallenges:  []string{},
		UnlockedAchievements: []string{},
		CurrentStreak:        0,
		LongestStreak:        0,
		LastPlayedDate:       time.Now(),
		TotalPlayTime:        0,
		Statistics: Statistics{
			TotalChallengesCompleted: 0,
			ChallengesByDifficulty:   make(map[string]int),
			ChallengesByCategory:     make(map[string]int),
			TotalHintsUsed:           0,
			AverageAttempts:          0,
			FastestCompletion:        0,
		},
	}
}

// CompleteChallenge marks a challenge as completed
func (p *Progress) CompleteChallenge(challengeID string, difficulty Difficulty, category Category) {
	// Add to completed challenges if not already there
	for _, id := range p.CompletedChallenges {
		if id == challengeID {
			return // Already completed
		}
	}

	p.CompletedChallenges = append(p.CompletedChallenges, challengeID)
	p.Statistics.TotalChallengesCompleted++
	p.Statistics.ChallengesByDifficulty[string(difficulty)]++
	p.Statistics.ChallengesByCategory[string(category)]++
	p.UpdateStreak()
}

// UpdateStreak updates the user's daily streak
func (p *Progress) UpdateStreak() {
	now := time.Now()
	lastPlayed := p.LastPlayedDate

	// Check if it's a new day
	if now.YearDay() != lastPlayed.YearDay() || now.Year() != lastPlayed.Year() {
		// Check if it's consecutive days
		daysDiff := int(now.Sub(lastPlayed).Hours() / 24)
		if daysDiff == 1 {
			p.CurrentStreak++
		} else if daysDiff > 1 {
			p.CurrentStreak = 1
		}

		if p.CurrentStreak > p.LongestStreak {
			p.LongestStreak = p.CurrentStreak
		}
	}

	p.LastPlayedDate = now
}

// UnlockAchievement adds an achievement to the user's collection
func (p *Progress) UnlockAchievement(achievementID string) bool {
	// Check if already unlocked
	for _, id := range p.UnlockedAchievements {
		if id == achievementID {
			return false
		}
	}

	p.UnlockedAchievements = append(p.UnlockedAchievements, achievementID)
	return true
}
