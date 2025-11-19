package handlers

import (
	"strings"
	"time"

	"auth-service/internal/auth"
	"auth-service/internal/config"
	"auth-service/internal/models"
	"auth-service/internal/redis"

	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{cfg: cfg}
}

// -------------------- REGISTER --------------------

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"` // ADMIN, STUDENT, TEACHER, INSTITUTION
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var body registerRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	// Fetch Role from DB
	var role models.Role
	if err := db.Where("name = ?", body.Role).First(&role).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid role name"})
	}

	user := models.User{
		Email:  body.Email,
		Name:   body.Name,
		RoleID: role.ID,
	}

	if err := user.SetPassword(body.Password); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to set password"})
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(409).JSON(fiber.Map{"error": "email already exists"})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"user": fiber.Map{
			"id":     user.ID,
			"email":  user.Email,
			"name":   user.Name,
			"role":   role.Name,
			"roleId": user.RoleID,
		},
	})
}

// -------------------- LOGIN --------------------

type loginRequest struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	DeviceInfo string `json:"device_info"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var body loginRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	var user models.User
	if err := db.Preload("Role").Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	if !user.CheckPassword(body.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// STUDENT RESTRICTION — ONLY 1 ACTIVE DEVICE
	if user.Role.Name == "STUDENT" {
		db.Model(&models.Session{}).
			Where("user_id = ? AND is_active = true", user.ID).
			Updates(map[string]interface{}{"is_active": false})

		// Also remove any old Redis session
		redis.RDB.Del(redis.Ctx, "session:"+user.ID.String())
	}

	// Generate JWT access token
	accessToken, err := auth.GenerateAccessToken(h.cfg, user.ID.String(), user.Email, user.Role.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to generate token"})
	}

	// Create session in DB
	session := models.Session{
		UserID:     user.ID,
		Token:      accessToken,
		DeviceInfo: datatypes.JSON([]byte(body.DeviceInfo)),
		ExpiresAt:  time.Now().Add(time.Hour * 1),
		IsActive:   true,
	}
	db.Create(&session)

	// Store in Redis (Fast validation)
	sessionKey := "session:" + user.ID.String()
	redis.RDB.Set(redis.Ctx, sessionKey, accessToken, time.Hour*1)

	return c.JSON(fiber.Map{
		"success":      true,
		"access_token": accessToken,
		"user": fiber.Map{
			"id":     user.ID,
			"email":  user.Email,
			"name":   user.Name,
			"role":   user.Role.Name,
			"roleId": user.RoleID,
		},
	})
}

// -------------------- LOGOUT --------------------

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*auth.Claims)
	db := c.Locals("db").(*gorm.DB)

	// 1) Delete from Redis
	sessionKey := "session:" + claims.UserID
	redis.RDB.Del(redis.Ctx, sessionKey)

	// 2) Mark session inactive in DB
	db.Model(&models.Session{}).
		Where("user_id = ? AND token = ?", claims.UserID, extractToken(c.Get("Authorization"))).
		Updates(map[string]interface{}{"is_active": false})

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Logged out successfully",
	})
}

// -------------------- REFRESH TOKEN --------------------

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	refresh := c.FormValue("refresh_token")
	if refresh == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Missing refresh_token"})
	}

	var token models.RefreshToken
	if err := db.Where("token = ? AND is_used = false", refresh).First(&token).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid or expired refresh token"})
	}

	// Mark refresh token as used
	db.Model(&models.RefreshToken{}).Where("id = ?", token.ID).
		Update("is_used", true)

	// Get user
	var user models.User
	db.First(&user, "id = ?", token.UserID)

	// Generate new JWT
	accessToken, err := auth.GenerateAccessToken(h.cfg, user.ID.String(), user.Email, user.Role.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Token generation failed"})
	}

	// Create new refresh token
	newRefresh := models.RefreshToken{
		UserID:    user.ID,
		SessionID: token.SessionID,
		Token:     auth.GenerateSecureToken(), // make sure this exists
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}
	db.Create(&newRefresh)

	return c.JSON(fiber.Map{
		"success":       true,
		"access_token":  accessToken,
		"refresh_token": newRefresh.Token,
	})
}

// -------------------- VERIFY --------------------

func (h *AuthHandler) Verify(c *fiber.Ctx) error {
	claims := c.Locals("claims").(*auth.Claims)
	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"valid": true,
			"user": fiber.Map{
				"id":    claims.UserID,
				"email": claims.Email,
				"role":  claims.Role,
			},
		},
	})
}

// --------------- UTILITY FUNCTION ----------------

func extractToken(header string) string {
	if header == "" {
		return ""
	}
	parts := strings.Split(header, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}
