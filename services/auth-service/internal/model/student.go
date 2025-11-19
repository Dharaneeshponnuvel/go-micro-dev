package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Student struct {
	UserID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	BatchID    uuid.UUID `gorm:"type:uuid;index" json:"batch_id"`
	RollNumber string    `gorm:"type:varchar(50);uniqueIndex:idx_batch_roll" json:"roll_number"` // composite unique with batch_id
	Branch     string    `gorm:"type:varchar(100)" json:"branch"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`

	Demographics datatypes.JSON `gorm:"type:jsonb" json:"demographics,omitempty"`

	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	User  User  `gorm:"foreignKey:UserID" json:"user"`
	Batch Batch `gorm:"foreignKey:BatchID" json:"batch"`
}
