package models

// NFT represents an NFT in the system
type NFT struct {
	ID          int64   `json:"_id_" db:"_id_"`
	Token       string  `json:"__token__" db:"__token__"`
	Supply      int     `json:"_supply" db:"_supply"`
	Floor       float64 `json:"_floor" db:"_floor"`
	CCY         string  `json:"_ccy" db:"_ccy"`
	Creator     string  `json:"_creator" db:"_creator"`
	Username    string  `json:"_username" db:"_username"`
	Image       string  `json:"_image" db:"_image"`
	Asset       string  `json:"_asset" db:"_asset"`
	Type        string  `json:"_type" db:"_type"`
	Metadata    string  `json:"_metadata" db:"_metadata"`
	Name        string  `json:"_name" db:"_name"`
	Description string  `json:"_description" db:"_description"`
	City        string  `json:"_city" db:"_city"`
	Latitude    float64 `json:"_latitude" db:"_latitude"`
	Longitude   float64 `json:"_longitude" db:"_longitude"`
	Visibility  string  `json:"_visibility" db:"_visibility"`
	MaxLat      float64 `json:"_maxLat" db:"_maxLat"`
	MinLat      float64 `json:"_minLat" db:"_minLat"`
	MaxLon      float64 `json:"_maxLon" db:"_maxLon"`
	MinLon      float64 `json:"_minLon" db:"_minLon"`
	Date        string  `json:"_date" db:"_date"`
	Time        string  `json:"_time" db:"_time"`
	Timestamp   int64   `json:"_timestamp" db:"_timestamp"`
}

// NFTEdition represents an edition of an NFT
type NFTEdition struct {
	Master    string `json:"__master__" db:"__master__"`
	Edition   string `json:"__edition__" db:"__edition__"`
	Minter    string `json:"_minter" db:"_minter"`
	ID        int    `json:"_id" db:"_id"`
	Date      string `json:"_date" db:"_date"`
	Time      string `json:"_time" db:"_time"`
	Timestamp int64  `json:"_timestamp" db:"_timestamp"`
}

// NFTCounter represents the NFT counter
type NFTCounter struct {
	N         int64 `json:"_n" db:"_n"`
	Timestamp int64 `json:"_timestamp" db:"_timestamp"`
}

// NFTOwner represents NFT ownership
type NFTOwner struct {
	Token     string `json:"_token" db:"_token"`
	Owner     string `json:"_owner" db:"_owner"`
	Timestamp int64  `json:"_timestamp" db:"_timestamp"`
}

// NFTTransaction represents an NFT transaction
type NFTTransaction struct {
	Owner     string  `json:"_owner" db:"_owner"`
	Pubkey    string  `json:"_pubkey" db:"_pubkey"`
	Type      string  `json:"_type" db:"_type"`
	Currency  string  `json:"_currency" db:"_currency"`
	Amount    float64 `json:"_amount" db:"_amount"`
	Hash      string  `json:"_hash" db:"_hash"`
	Timestamp int64   `json:"_timestamp" db:"_timestamp"`
}

// MintNFTRequest represents the request to mint an NFT
type MintNFTRequest struct {
	Asset       []byte  `json:"asset"`
	Type        string  `json:"type"`
	Creator     string  `json:"creator"`
	Supply      int     `json:"supply"`
	Floor       float64 `json:"floor"`
	Username    string  `json:"username"`
	Description string  `json:"description"`
	City        string  `json:"city"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Visibility  string  `json:"visibility"`
	MaxLat      float64 `json:"maxLat"`
	MinLat      float64 `json:"minLat"`
	MaxLon      float64 `json:"maxLon"`
	MinLon      float64 `json:"minLon"`
	Image       []byte  `json:"image,omitempty"`
	Royalties   int     `json:"royalties,omitempty"`
	MintCCY     string  `json:"mintCcy,omitempty"`
}
