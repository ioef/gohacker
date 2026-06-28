package config

import (
	"time"
)

// Config holds all application configuration
type Config struct {
	Server  ServerConfig
	JWT     JWTConfig
	Game    GameConfig
	Storage StorageConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port            string
	Host            string
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
}

// JWTConfig holds JWT authentication configuration
type JWTConfig struct {
	Secret     string
	Expiration time.Duration
}

// GameConfig holds game-related configuration
type GameConfig struct {
	MaxCodeLength      int
	ExecutionTimeout   time.Duration
	MaxConcurrentExecs int
	AutoSaveInterval   time.Duration
}

// StorageConfig holds storage configuration
type StorageConfig struct {
	DataDir         string
	PersistenceFile string
	SaveInterval    time.Duration
}

// NewConfig returns a new Config with default values
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:            "8080",
			Host:            "0.0.0.0",
			ReadTimeout:     10 * time.Second,
			WriteTimeout:    10 * time.Second,
			ShutdownTimeout: 5 * time.Second,
		},
		JWT: JWTConfig{
			Secret:     "your-super-secret-jwt-key-change-this-in-production",
			Expiration: 24 * time.Hour,
		},
		Game: GameConfig{
			MaxCodeLength:      10000,
			ExecutionTimeout:   5 * time.Second,
			MaxConcurrentExecs: 10,
			AutoSaveInterval:   30 * time.Second,
		},
		Storage: StorageConfig{
			DataDir:         "./data",
			PersistenceFile: "./data/gohacker.json",
			SaveInterval:    60 * time.Second,
		},
	}
}
