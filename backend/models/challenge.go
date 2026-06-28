package models

// Difficulty represents challenge difficulty levels
type Difficulty string

const (
	DifficultyNewbie       Difficulty = "newbie"
	DifficultyIntermediate Difficulty = "intermediate"
	DifficultyAdvanced     Difficulty = "advanced"
	DifficultyExpert       Difficulty = "expert"
)

// Category represents challenge categories
type Category string

const (
	CategoryBasics      Category = "basics"
	CategoryStructs     Category = "structs"
	CategoryConcurrency Category = "concurrency"
	CategoryWeb         Category = "web"
	CategoryTesting     Category = "testing"
)

// Challenge represents a coding challenge
type Challenge struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Story         string     `json:"story"`
	Difficulty    Difficulty `json:"difficulty"`
	Category      Category   `json:"category"`
	XPReward      int        `json:"xp_reward"`
	RequiredLevel int        `json:"required_level"`
	Hints         []string   `json:"hints"`
	StarterCode   string     `json:"starter_code"`
	Solution      string     `json:"-"` // Hidden from client
	TestCases     []TestCase `json:"-"` // Hidden from client
	Order         int        `json:"order"`
}

// TestCase represents a test case for validating solutions
type TestCase struct {
	Input          string `json:"input"`
	ExpectedOutput string `json:"expected_output"`
	Description    string `json:"description"`
	Hidden         bool   `json:"hidden"` // Hidden test cases for anti-cheat
}

// ChallengeProgress represents a user's progress on a challenge
type ChallengeProgress struct {
	ChallengeID string `json:"challenge_id"`
	UserID      string `json:"user_id"`
	Completed   bool   `json:"completed"`
	Attempts    int    `json:"attempts"`
	HintsUsed   int    `json:"hints_used"`
	LastCode    string `json:"last_code"`
	BestTime    int    `json:"best_time"` // in seconds
	CompletedAt string `json:"completed_at,omitempty"`
}

// ChallengeResponse is what gets sent to the client
type ChallengeResponse struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Story         string     `json:"story"`
	Difficulty    Difficulty `json:"difficulty"`
	Category      Category   `json:"category"`
	XPReward      int        `json:"xp_reward"`
	RequiredLevel int        `json:"required_level"`
	Hints         []string   `json:"hints"`
	StarterCode   string     `json:"starter_code"`
	Order         int        `json:"order"`
	Completed     bool       `json:"completed"`
	Locked        bool       `json:"locked"`
}
