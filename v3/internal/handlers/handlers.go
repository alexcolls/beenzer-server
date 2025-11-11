package handlers

import (
	"github.com/beenzer/beenzer-server/v3/internal/database"
	"github.com/gofiber/fiber/v2"
)

// RegisterUserRoutes registers user-related routes
func RegisterUserRoutes(router fiber.Router, db *database.DB) {
	// TODO: Implement all user endpoints from v2
	// GET /api/users/:pubkey - Get user by public key
	// POST /api/users - Create new user
	// PUT /api/users/:pubkey - Update user
	// GET /api/users/search/:query - Search users
	// POST /api/users/:pubkey/friends/:friendPubkey - Add friend
	// DELETE /api/users/:pubkey/friends/:friendPubkey - Remove friend
	// GET /api/users/:pubkey/friends - Get user friends
}

// RegisterMessageRoutes registers message-related routes
func RegisterMessageRoutes(router fiber.Router, db *database.DB) {
	// TODO: Implement all message endpoints from v2
	// GET /api/messages/:pubkey1/:pubkey2 - Get messages between two users
	// POST /api/messages - Send new message
	// PUT /api/messages/like - Like a message
	// PUT /api/messages/emoji - Add emoji to message
}

// RegisterNFTRoutes registers NFT-related routes
func RegisterNFTRoutes(router fiber.Router, db *database.DB) {
	// TODO: Implement all NFT endpoints from v2
	// POST /api/nfts/mint - Mint new NFT
	// POST /api/nfts/print - Print NFT edition
	// GET /api/nfts/user/:pubkey - Get user's NFTs
	// GET /api/nfts/map - Get NFTs by location
	// GET /api/nfts/feed - Get user feed
}

// RegisterTokenRoutes registers token-related routes
func RegisterTokenRoutes(router fiber.Router, db *database.DB) {
	// TODO: Implement all token endpoints from v2
	// GET /api/tokens/holders - Get token holders
	// GET /api/tokens/transactions - Get token transactions
	// POST /api/tokens/transactions - Add token transaction
}

// RegisterServiceRoutes registers service-related routes
func RegisterServiceRoutes(router fiber.Router, db *database.DB) {
	// TODO: Implement service endpoints from v2
	// POST /api/services/video-to-gif - Convert video to GIF
}
