package storage

import (
	"encoding/json"
	"fmt"
	"gohacker/models"
	"os"
	"sync"
	"time"
)

// PersistenceData represents all data to be persisted
type PersistenceData struct {
	Users             []*models.User              `json:"users"`
	Progress          []*models.Progress          `json:"progress"`
	ChallengeProgress []*models.ChallengeProgress `json:"challenge_progress"`
	LastSaved         time.Time                   `json:"last_saved"`
}

// PersistenceManager handles saving and loading data
type PersistenceManager struct {
	store    *Store
	filePath string
	mu       sync.RWMutex
	stopCh   chan struct{}
}

// NewPersistenceManager creates a new persistence manager
func NewPersistenceManager(store *Store, filePath string) *PersistenceManager {
	return &PersistenceManager{
		store:    store,
		filePath: filePath,
		stopCh:   make(chan struct{}),
	}
}

// Save saves all data to disk
func (pm *PersistenceManager) Save() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Collect all data from memdb
	users, err := pm.store.GetAllUsers()
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	// Get all progress (we'll need to iterate through users)
	var allProgress []*models.Progress
	var allChallengeProgress []*models.ChallengeProgress
	for _, user := range users {
		if progress, err := pm.store.GetProgress(user.ID); err == nil {
			allProgress = append(allProgress, progress)
		}
		if cp, err := pm.store.GetUserChallengeProgress(user.ID); err == nil {
			allChallengeProgress = append(allChallengeProgress, cp...)
		}
	}

	data := PersistenceData{
		Users:             users,
		Progress:          allProgress,
		ChallengeProgress: allChallengeProgress,
		LastSaved:         time.Now(),
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	// Write to file
	if err := os.WriteFile(pm.filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Load loads all data from disk
func (pm *PersistenceManager) Load() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	// Check if file exists
	if _, err := os.Stat(pm.filePath); os.IsNotExist(err) {
		return nil // No file to load, that's okay
	}

	// Read file
	jsonData, err := os.ReadFile(pm.filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Unmarshal JSON
	var data PersistenceData
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return fmt.Errorf("failed to unmarshal data: %w", err)
	}

	// Load data into memdb
	for _, user := range data.Users {
		if err := pm.store.CreateUser(user); err != nil {
			return fmt.Errorf("failed to load user: %w", err)
		}
	}

	for _, progress := range data.Progress {
		if err := pm.store.CreateProgress(progress); err != nil {
			return fmt.Errorf("failed to load progress: %w", err)
		}
	}

	for _, cp := range data.ChallengeProgress {
		if err := pm.store.CreateChallengeProgress(cp); err != nil {
			return fmt.Errorf("failed to load challenge progress: %w", err)
		}
	}

	return nil
}

// StartAutoSave starts automatic periodic saving
func (pm *PersistenceManager) StartAutoSave(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := pm.Save(); err != nil {
					fmt.Printf("Auto-save error: %v\n", err)
				}
			case <-pm.stopCh:
				ticker.Stop()
				return
			}
		}
	}()
}

// Stop stops the auto-save goroutine
func (pm *PersistenceManager) Stop() {
	close(pm.stopCh)
	// Do a final save
	if err := pm.Save(); err != nil {
		fmt.Printf("Final save error: %v\n", err)
	}
}
