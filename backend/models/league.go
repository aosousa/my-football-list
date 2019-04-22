package models

// League represents a national football league
type League struct {
	LeagueID int    `json:"leagueId"` // Unique ID of the league
	Name     string `json:"name"`     // Name of the league
	Country  string `json:"country"`  // Country of the league
	Season   int    `json:"season"`   // Year in which the current season in the league started
	LogoURL  string `json:"logoUrl"`  // URL of the logo of the league
	FlagURL  string `json:"flagUrl"`  // URL of the flag of the league's country
}

// Leagues represents a slice of League structs
type Leagues []League
