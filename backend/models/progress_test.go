package models

import (
	"testing"
)

func TestNewProgress(t *testing.T) {
	userID := "test-user-123"
	progress := NewProgress(userID)

	if progress.UserID != userID {
		t.Errorf("Expected UserID=%s, got %s", userID, progress.UserID)
	}

	if progress.CompletedChallenges == nil {
		t.Error("CompletedChallenges should be initialized")
	}

	if progress.UnlockedAchievements == nil {
		t.Error("UnlockedAchievements should be initialized")
	}

	if progress.Statistics.ChallengesByDifficulty == nil {
		t.Error("ChallengesByDifficulty should be initialized")
	}

	if progress.Statistics.ChallengesByCategory == nil {
		t.Error("ChallengesByCategory should be initialized")
	}
}

func TestProgressCompleteChallenge(t *testing.T) {
	progress := NewProgress("test-user")

	// Complete first challenge
	progress.CompleteChallenge("001", DifficultyNewbie, CategoryBasics)

	if len(progress.CompletedChallenges) != 1 {
		t.Errorf("Expected 1 completed challenge, got %d", len(progress.CompletedChallenges))
	}

	if progress.Statistics.TotalChallengesCompleted != 1 {
		t.Errorf("Expected TotalChallengesCompleted=1, got %d",
			progress.Statistics.TotalChallengesCompleted)
	}

	if progress.Statistics.ChallengesByDifficulty["newbie"] != 1 {
		t.Error("ChallengesByDifficulty not updated correctly")
	}

	if progress.Statistics.ChallengesByCategory["basics"] != 1 {
		t.Error("ChallengesByCategory not updated correctly")
	}

	// Try to complete same challenge again
	progress.CompleteChallenge("001", DifficultyNewbie, CategoryBasics)

	if len(progress.CompletedChallenges) != 1 {
		t.Error("Should not add duplicate completed challenges")
	}
}

func TestProgressUnlockAchievement(t *testing.T) {
	progress := NewProgress("test-user")

	// Unlock first achievement
	unlocked := progress.UnlockAchievement("first_steps")
	if !unlocked {
		t.Error("Should return true when unlocking new achievement")
	}

	if len(progress.UnlockedAchievements) != 1 {
		t.Errorf("Expected 1 unlocked achievement, got %d", len(progress.UnlockedAchievements))
	}

	// Try to unlock same achievement again
	unlocked = progress.UnlockAchievement("first_steps")
	if unlocked {
		t.Error("Should return false when achievement already unlocked")
	}

	if len(progress.UnlockedAchievements) != 1 {
		t.Error("Should not add duplicate achievements")
	}
}
