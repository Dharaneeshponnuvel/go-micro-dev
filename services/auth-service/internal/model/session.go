package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Session struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID       uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	Token        string         `gorm:"type:text;not null;uniqueIndex" json:"token"` // JWT token or hashed token
	RefreshToken string         `gorm:"type:text;uniqueIndex" json:"refresh_token"`
	DeviceInfo   datatypes.JSON `gorm:"type:jsonb" json:"device_info"` // Device fingerprint, IP, user agent
	ExpiresAt    time.Time      `gorm:"not null;index" json:"expires_at"`
	IsActive     bool           `gorm:"default:true;index" json:"is_active"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"user"`
}
