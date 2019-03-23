package models

/*Team represents a football team. Fields:
 * TeamID (int) - Unique ID of the team
 * Name (string) - Name of the team
 * LogoURL (string) - URL of the logo of the team
 */
type Team struct {
	TeamID  int    `json:"teamId"`
	Name    string `json:"name"`
	LogoURL string `json:"logoUrl"`
}
