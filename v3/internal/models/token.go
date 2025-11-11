package models

// TokenTransaction represents a token transaction
type TokenTransaction struct {
	Date      string  `json:"_date" db:"_date"`
	Time      string  `json:"_time" db:"_time"`
	Type      string  `json:"_type" db:"_type"`
	Amount    float64 `json:"_amount" db:"_amount"`
	Pubkey    string  `json:"_pubkey" db:"_pubkey"`
	Flag      string  `json:"_flag" db:"_flag"`
	Timestamp int64   `json:"_timestamp" db:"_timestamp"`
}

// TokenHolder represents a token holder
type TokenHolder struct {
	Position   int     `json:"__position__" db:"__position__"`
	Percentage float64 `json:"_percentage" db:"_percentage"`
	Amount     float64 `json:"_amount" db:"_amount"`
	Pubkey     string  `json:"_pubkey" db:"_pubkey"`
	Flag       string  `json:"_flag" db:"_flag"`
	Timestamp  int64   `json:"_timestamp" db:"_timestamp"`
}
