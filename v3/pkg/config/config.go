package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	// Server
	Port  string
	GoEnv string

	// Database URLs
	UsersDBURL    string
	MessagesDBURL string
	NFTsDBURL     string
	TokenDBURL    string

	// Solana
	SolanaRPCURL         string
	MasterWallet         string
	MasterWalletKeypair  []byte
	NFTSymbol            string
	NFTMaster            string
	NFTMasterSupply      int
	CurrencySymbol       string
	CurrencyToken        string
	StockSymbol          string
	StockToken           string
	TokenAuthority       string
	TokenAuthorityKeypair []byte
	TokenAccount         string
	TokenCreator         string

	// Metaplex
	MetaplexBundlrURI string

	// Marketplace & DAO
	MarketWallet string
	MarketURL    string
	DAOWallet    string
	DAOURL       string

	// Client URLs
	ClientAppURL    string
	ClientDAOURL    string
	ClientMarketURL string
	LandingURL      string

	// Social Media
	Twitter   string
	Instagram string
	Discord   string
	Telegram  string
	TikTok    string
	YouTube   string
	OpenSea   string
	MagicEden string

	// Google Maps
	GoogleMapsAPIKey string

	// Application Settings
	AsyncTries  int
	Description string
}

var AppConfig *Config

// Load reads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		// Server
		Port:  getEnv("PORT", "4000"),
		GoEnv: getEnv("GO_ENV", "development"),

		// Database
		UsersDBURL:    getEnv("USERS_DB_URL", ""),
		MessagesDBURL: getEnv("MESSAGES_DB_URL", ""),
		NFTsDBURL:     getEnv("NFTS_DB_URL", ""),
		TokenDBURL:    getEnv("TOKEN_DB_URL", ""),

		// Solana
		SolanaRPCURL:    getEnv("SOLANA_RPC_URL", "https://api.devnet.solana.com"),
		MasterWallet:    getEnv("MASTER_WALLET", ""),
		NFTSymbol:       getEnv("NFT_SYMBOL", "BEENZER"),
		NFTMaster:       getEnv("NFT_MASTER", ""),
		NFTMasterSupply: getEnvAsInt("NFT_MASTER_SUPPLY", 1000),
		CurrencySymbol:  getEnv("CURRENCY_SYMBOL", "BEEN"),
		CurrencyToken:   getEnv("CURRENCY_TOKEN", ""),
		StockSymbol:     getEnv("STOCK_SYMBOL", "BEENS"),
		StockToken:      getEnv("STOCK_TOKEN", ""),
		TokenAuthority:  getEnv("TOKEN_AUTHORITY", ""),
		TokenAccount:    getEnv("TOKEN_ACCOUNT", ""),
		TokenCreator:    getEnv("TOKEN_CREATOR", ""),

		// Metaplex
		MetaplexBundlrURI: getEnv("METAPLEX_BUNDLR_URI", "https://node1.bundlr.network"),

		// Marketplace & DAO
		MarketWallet: getEnv("MARKET_WALLET", ""),
		MarketURL:    getEnv("MARKET_URL", "https://market.beenzer.app"),
		DAOWallet:    getEnv("DAO_WALLET", ""),
		DAOURL:       getEnv("DAO_URL", "https://dao.beenzer.app"),

		// Client URLs
		ClientAppURL:    getEnv("CLIENT_APP_URL", ""),
		ClientDAOURL:    getEnv("CLIENT_DAO_URL", ""),
		ClientMarketURL: getEnv("CLIENT_MARKET_URL", ""),
		LandingURL:      getEnv("LANDING_URL", ""),

		// Social Media
		Twitter:   getEnv("TWITTER", "https://twitter.com/beenzer_app"),
		Instagram: getEnv("INSTAGRAM", "https://instagram.com/beenzer_app"),
		Discord:   getEnv("DISCORD", "https://discord.gg/Ta9X6zbg"),
		Telegram:  getEnv("TELEGRAM", "https://t.me/+VgZorKQGP0gwY2Fk"),
		TikTok:    getEnv("TIKTOK", "https://tiktok.com/beenzer_app"),
		YouTube:   getEnv("YOUTUBE", "https://youtube.com/@beenzer"),
		OpenSea:   getEnv("OPENSEA", "https://opensea.com/beenzer_dapp"),
		MagicEden: getEnv("MAGICEDEN", "https://magiceden.io/beenzer_dapp"),

		// Google Maps
		GoogleMapsAPIKey: getEnv("GOOGLE_MAPS_API_KEY", ""),

		// Application Settings
		AsyncTries:  getEnvAsInt("ASYNC_TRIES", 3),
		Description: getEnv("DESCRIPTION", "Beenzer Collection is the NFT Master Edition for BEENZER official NFTs. Check our links to be part of the best web3 community! 💚"),
	}

	// Parse keypairs
	var err error
	config.MasterWalletKeypair, err = parseKeypair(getEnv("MASTER_WALLET_KEYPAIR", ""))
	if err != nil {
		return nil, fmt.Errorf("failed to parse master wallet keypair: %w", err)
	}

	config.TokenAuthorityKeypair, err = parseKeypair(getEnv("TOKEN_AUTHORITY_KEYPAIR", ""))
	if err != nil {
		return nil, fmt.Errorf("failed to parse token authority keypair: %w", err)
	}

	// Validate required fields
	if err := config.Validate(); err != nil {
		return nil, err
	}

	AppConfig = config
	return config, nil
}

// Validate checks if all required configuration is set
func (c *Config) Validate() error {
	if c.UsersDBURL == "" {
		return fmt.Errorf("USERS_DB_URL is required")
	}
	if c.MessagesDBURL == "" {
		return fmt.Errorf("MESSAGES_DB_URL is required")
	}
	if c.NFTsDBURL == "" {
		return fmt.Errorf("NFTS_DB_URL is required")
	}
	if c.TokenDBURL == "" {
		return fmt.Errorf("TOKEN_DB_URL is required")
	}
	return nil
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func parseKeypair(keypairStr string) ([]byte, error) {
	if keypairStr == "" {
		return make([]byte, 64), nil
	}

	parts := strings.Split(keypairStr, ",")
	if len(parts) != 64 {
		return nil, fmt.Errorf("keypair must have 64 bytes, got %d", len(parts))
	}

	keypair := make([]byte, 64)
	for i, part := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			return nil, fmt.Errorf("invalid byte at position %d: %w", i, err)
		}
		if val < 0 || val > 255 {
			return nil, fmt.Errorf("byte at position %d out of range: %d", i, val)
		}
		keypair[i] = byte(val)
	}

	return keypair, nil
}
