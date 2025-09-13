package entities

import (
	"time"
)

type EmployeeAccess struct {
	ID            int        `json:"id"`
	EmployeeID    int        `json:"id_pegawai"`
	Username      string     `json:"username"`
	Password      string     `json:"password,omitempty"` // Omit when marshaling to JSON
	Role          string     `json:"role"`
	IsActive      bool       `json:"is_active"`
	LastLogin     *time.Time `json:"last_login,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	EmployeeName  string     `json:"employee_name,omitempty"`
	EmployeeEmail string     `json:"employee_email,omitempty"`
	OutletID      int        `json:"id_outlet,omitempty"`
}

type CreateEmployeeAccessRequest struct {
	EmployeeID int    `json:"id_pegawai" validate:"required"`
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required,min=6"`
	Role       string `json:"role" validate:"required"`
	IsActive   bool   `json:"is_active"`
}

type UpdateEmployeeAccessRequest struct {
	Username string `json:"username" validate:"required"`
	Role     string `json:"role" validate:"required"`
	IsActive bool   `json:"is_active"`
}

type UpdateEmployeePasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}

type EmployeeLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type EmployeeLoginResponse struct {
	AccessToken  string         `json:"access_token"`
	RefreshToken string         `json:"refresh_token,omitempty"`
	User         EmployeeAccess `json:"user"`
}
