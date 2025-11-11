package models

// Message represents a message in a conversation
type Message struct {
	Owner     string `json:"_owner" db:"_owner"`
	Message   string `json:"_message" db:"_message"`
	Liked     bool   `json:"_liked" db:"_liked"`
	Emoji     string `json:"_emoji" db:"_emoji"`
	Timestamp int64  `json:"_timestamp" db:"_timestamp"`
}

// NewMessageRequest represents the request to send a new message
type NewMessageRequest struct {
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Message  string `json:"message"`
}

// GetMessagesRequest represents the request to get messages
type GetMessagesRequest struct {
	Pubkey1 string `json:"pubkey"`
	Pubkey2 string `json:"pubkey2"`
}
