package websocket

import (
	"encoding/json"

	"github.com/beenzer/beenzer-server/v3/internal/database"
	"github.com/beenzer/beenzer-server/v3/pkg/logger"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Conn wraps the websocket connection
type Conn = websocket.Conn

// UpgradeConnection upgrades HTTP connection to WebSocket
func UpgradeConnection(hub *Hub, db *database.DB) fiber.Handler {
	return websocket.New(func(c *websocket.Conn) {
		clientID := uuid.New().String()
		
		client := &Client{
			ID:   clientID,
			Hub:  hub,
			Send: make(chan []byte, 256),
		}

		hub.Register <- client

		// Start goroutines for reading and writing
		go client.writePump(c)
		client.readPump(c, db)
	})
}

// readPump reads messages from the WebSocket connection
func (c *Client) readPump(conn *websocket.Conn, db *database.DB) {
	defer func() {
		c.Hub.Unregister <- c
		conn.Close()
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Log.Error().Err(err).Msg("WebSocket error")
			}
			break
		}

		if messageType == websocket.TextMessage {
			// Parse the message
			var msg Message
			if err := json.Unmarshal(message, &msg); err != nil {
				logger.Log.Error().Err(err).Msg("Failed to parse WebSocket message")
				continue
			}

			// Handle the message event
			HandleWebSocketEvent(c, db, msg.Event, msg.Data)
		}
	}
}

// writePump writes messages to the WebSocket connection
func (c *Client) writePump(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}

// HandleWebSocketEvent processes incoming WebSocket events
func HandleWebSocketEvent(client *Client, db *database.DB, event string, data interface{}) {
	logger.Log.Debug().Msgf("📨 WebSocket event: %s", event)

	// Route to appropriate handler based on event
	switch event {
	// User events
	case "newConnection":
		handleNewConnection(client, db, data)
	case "newUser":
		handleNewUser(client, db, data)
	case "getUser":
		handleGetUser(client, db, data)
	case "updateUser":
		handleUpdateUser(client, db, data)
	case "searchUsers":
		handleSearchUsers(client, db, data)
	case "addFriend":
		handleAddFriend(client, db, data)
	case "deleteFriend":
		handleDeleteFriend(client, db, data)
	case "getUserFriends":
		handleGetUserFriends(client, db, data)
	
	// Message events
	case "getMessages":
		handleGetMessages(client, db, data)
	case "newMessage":
		handleNewMessage(client, db, data)
	case "likeMessage":
		handleLikeMessage(client, db, data)
	case "addEmoji":
		handleAddEmoji(client, db, data)
	
	// NFT events
	case "mintNFT":
		handleMintNFT(client, db, data)
	case "printNFT":
		handlePrintNFT(client, db, data)
	case "getUserNFTs":
		handleGetUserNFTs(client, db, data)
	case "getMapNFTs":
		handleGetMapNFTs(client, db, data)
	case "getFeed":
		handleGetFeed(client, db, data)
	
	// Token events
	case "getTokenHolders":
		handleGetTokenHolders(client, db, data)
	case "getTokenTransactions":
		handleGetTokenTransactions(client, db, data)
	case "addTokenTransaction":
		handleAddTokenTransaction(client, db, data)
	
	default:
		logger.Log.Warn().Msgf("Unknown WebSocket event: %s", event)
	}
}

// Placeholder handlers - These would need full implementation
func handleNewConnection(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement full logic from v2/src/services/sockets/users.socket.ts
	client.Hub.SendToClient(client, "isNewUser", false)
}

func handleNewUser(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleGetUser(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleUpdateUser(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleSearchUsers(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleAddFriend(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleDeleteFriend(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleGetUserFriends(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleGetMessages(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleNewMessage(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleLikeMessage(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleAddEmoji(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleMintNFT(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement full NFT minting logic
}

func handlePrintNFT(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleGetUserNFTs(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleGetMapNFTs(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleGetFeed(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleGetTokenHolders(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleGetTokenTransactions(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}

func handleAddTokenTransaction(client *Client, db *database.DB, data interface{}) {
	// TODO: Implement
}
