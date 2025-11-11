package database

import (
	"context"
	"fmt"

	"github.com/beenzer/beenzer-server/v3/internal/models"
	"github.com/beenzer/beenzer-server/v3/internal/utils"
)

// UserRepository handles user database operations
type UserRepository struct {
	db *DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUser retrieves a user by public key
func (r *UserRepository) GetUser(ctx context.Context, pubkey string) (*models.User, error) {
	query := `SELECT * FROM users WHERE __pubkey__ = $1`
	
	var user models.User
	err := r.db.Users.QueryRow(ctx, query, pubkey).Scan(
		&user.Pubkey, &user.Username, &user.PFP, &user.Name, &user.Lastname,
		&user.Description, &user.Birthdate, &user.Country, &user.Flag, &user.City,
		&user.Phone, &user.Email, &user.Verified, &user.Twitter, &user.Instagram,
		&user.Discord, &user.Telegram, &user.YouTube, &user.TikTok,
		&user.MagicEden, &user.OpenSea, &user.AppUser, &user.CreatedAt, &user.Timestamp,
	)
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

// CreateUser creates a new user
func (r *UserRepository) CreateUser(ctx context.Context, pubkey, username string, appuser bool) error {
	timestamp := utils.GetTimestamp()
	query := `
		INSERT INTO users (__pubkey__, _username_, _appuser, _created_at, _timestamp)
		VALUES ($1, $2, $3, $4, $5)
	`
	
	_, err := r.db.Users.Exec(ctx, query, pubkey, username, appuser, timestamp, timestamp)
	return err
}

// UpdateUser updates a user field
func (r *UserRepository) UpdateUser(ctx context.Context, pubkey, field, value string) error {
	// Sanitize field name to prevent SQL injection
	allowedFields := map[string]bool{
		"_username_": true, "_pfp": true, "_name": true, "_lastname": true,
		"_description": true, "_birthdate": true, "_country": true, "_flag": true,
		"_city": true, "_phone": true, "_email": true, "_verified": true,
		"_twitter": true, "_instagram": true, "_discord": true, "_telegram": true,
		"_youtube": true, "_tiktok": true, "_magiceden": true, "_opensea": true,
	}
	
	if !allowedFields[field] {
		return fmt.Errorf("invalid field: %s", field)
	}
	
	timestamp := utils.GetTimestamp()
	query := fmt.Sprintf(`UPDATE users SET %s = $1, _timestamp = $2 WHERE __pubkey__ = $3`, field)
	
	_, err := r.db.Users.Exec(ctx, query, value, timestamp, pubkey)
	return err
}

// IsNewUser checks if a user exists
func (r *UserRepository) IsNewUser(ctx context.Context, pubkey string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE __pubkey__ = $1`
	
	var count int
	err := r.db.Users.QueryRow(ctx, query, pubkey).Scan(&count)
	if err != nil {
		return false, err
	}
	
	return count == 0, nil
}

// IsUsernameTaken checks if a username exists
func (r *UserRepository) IsUsernameTaken(ctx context.Context, username string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE _username_ = $1`
	
	var count int
	err := r.db.Users.QueryRow(ctx, query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}

// SearchUsers searches for users by username
func (r *UserRepository) SearchUsers(ctx context.Context, searchQuery string) ([]models.User, error) {
	query := `
		SELECT * FROM users 
		WHERE _username_ ILIKE $1 
		ORDER BY _username_ 
		LIMIT 50
	`
	
	rows, err := r.db.Users.Query(ctx, query, "%"+searchQuery+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.Pubkey, &user.Username, &user.PFP, &user.Name, &user.Lastname,
			&user.Description, &user.Birthdate, &user.Country, &user.Flag, &user.City,
			&user.Phone, &user.Email, &user.Verified, &user.Twitter, &user.Instagram,
			&user.Discord, &user.Telegram, &user.YouTube, &user.TikTok,
			&user.MagicEden, &user.OpenSea, &user.AppUser, &user.CreatedAt, &user.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, nil
}

// AddFriend adds a friendship connection
func (r *UserRepository) AddFriend(ctx context.Context, pubkey1, pubkey2 string) error {
	timestamp := utils.GetTimestamp()
	query := `INSERT INTO friends (__pubkey__, __pubkey2__, _timestamp) VALUES ($1, $2, $3)`
	
	_, err := r.db.Users.Exec(ctx, query, pubkey1, pubkey2, timestamp)
	return err
}

// RemoveFriend removes a friendship connection
func (r *UserRepository) RemoveFriend(ctx context.Context, pubkey1, pubkey2 string) error {
	query := `DELETE FROM friends WHERE __pubkey__ = $1 AND __pubkey2__ = $2`
	
	_, err := r.db.Users.Exec(ctx, query, pubkey1, pubkey2)
	return err
}

// IsFriend checks if two users are friends
func (r *UserRepository) IsFriend(ctx context.Context, pubkey1, pubkey2 string) (bool, error) {
	query := `SELECT COUNT(*) FROM friends WHERE __pubkey__ = $1 AND __pubkey2__ = $2`
	
	var count int
	err := r.db.Users.QueryRow(ctx, query, pubkey1, pubkey2).Scan(&count)
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}

// GetUserFriends retrieves all friends of a user with full details
func (r *UserRepository) GetUserFriends(ctx context.Context, pubkey string) ([]models.User, error) {
	// First get friend pubkeys
	query1 := `SELECT __pubkey2__ FROM friends WHERE __pubkey__ = $1`
	rows, err := r.db.Users.Query(ctx, query1, pubkey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var friendPubkeys []string
	for rows.Next() {
		var friendPubkey string
		if err := rows.Scan(&friendPubkey); err != nil {
			return nil, err
		}
		friendPubkeys = append(friendPubkeys, friendPubkey)
	}
	
	// Then get full user details for each friend
	friends := make([]models.User, 0, len(friendPubkeys))
	for _, friendPubkey := range friendPubkeys {
		user, err := r.GetUser(ctx, friendPubkey)
		if err != nil {
			continue // Skip if user not found
		}
		friends = append(friends, *user)
	}
	
	return friends, nil
}

// AddLog adds a user activity log
func (r *UserRepository) AddLog(ctx context.Context, pubkey, log string) error {
	timestamp := utils.GetTimestamp()
	query := `INSERT INTO logs (_pubkey, _logs, _timestamp) VALUES ($1, $2, $3)`
	
	_, err := r.db.Users.Exec(ctx, query, pubkey, log, timestamp)
	return err
}

// GetLogs retrieves user activity logs
func (r *UserRepository) GetLogs(ctx context.Context, pubkey string) ([]models.UserLog, error) {
	query := `SELECT * FROM logs WHERE _pubkey = $1 ORDER BY _timestamp DESC LIMIT 100`
	
	rows, err := r.db.Users.Query(ctx, query, pubkey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var logs []models.UserLog
	for rows.Next() {
		var log models.UserLog
		err := rows.Scan(&log.Pubkey, &log.Logs, &log.Timestamp)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	
	return logs, nil
}
