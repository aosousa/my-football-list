package models

// Team represents a football team
type Team struct {
	TeamID  int    `json:"teamId"`  // Unique ID of the team
	Name    string `json:"name"`    // Name of the team
	LogoURL string `json:"logoUrl"` // URL of the logo of the team
}
