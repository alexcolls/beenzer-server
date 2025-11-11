package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/beenzer/beenzer-server/v3/internal/database"
	"github.com/beenzer/beenzer-server/v3/internal/handlers"
	"github.com/beenzer/beenzer-server/v3/internal/middleware"
	"github.com/beenzer/beenzer-server/v3/internal/websocket"
	"github.com/beenzer/beenzer-server/v3/pkg/config"
	"github.com/beenzer/beenzer-server/v3/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Beenzer Server API v3
// @version 3.0.0
// @description Geo-social backend platform with Solana blockchain integration - Go implementation
// @termsOfService https://beenzer.com/terms

// @contact.name Beenzer Support
// @contact.url https://beenzer.com/support
// @contact.email support@beenzer.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:4000
// @BasePath /
// @schemes http https

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("❌ Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	logger.Init(cfg.GoEnv)
	logger.Log.Info().Msg("🚀 Starting Beenzer Server v3")

	// Connect to databases
	db, err := database.Connect(cfg)
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("Failed to connect to databases")
	}
	defer db.Close()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:               "Beenzer Server v3",
		ServerHeader:          "Beenzer",
		BodyLimit:             1 * 1024 * 1024 * 1024, // 1GB (matching v2)
		ReadTimeout:           60 * time.Second,
		WriteTimeout:          60 * time.Second,
		IdleTimeout:           120 * time.Second,
		DisableStartupMessage: false,
		EnablePrintRoutes:     cfg.GoEnv == "development",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // In production, restrict to specific origins
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
		MaxAge:           86400,
	}))

	// Custom middleware
	app.Use(middleware.RequestID())

	// Routes
	setupRoutes(app, db)

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// WebSocket hub
	hub := websocket.NewHub()
	go hub.Run()

	// WebSocket endpoint
	app.Get("/ws", websocket.UpgradeConnection(hub, db))

	// Health check
	app.Get("/", healthCheck)
	app.Get("/health", healthCheck)

	// Start background jobs
	go startBackgroundJobs(cfg, db)

	// Start server
	port := cfg.Port
	logger.Log.Info().Msgf("🌐 Server running on port %s", port)
	logger.Log.Info().Msgf("📚 Swagger docs available at http://localhost:%s/swagger/", port)
	logger.Log.Info().Msgf("🔌 WebSocket endpoint at ws://localhost:%s/ws", port)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":" + port); err != nil {
			logger.Log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	<-quit
	logger.Log.Info().Msg("🛑 Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		logger.Log.Error().Err(err).Msg("Server forced to shutdown")
	}

	logger.Log.Info().Msg("✅ Server exited gracefully")
}

func setupRoutes(app *fiber.App, db *database.DB) {
	api := app.Group("/api")

	// User routes
	users := api.Group("/users")
	handlers.RegisterUserRoutes(users, db)

	// Message routes
	messages := api.Group("/messages")
	handlers.RegisterMessageRoutes(messages, db)

	// NFT routes
	nfts := api.Group("/nfts")
	handlers.RegisterNFTRoutes(nfts, db)

	// Token routes
	tokens := api.Group("/tokens")
	handlers.RegisterTokenRoutes(tokens, db)

	// Service routes
	services := api.Group("/services")
	handlers.RegisterServiceRoutes(services, db)
}

// healthCheck godoc
// @Summary Health check endpoint
// @Description Check if the server is running
// @Tags health
// @Accept json
// @Produce html
// @Success 200 {string} string "Server is running"
// @Router / [get]
func healthCheck(c *fiber.Ctx) error {
	return c.SendString("<h1>Beenzer Server v3</h1><p>Go + Fiber Implementation</p>")
}

func startBackgroundJobs(cfg *config.Config, db *database.DB) {
	// Balance check every 15 minutes (matching v2)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	logger.Log.Info().Msg("⏰ Background jobs started")

	for range ticker.C {
		now := time.Now()
		
		// Log time every minute at :00 seconds
		if now.Second() == 0 {
			logger.Log.Info().Msgf("🕐 %s", now.Format("15:04:05"))
			
			// Check balances every 15 minutes
			if now.Minute()%15 == 0 {
				logger.Log.Info().Msg("💰 Checking balances...")
				// TODO: Implement balance check
				// This would call the Solana blockchain service to check balances
			}
		}
	}
}
