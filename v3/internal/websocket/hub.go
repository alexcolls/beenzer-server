package websocket

import (
	"encoding/json"
	"sync"

	"github.com/beenzer/beenzer-server/v3/pkg/logger"
)

// Client represents a WebSocket client
type Client struct {
	ID   string
	Conn *Conn
	Hub  *Hub
	Send chan []byte
}

// Hub manages all active WebSocket connections
type Hub struct {
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	mu         sync.RWMutex
}

// Message represents a WebSocket message
type Message struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

// NewHub creates a new WebSocket hub
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan []byte, 256),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Run starts the hub's main loop
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.mu.Lock()
			h.Clients[client] = true
			h.mu.Unlock()
			logger.Log.Info().Msgf("🔌 Client connected. Total: %d", len(h.Clients))
			
			// Send connection confirmation
			h.SendToClient(client, "serverConnection", "Client connected to server successfully")
			h.SendToClient(client, "nUsers", len(h.Clients))

		case client := <-h.Unregister:
			h.mu.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
			h.mu.Unlock()
			logger.Log.Info().Msgf("🔌 Client disconnected. Total: %d", len(h.Clients))

		case message := <-h.Broadcast:
			h.mu.RLock()
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
			h.mu.RUnlock()
		}
	}
}

// SendToClient sends a message to a specific client
func (h *Hub) SendToClient(client *Client, event string, data interface{}) {
	msg := Message{
		Event: event,
		Data:  data,
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to marshal message")
		return
	}
	
	select {
	case client.Send <- jsonData:
	default:
		logger.Log.Warn().Msg("Client send channel is full")
	}
}

// BroadcastToAll broadcasts a message to all connected clients
func (h *Hub) BroadcastToAll(event string, data interface{}) {
	msg := Message{
		Event: event,
		Data:  data,
	}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		logger.Log.Error().Err(err).Msg("Failed to marshal broadcast message")
		return
	}
	
	h.Broadcast <- jsonData
}

// GetConnectedCount returns the number of connected clients
func (h *Hub) GetConnectedCount() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.Clients)
}
