package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();uniqueIndex"`
	Name     string    `gorm:"type:varchar(255);default:uuid_generate_v4();not null"`
	Phone    string    `gorm:"type:varchar(17);uniqueIndex"`
	Email    string    `gorm:"type:varchar(255);uniqueIndex"`
	Password string    `gorm:"not null"`
	Role     string    `gorm:"type:varchar(255);default:'user';not null"`
	Provider string    `gorm:"type:varchar(255)"`
	Photo    string    `gorm:"type:varchar(255)"`

	// Authentication part
	VerificationCode   string `gorm:"type:varchar(255)"`
	PasswordResetToken string `gorm:"type:varchar(255)"`
	PasswordResetAt    time.Time
	Verified           bool `gorm:"default:false;not null"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Phone           string `json:"phone" binding:"required"`
	Email           string `json:"email"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,min=8"`
	Photo           string `json:"photo"`
}

type SignInInput struct {
	Phone    string `json:"phone"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID   uuid.UUID `json:"id,omitempty"`
	Name string    `json:"name,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ForgotPasswordInput struct {
	// Make Phone and Email json field optional, however one of each should be required
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}
