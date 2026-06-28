package models

import (
	"testing"
)

func TestUserSetPassword(t *testing.T) {
	user := &User{}
	password := "testpassword123"

	err := user.SetPassword(password)
	if err != nil {
		t.Fatalf("SetPassword failed: %v", err)
	}

	if user.PasswordHash == "" {
		t.Error("PasswordHash should not be empty")
	}

	if user.PasswordHash == password {
		t.Error("PasswordHash should not be the plain password")
	}
}

func TestUserCheckPassword(t *testing.T) {
	user := &User{}
	password := "testpassword123"

	user.SetPassword(password)

	// Test correct password
	if !user.CheckPassword(password) {
		t.Error("CheckPassword should return true for correct password")
	}

	// Test incorrect password
	if user.CheckPassword("wrongpassword") {
		t.Error("CheckPassword should return false for incorrect password")
	}
}

func TestUserCalculateLevel(t *testing.T) {
	tests := []struct {
		totalXP       int
		expectedLevel int
	}{
		{0, 1},
		{100, 2},
		{300, 3},
		{600, 4},
		{1000, 5},
	}

	for _, tt := range tests {
		user := &User{TotalXP: tt.totalXP}
		user.CalculateLevel()

		if user.Level != tt.expectedLevel {
			t.Errorf("For TotalXP=%d, expected Level=%d, got Level=%d",
				tt.totalXP, tt.expectedLevel, user.Level)
		}
	}
}

func TestUserGetRank(t *testing.T) {
	tests := []struct {
		level        int
		expectedRank string
	}{
		{1, "Newbie Hacker"},
		{10, "Newbie Hacker"},
		{11, "Script Kiddie"},
		{20, "Script Kiddie"},
		{21, "Elite Coder"},
		{35, "Elite Coder"},
		{36, "Master Hacker"},
		{50, "Master Hacker"},
	}

	for _, tt := range tests {
		user := &User{Level: tt.level}
		rank := user.GetRank()

		if rank != tt.expectedRank {
			t.Errorf("For Level=%d, expected Rank=%s, got Rank=%s",
				tt.level, tt.expectedRank, rank)
		}
	}
}

func TestUserAddXP(t *testing.T) {
	user := &User{
		Level:   1,
		XP:      0,
		TotalXP: 0,
	}

	user.AddXP(150)

	if user.TotalXP != 150 {
		t.Errorf("Expected TotalXP=150, got %d", user.TotalXP)
	}

	if user.Level < 2 {
		t.Errorf("Expected Level to increase, got %d", user.Level)
	}
}

func TestUserToResponse(t *testing.T) {
	user := &User{
		ID:           "test-id",
		Username:     "testuser",
		Email:        "test@example.com",
		PasswordHash: "secret-hash",
		Level:        5,
		XP:           50,
		TotalXP:      500,
		Rank:         "Script Kiddie",
	}

	response := user.ToResponse()

	if response.ID != user.ID {
		t.Error("Response ID mismatch")
	}

	if response.Username != user.Username {
		t.Error("Response Username mismatch")
	}

	// Password hash should not be in response (it's not a field in UserResponse)
	// This is implicitly tested by the struct definition
}
