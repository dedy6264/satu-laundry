package entities

import (
	"time"
)

type UserAccess struct {
	ID             int        `json:"id"`
	Username       string     `json:"username"`
	Password       string     `json:"password,omitempty"` // Omit when marshaling to JSON
	Role           string     `json:"role"`               //staft, cashier,warehouse, manager and owner
	IsActive       bool       `json:"is_active"`
	LastLogin      *time.Time `json:"last_login,omitempty"`
	ReferenceLevel string     `json:"reference_level"` //pegawai, outlet or cabang
	ReferenceID    int        `json:"reference_id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type CreateUserAccessRequest struct {
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required,min=6"`
	Role           string `json:"role" validate:"required"`
	IsActive       bool   `json:"is_active"`
	ReferenceLevel string `json:"reference_level"`
	ReferenceID    int    `json:"reference_id"`
}

type UpdateUserAccessRequest struct {
	Username       string `json:"username" validate:"required"`
	Role           string `json:"role" validate:"required"`
	IsActive       bool   `json:"is_active"`
	ReferenceLevel string `json:"reference_level"`
	ReferenceID    int    `json:"reference_id"`
}

type UpdateUserPasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginResponse struct {
	AccessToken  string     `json:"access_token"`
	RefreshToken string     `json:"refresh_token,omitempty"`
	User         UserAccess `json:"user"`
}
