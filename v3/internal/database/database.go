package database

import (
	"context"
	"fmt"
	"time"

	"github.com/beenzer/beenzer-server/v3/pkg/config"
	"github.com/beenzer/beenzer-server/v3/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DB holds all database connections
type DB struct {
	Users    *pgxpool.Pool
	Messages *pgxpool.Pool
	NFTs     *pgxpool.Pool
	Tokens   *pgxpool.Pool
}

var Database *DB

// Connect establishes connections to all databases
func Connect(cfg *config.Config) (*DB, error) {
	ctx := context.Background()

	// Create pool configs
	usersPool, err := createPool(ctx, cfg.UsersDBURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to users database: %w", err)
	}

	messagesPool, err := createPool(ctx, cfg.MessagesDBURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to messages database: %w", err)
	}

	nftsPool, err := createPool(ctx, cfg.NFTsDBURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to nfts database: %w", err)
	}

	tokensPool, err := createPool(ctx, cfg.TokenDBURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to tokens database: %w", err)
	}

	db := &DB{
		Users:    usersPool,
		Messages: messagesPool,
		NFTs:     nftsPool,
		Tokens:   tokensPool,
	}

	Database = db
	logger.Log.Info().Msg("✅ Successfully connected to all databases")

	// Initialize database tables
	if err := db.InitializeTables(ctx); err != nil {
		return nil, fmt.Errorf("failed to initialize tables: %w", err)
	}

	return db, nil
}

// createPool creates a new connection pool
func createPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	// Configure pool settings
	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = time.Minute * 30
	config.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}

// Close closes all database connections
func (db *DB) Close() {
	if db.Users != nil {
		db.Users.Close()
	}
	if db.Messages != nil {
		db.Messages.Close()
	}
	if db.NFTs != nil {
		db.NFTs.Close()
	}
	if db.Tokens != nil {
		db.Tokens.Close()
	}
	logger.Log.Info().Msg("🔒 Database connections closed")
}

// InitializeTables creates all required tables if they don't exist
func (db *DB) InitializeTables(ctx context.Context) error {
	// Users database tables
	if err := db.createUsersTable(ctx); err != nil {
		return err
	}

	// NFTs database tables
	if err := db.createNFTsTables(ctx); err != nil {
		return err
	}

	// Tokens database tables
	if err := db.createTokensTables(ctx); err != nil {
		return err
	}

	logger.Log.Info().Msg("✅ Database tables initialized")
	return nil
}

func (db *DB) createUsersTable(ctx context.Context) error {
	// Users table
	_, err := db.Users.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			__pubkey__ VARCHAR(255) PRIMARY KEY,
			_username_ VARCHAR(55) UNIQUE,
			_pfp VARCHAR(555),
			_name VARCHAR(55),
			_lastname VARCHAR(55),
			_description VARCHAR(255),
			_birthdate VARCHAR(55),
			_country VARCHAR(155),
			_flag VARCHAR(10),
			_city VARCHAR(155),
			_phone VARCHAR(55),
			_email VARCHAR(55),
			_verified BOOLEAN DEFAULT FALSE,
			_twitter VARCHAR(55),
			_instagram VARCHAR(55),
			_discord VARCHAR(55),
			_telegram VARCHAR(55),
			_youtube VARCHAR(55),
			_tiktok VARCHAR(55),
			_magiceden VARCHAR(55),
			_opensea VARCHAR(55),
			_appuser BOOLEAN DEFAULT FALSE,
			_created_at BIGINT,
			_timestamp BIGINT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Logs table
	_, err = db.Users.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS logs (
			_pubkey VARCHAR(255),
			_logs VARCHAR(255),
			_timestamp BIGINT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create logs table: %w", err)
	}

	// Friends table
	_, err = db.Users.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS friends (
			__pubkey__ VARCHAR(255),
			__pubkey2__ VARCHAR(255),
			_timestamp BIGINT,
			PRIMARY KEY (__pubkey__, __pubkey2__)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create friends table: %w", err)
	}

	return nil
}

func (db *DB) createNFTsTables(ctx context.Context) error {
	// NFTs table
	_, err := db.NFTs.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS nfts (
			_id_ BIGINT NOT NULL UNIQUE,
			__token__ VARCHAR(255) PRIMARY KEY,
			_supply INT,
			_floor FLOAT,
			_ccy VARCHAR(10),
			_creator VARCHAR(255),
			_username VARCHAR(255),
			_image VARCHAR(255),
			_asset VARCHAR(255),
			_type VARCHAR(25),
			_metadata VARCHAR(255),
			_name VARCHAR(255),
			_description VARCHAR(2048),
			_city VARCHAR(75),
			_latitude DOUBLE PRECISION,
			_longitude DOUBLE PRECISION,
			_visibility VARCHAR(55),
			_maxLat DOUBLE PRECISION,
			_minLat DOUBLE PRECISION,
			_maxLon DOUBLE PRECISION,
			_minLon DOUBLE PRECISION,
			_date VARCHAR(22),
			_time VARCHAR(22),
			_timestamp BIGINT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create nfts table: %w", err)
	}

	// Editions table
	_, err = db.NFTs.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS editions (
			__master__ VARCHAR(255),
			__edition__ VARCHAR(255),
			_minter VARCHAR(255),
			_id INT,
			_date VARCHAR(55),
			_time VARCHAR(55),
			_timestamp BIGINT,
			PRIMARY KEY (__master__, __edition__)
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create editions table: %w", err)
	}

	// Counter table
	_, err = db.NFTs.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS counter (
			_n BIGINT,
			_timestamp BIGINT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create counter table: %w", err)
	}

	// Initialize counter if empty
	_, err = db.NFTs.Exec(ctx, `
		INSERT INTO counter (_n, _timestamp)
		SELECT 0, 0
		WHERE NOT EXISTS (SELECT 1 FROM counter)
	`)
	if err != nil {
		return fmt.Errorf("failed to initialize counter: %w", err)
	}

	// Transactions table
	_, err = db.NFTs.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS transactions (
			_owner VARCHAR(255),
			_pubkey VARCHAR(255),
			_type VARCHAR(5),
			_currency VARCHAR(5),
			_amount FLOAT,
			_hash VARCHAR(255),
			_timestamp BIGINT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create transactions table: %w", err)
	}

	return nil
}

func (db *DB) createTokensTables(ctx context.Context) error {
	// Token transactions table
	_, err := db.Tokens.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS transactions (
			_date VARCHAR(20),
			_time VARCHAR(20),
			_type VARCHAR(10),
			_amount FLOAT,
			_pubkey VARCHAR(255),
			_flag VARCHAR(10),
			_timestamp BIGINT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create token transactions table: %w", err)
	}

	// Token holders table
	_, err = db.Tokens.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS holders (
			__position__ INT PRIMARY KEY,
			_percentage FLOAT,
			_amount FLOAT,
			_pubkey VARCHAR(255),
			_flag VARCHAR(10),
			_timestamp BIGINT
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create holders table: %w", err)
	}

	return nil
}
