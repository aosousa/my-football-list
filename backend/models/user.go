package models

import "database/sql"

// User represents a user in the platform
type User struct {
	UserID                     int    `json:"userId"`   // Unique ID of the user in the platform
	Username                   string `json:"username"` // User's username in the platform
	Password                   string `json:"password"` // User's password in the platform
	ConfirmPassword 		   string `json:"currentPassword"` // User's password confirmation
	PasswordResetToken         sql.NullString // Token to use for password reset
	PasswordResetTokenValidity sql.NullString // Validity of a user's password reset token
	Email                      string `json:"email"`       // User's email (used for password reset)
	CreateTime                 string `json:"createTime"`  // Timestamp of when the account was created
	UpdateTime                 string `json:"updateTime"`  // Timestamp of when the account was last updated
	Status                     int    `json:"status"`      // Status of the account
	SpoilerMode                bool   `json:"spoilerMode"` // Whether or not to show results for this user (0 = Show, 1 = Hide)
}

// Users represents a slice of User structs
type Users []User

// ChangePasswordRequest is the struct used for a password change request
type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword"` // User's current password
	NewPassword string `json:"newPassword"` // User's new password
	ConfirmNewPassword string `json:"confirmNewPassword"` // User's confirmation of new password
}