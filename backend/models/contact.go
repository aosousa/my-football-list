package models

// Contact is the struct used to receive information from the contact form
type Contact struct {
	Type    string `json:"type"`    // Type of request (bug fix, suggestion, etc)
	Subject string `json:"subject"` // Message subject
	Message string `json:"message"` // Message body
}
