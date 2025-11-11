package handlers

import (
	"github.com/beenzer/beenzer-server/v3/internal/database"
	"github.com/beenzer/beenzer-server/v3/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	repo *database.UserRepository
}

func NewUserHandler(db *database.DB) *UserHandler {
	return &UserHandler{
		repo: database.NewUserRepository(db),
	}
}

// GetUser godoc
// @Summary Get user by public key
// @Description Retrieve user profile information
// @Tags users
// @Accept json
// @Produce json
// @Param pubkey path string true "User public key"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/users/{pubkey} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	pubkey := c.Params("pubkey")
	if pubkey == "" {
		return c.Status(400).JSON(fiber.Map{"error": "pubkey is required"})
	}

	user, err := h.repo.GetUser(c.Context(), pubkey)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}

	return c.JSON(user)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Register a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body map[string]interface{} true "User creation data"
// @Success 201 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Router /api/users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var body struct {
		Pubkey   string `json:"pubkey"`
		Username string `json:"username"`
		AppUser  bool   `json:"appuser"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	// Validate input
	if len(body.Username) < 3 {
		return c.Status(400).JSON(fiber.Map{"error": "username must be at least 3 characters"})
	}

	// Sanitize username
	username := utils.SQLFilter(body.Username)
	if len(username) < 3 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid username"})
	}

	// Check if username is taken
	taken, err := h.repo.IsUsernameTaken(c.Context(), username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "database error"})
	}
	if taken {
		return c.Status(400).JSON(fiber.Map{"error": "username already exists"})
	}

	// Create user
	err = h.repo.CreateUser(c.Context(), body.Pubkey, username, body.AppUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create user"})
	}

	return c.Status(201).JSON(fiber.Map{"success": true})
}

// UpdateUser godoc
// @Summary Update user information
// @Description Update a specific field of user profile
// @Tags users
// @Accept json
// @Produce json
// @Param pubkey path string true "User public key"
// @Param update body map[string]interface{} true "Update data"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Router /api/users/{pubkey} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	pubkey := c.Params("pubkey")
	
	var body struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	// Sanitize value
	value := utils.SQLFilter(body.Value)
	if body.Field == "_pfp" || body.Field == "_description" {
		// Allow these fields to have special characters
		value = body.Value
	}

	err := h.repo.UpdateUser(c.Context(), pubkey, body.Field, value)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

// SearchUsers godoc
// @Summary Search users by username
// @Description Search for users matching the query string
// @Tags users
// @Accept json
// @Produce json
// @Param query path string true "Search query"
// @Success 200 {array} models.User
// @Failure 400 {object} map[string]string
// @Router /api/users/search/{query} [get]
func (h *UserHandler) SearchUsers(c *fiber.Ctx) error {
	query := c.Params("query")
	if len(query) < 3 {
		return c.Status(400).JSON(fiber.Map{"error": "query must be at least 3 characters"})
	}

	query = utils.SQLFilter(query)
	users, err := h.repo.SearchUsers(c.Context(), query)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "search failed"})
	}

	return c.JSON(users)
}

// CheckUsername godoc
// @Summary Check username availability
// @Description Check if a username is available
// @Tags users
// @Accept json
// @Produce json
// @Param username path string true "Username to check"
// @Success 200 {object} map[string]bool
// @Router /api/users/check/{username} [get]
func (h *UserHandler) CheckUsername(c *fiber.Ctx) error {
	username := c.Params("username")
	username = utils.SQLFilter(username)

	taken, err := h.repo.IsUsernameTaken(c.Context(), username)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "database error"})
	}

	return c.JSON(fiber.Map{"available": !taken})
}

// AddFriend godoc
// @Summary Add a friend
// @Description Create a friendship connection between two users
// @Tags users
// @Accept json
// @Produce json
// @Param pubkey path string true "User public key"
// @Param friendPubkey path string true "Friend public key"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Router /api/users/{pubkey}/friends/{friendPubkey} [post]
func (h *UserHandler) AddFriend(c *fiber.Ctx) error {
	pubkey := c.Params("pubkey")
	friendPubkey := c.Params("friendPubkey")

	if len(pubkey) < 22 || len(friendPubkey) < 22 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid public keys"})
	}

	err := h.repo.AddFriend(c.Context(), pubkey, friendPubkey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to add friend"})
	}

	return c.JSON(fiber.Map{"success": true})
}

// RemoveFriend godoc
// @Summary Remove a friend
// @Description Remove a friendship connection
// @Tags users
// @Accept json
// @Produce json
// @Param pubkey path string true "User public key"
// @Param friendPubkey path string true "Friend public key"
// @Success 200 {object} map[string]bool
// @Router /api/users/{pubkey}/friends/{friendPubkey} [delete]
func (h *UserHandler) RemoveFriend(c *fiber.Ctx) error {
	pubkey := c.Params("pubkey")
	friendPubkey := c.Params("friendPubkey")

	err := h.repo.RemoveFriend(c.Context(), pubkey, friendPubkey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to remove friend"})
	}

	return c.JSON(fiber.Map{"success": true})
}

// GetUserFriends godoc
// @Summary Get user's friends
// @Description Retrieve all friends of a user with full profile information
// @Tags users
// @Accept json
// @Produce json
// @Param pubkey path string true "User public key"
// @Success 200 {array} models.User
// @Router /api/users/{pubkey}/friends [get]
func (h *UserHandler) GetUserFriends(c *fiber.Ctx) error {
	pubkey := c.Params("pubkey")

	friends, err := h.repo.GetUserFriends(c.Context(), pubkey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get friends"})
	}

	return c.JSON(friends)
}

// GetUserLogs godoc
// @Summary Get user activity logs
// @Description Retrieve user activity history
// @Tags users
// @Accept json
// @Produce json
// @Param pubkey path string true "User public key"
// @Success 200 {array} models.UserLog
// @Router /api/users/{pubkey}/logs [get]
func (h *UserHandler) GetUserLogs(c *fiber.Ctx) error {
	pubkey := c.Params("pubkey")

	logs, err := h.repo.GetLogs(c.Context(), pubkey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to get logs"})
	}

	return c.JSON(logs)
}

// CheckNewUser godoc
// @Summary Check if user is new
// @Description Check if a public key is registered
// @Tags users
// @Accept json
// @Produce json
// @Param pubkey path string true "User public key"
// @Success 200 {object} map[string]bool
// @Router /api/users/{pubkey}/new [get]
func (h *UserHandler) CheckNewUser(c *fiber.Ctx) error {
	pubkey := c.Params("pubkey")

	isNew, err := h.repo.IsNewUser(c.Context(), pubkey)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "database error"})
	}

	return c.JSON(fiber.Map{"isNew": isNew})
}
