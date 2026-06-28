package main

import (
	"fmt"
	"gohacker/api/routes"
	"gohacker/config"
	"gohacker/game"
	"gohacker/storage"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.NewConfig()

	// Create data directory
	if err := os.MkdirAll(cfg.Storage.DataDir, 0755); err != nil {
		log.Fatalf("Failed to create data directory: %v", err)
	}

	// Initialize storage
	store, err := storage.NewStore()
	if err != nil {
		log.Fatalf("Failed to create store: %v", err)
	}

	// Initialize persistence manager
	pm := storage.NewPersistenceManager(store, cfg.Storage.PersistenceFile)

	// Load existing data
	if err := pm.Load(); err != nil {
		log.Printf("Warning: Failed to load data: %v", err)
	} else {
		log.Println("✅ Data loaded successfully")
	}

	// Start auto-save
	pm.StartAutoSave(cfg.Storage.SaveInterval)
	defer pm.Stop()

	// Initialize game engine
	engine := game.NewEngine(store)
	log.Printf("✅ Game engine initialized with %d challenges", len(engine.GetAllChallenges()))

	// Setup Gin
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, store, engine, cfg)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		log.Println("\n🛑 Shutting down gracefully...")
		pm.Stop()
		os.Exit(0)
	}()

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	log.Printf("🚀 GoHacker API starting on %s", addr)
	log.Printf("🎮 Ready to hack! Visit http://localhost:%s/api/health", cfg.Server.Port)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
