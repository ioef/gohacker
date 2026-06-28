package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User represents a player in the game
type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Never send password hash to client
	Level        int       `json:"level"`
	XP           int       `json:"xp"`
	TotalXP      int       `json:"total_xp"`
	Rank         string    `json:"rank"`
	CreatedAt    time.Time `json:"created_at"`
	LastLogin    time.Time `json:"last_login"`
}

// SetPassword hashes and sets the user's password
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

// CheckPassword verifies if the provided password matches the user's password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

// CalculateLevel calculates the user's level based on total XP
func (u *User) CalculateLevel() {
	// Level formula: Level = floor(sqrt(TotalXP / 100))
	// This creates a nice progression curve
	xpPerLevel := 100
	level := 1
	requiredXP := xpPerLevel

	for u.TotalXP >= requiredXP {
		level++
		requiredXP += xpPerLevel * level
	}

	u.Level = level
	u.XP = u.TotalXP - (requiredXP - (xpPerLevel * level))
}

// GetRank returns the user's rank based on level
func (u *User) GetRank() string {
	switch {
	case u.Level >= 36:
		return "Master Hacker"
	case u.Level >= 21:
		return "Elite Coder"
	case u.Level >= 11:
		return "Script Kiddie"
	default:
		return "Newbie Hacker"
	}
}

// AddXP adds experience points and recalculates level
func (u *User) AddXP(xp int) {
	u.TotalXP += xp
	u.CalculateLevel()
	u.Rank = u.GetRank()
}

// UserResponse is the sanitized user data sent to clients
type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Level     int       `json:"level"`
	XP        int       `json:"xp"`
	TotalXP   int       `json:"total_xp"`
	Rank      string    `json:"rank"`
	CreatedAt time.Time `json:"created_at"`
	LastLogin time.Time `json:"last_login"`
}

// ToResponse converts User to UserResponse (removes sensitive data)
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Level:     u.Level,
		XP:        u.XP,
		TotalXP:   u.TotalXP,
		Rank:      u.Rank,
		CreatedAt: u.CreatedAt,
		LastLogin: u.LastLogin,
	}
}
