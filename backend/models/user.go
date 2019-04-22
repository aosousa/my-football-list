package models

// User represents a user in the platform
type User struct {
	UserID                     int    `json:"userId"`   // Unique ID of the user in the platform
	Username                   string `json:"username"` // User's username in the platform
	Password                   string `json:"password"` // User's password in the platform
	PasswordResetToken         string // Token to use for password reset
	PasswordResetTokenValidity string // Validity of a user's password reset token
	Email                      string `json:"email"`       // User's email (used for password reset)
	CreateTime                 string `json:"createTime"`  // Timestamp of when the account was created
	UpdateTime                 string `json:"updateTime"`  // Timestamp of when the account was last updated
	Status                     int    `json:"status"`      // Status of the account
	SpoilerMode                bool   `json:"spoilerMode"` // Whether or not to show results for this user (0 = Show, 1 = Hide)
}

// Users represents a slice of User structs
type Users []User
