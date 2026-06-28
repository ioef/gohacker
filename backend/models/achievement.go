package models

// Achievement represents an unlockable achievement
type Achievement struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	XPBonus     int    `json:"xp_bonus"`
	Rarity      string `json:"rarity"` // common, rare, epic, legendary
	Hidden      bool   `json:"hidden"` // Hidden until unlocked
}

// GetAllAchievements returns all available achievements
func GetAllAchievements() []Achievement {
	return []Achievement{
		{
			ID:          "first_steps",
			Title:       "First Steps",
			Description: "Complete your first challenge",
			Icon:        "🎯",
			XPBonus:     50,
			Rarity:      "common",
			Hidden:      false,
		},
		{
			ID:          "hello_world",
			Title:       "Hello, World!",
			Description: "Write your first Go program",
			Icon:        "👋",
			XPBonus:     25,
			Rarity:      "common",
			Hidden:      false,
		},
		{
			ID:          "quick_learner",
			Title:       "Quick Learner",
			Description: "Complete 5 challenges in one day",
			Icon:        "⚡",
			XPBonus:     100,
			Rarity:      "rare",
			Hidden:      false,
		},
		{
			ID:          "perfectionist",
			Title:       "Perfectionist",
			Description: "Complete a challenge without using hints",
			Icon:        "💎",
			XPBonus:     75,
			Rarity:      "rare",
			Hidden:      false,
		},
		{
			ID:          "streak_master",
			Title:       "Streak Master",
			Description: "Maintain a 7-day streak",
			Icon:        "🔥",
			XPBonus:     200,
			Rarity:      "epic",
			Hidden:      false,
		},
		{
			ID:          "go_guru",
			Title:       "Go Guru",
			Description: "Complete all challenges",
			Icon:        "🏆",
			XPBonus:     1000,
			Rarity:      "legendary",
			Hidden:      false,
		},
		{
			ID:          "concurrent_master",
			Title:       "Concurrent Master",
			Description: "Complete all concurrency challenges",
			Icon:        "🔀",
			XPBonus:     300,
			Rarity:      "epic",
			Hidden:      false,
		},
		{
			ID:          "speed_demon",
			Title:       "Speed Demon",
			Description: "Complete a challenge in under 60 seconds",
			Icon:        "🚀",
			XPBonus:     150,
			Rarity:      "rare",
			Hidden:      false,
		},
		{
			ID:          "night_owl",
			Title:       "Night Owl",
			Description: "Complete a challenge between midnight and 4 AM",
			Icon:        "🦉",
			XPBonus:     50,
			Rarity:      "common",
			Hidden:      true,
		},
		{
			ID:          "persistent",
			Title:       "Persistent",
			Description: "Attempt a challenge 10 times before succeeding",
			Icon:        "💪",
			XPBonus:     100,
			Rarity:      "rare",
			Hidden:      true,
		},
		{
			ID:          "level_10",
			Title:       "Rising Star",
			Description: "Reach level 10",
			Icon:        "⭐",
			XPBonus:     100,
			Rarity:      "common",
			Hidden:      false,
		},
		{
			ID:          "level_25",
			Title:       "Elite Hacker",
			Description: "Reach level 25",
			Icon:        "💻",
			XPBonus:     250,
			Rarity:      "rare",
			Hidden:      false,
		},
		{
			ID:          "level_50",
			Title:       "Legendary Coder",
			Description: "Reach level 50",
			Icon:        "👑",
			XPBonus:     500,
			Rarity:      "legendary",
			Hidden:      false,
		},
	}
}
