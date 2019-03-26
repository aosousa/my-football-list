package models

/*League represents a national football league. Fields:
 * LeagueID (int) - Unique ID of the league
 * Name (string) - Name of the league
 * Country (string) - Country of the league
 * Season (int) - Year in which the current season in the league started
 * LogoURL (string) - URL of the logo of the league
 * FlagURL (string) - URL of the flag of the league's country
 */
type League struct {
	LeagueID int    `json:"leagueId"`
	Name     string `json:"name"`
	Country  string `json:"country"`
	Season   int    `json:"season"`
	LogoURL  string `json:"logoUrl"`
	FlagURL  string `json:"flagUrl"`
}

// Leagues represents a slice of League structs
type Leagues []League
