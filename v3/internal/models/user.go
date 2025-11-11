package models

// User represents a user in the system
type User struct {
	Pubkey      string `json:"__pubkey__" db:"__pubkey__"`
	Username    string `json:"_username_" db:"_username_"`
	PFP         string `json:"_pfp" db:"_pfp"`
	Name        string `json:"_name" db:"_name"`
	Lastname    string `json:"_lastname" db:"_lastname"`
	Description string `json:"_description" db:"_description"`
	Birthdate   string `json:"_birthdate" db:"_birthdate"`
	Country     string `json:"_country" db:"_country"`
	Flag        string `json:"_flag" db:"_flag"`
	City        string `json:"_city" db:"_city"`
	Phone       string `json:"_phone" db:"_phone"`
	Email       string `json:"_email" db:"_email"`
	Verified    bool   `json:"_verified" db:"_verified"`
	Twitter     string `json:"_twitter" db:"_twitter"`
	Instagram   string `json:"_instagram" db:"_instagram"`
	Discord     string `json:"_discord" db:"_discord"`
	Telegram    string `json:"_telegram" db:"_telegram"`
	YouTube     string `json:"_youtube" db:"_youtube"`
	TikTok      string `json:"_tiktok" db:"_tiktok"`
	MagicEden   string `json:"_magiceden" db:"_magiceden"`
	OpenSea     string `json:"_opensea" db:"_opensea"`
	AppUser     bool   `json:"_appuser" db:"_appuser"`
	CreatedAt   int64  `json:"_created_at" db:"_created_at"`
	Timestamp   int64  `json:"_timestamp" db:"_timestamp"`
}

// UserLog represents a user activity log
type UserLog struct {
	Pubkey    string `json:"_pubkey" db:"_pubkey"`
	Logs      string `json:"_logs" db:"_logs"`
	Timestamp int64  `json:"_timestamp" db:"_timestamp"`
}

// Friend represents a friendship connection between two users
type Friend struct {
	Pubkey1   string `json:"__pubkey__" db:"__pubkey__"`
	Pubkey2   string `json:"__pubkey2__" db:"__pubkey2__"`
	Timestamp int64  `json:"_timestamp" db:"_timestamp"`
}

// UserSearchResult represents a simplified user for search results
type UserSearchResult struct {
	Pubkey   string `json:"__pubkey__"`
	Username string `json:"_username_"`
	PFP      string `json:"_pfp"`
	Verified bool   `json:"_verified"`
}
