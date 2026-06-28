package storage

import (
	"fmt"
	"gohacker/models"

	"github.com/hashicorp/go-memdb"
)

// Store represents the in-memory database
type Store struct {
	db *memdb.MemDB
}

// NewStore creates a new in-memory store
func NewStore() (*Store, error) {
	schema := CreateSchema()
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		return nil, fmt.Errorf("failed to create memdb: %w", err)
	}

	return &Store{db: db}, nil
}

// User operations

// CreateUser creates a new user
func (s *Store) CreateUser(user *models.User) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("users", user); err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	txn.Commit()
	return nil
}

// GetUserByID retrieves a user by ID
func (s *Store) GetUserByID(id string) (*models.User, error) {
	txn := s.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("users", "id", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if raw == nil {
		return nil, fmt.Errorf("user not found")
	}

	return raw.(*models.User), nil
}

// GetUserByUsername retrieves a user by username
func (s *Store) GetUserByUsername(username string) (*models.User, error) {
	txn := s.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("users", "username", username)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if raw == nil {
		return nil, fmt.Errorf("user not found")
	}

	return raw.(*models.User), nil
}

// GetUserByEmail retrieves a user by email
func (s *Store) GetUserByEmail(email string) (*models.User, error) {
	txn := s.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("users", "email", email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	if raw == nil {
		return nil, fmt.Errorf("user not found")
	}

	return raw.(*models.User), nil
}

// UpdateUser updates an existing user
func (s *Store) UpdateUser(user *models.User) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("users", user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	txn.Commit()
	return nil
}

// GetAllUsers retrieves all users
func (s *Store) GetAllUsers() ([]*models.User, error) {
	txn := s.db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("users", "id")
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	var users []*models.User
	for obj := it.Next(); obj != nil; obj = it.Next() {
		users = append(users, obj.(*models.User))
	}

	return users, nil
}

// Progress operations

// CreateProgress creates new progress for a user
func (s *Store) CreateProgress(progress *models.Progress) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("progress", progress); err != nil {
		return fmt.Errorf("failed to insert progress: %w", err)
	}

	txn.Commit()
	return nil
}

// GetProgress retrieves progress for a user
func (s *Store) GetProgress(userID string) (*models.Progress, error) {
	txn := s.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("progress", "id", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get progress: %w", err)
	}

	if raw == nil {
		return nil, fmt.Errorf("progress not found")
	}

	return raw.(*models.Progress), nil
}

// UpdateProgress updates user progress
func (s *Store) UpdateProgress(progress *models.Progress) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("progress", progress); err != nil {
		return fmt.Errorf("failed to update progress: %w", err)
	}

	txn.Commit()
	return nil
}

// Challenge Progress operations

// CreateChallengeProgress creates new challenge progress
func (s *Store) CreateChallengeProgress(cp *models.ChallengeProgress) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("challenge_progress", cp); err != nil {
		return fmt.Errorf("failed to insert challenge progress: %w", err)
	}

	txn.Commit()
	return nil
}

// GetChallengeProgress retrieves challenge progress for a user and challenge
func (s *Store) GetChallengeProgress(userID, challengeID string) (*models.ChallengeProgress, error) {
	txn := s.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("challenge_progress", "id", userID, challengeID)
	if err != nil {
		return nil, fmt.Errorf("failed to get challenge progress: %w", err)
	}

	if raw == nil {
		return nil, fmt.Errorf("challenge progress not found")
	}

	return raw.(*models.ChallengeProgress), nil
}

// GetUserChallengeProgress retrieves all challenge progress for a user
func (s *Store) GetUserChallengeProgress(userID string) ([]*models.ChallengeProgress, error) {
	txn := s.db.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("challenge_progress", "user_id", userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get challenge progress: %w", err)
	}

	var progress []*models.ChallengeProgress
	for obj := it.Next(); obj != nil; obj = it.Next() {
		progress = append(progress, obj.(*models.ChallengeProgress))
	}

	return progress, nil
}

// UpdateChallengeProgress updates challenge progress
func (s *Store) UpdateChallengeProgress(cp *models.ChallengeProgress) error {
	txn := s.db.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("challenge_progress", cp); err != nil {
		return fmt.Errorf("failed to update challenge progress: %w", err)
	}

	txn.Commit()
	return nil
}
