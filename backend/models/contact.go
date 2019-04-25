package models

// Contact is the struct used to receive information from the contact form
type Contact struct {
	Type    string `json:"type"`    // Type of request (bug fix, suggestion, etc)
	Subject string `json:"subject"` // Message subject
	Message string `json:"message"` // Message body
}

// ResetPassword is the struct used to receive the recipient's email address for a reset password request
type ResetPassword struct {
	Token string `json:"token"` // Password reset token previously sent to recipient's email address
	Email string `json:"email"` // Recipient email address
	Password string `json:"password"` // User's new password
}