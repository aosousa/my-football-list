package handlers

import (
	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
)

/*GetTeam queries for a Team row in the database through team ID
 * Receives:
 * teamID (int) - ID of the team
 *
 * Returns:
 * Team - Team struct found
 * err - Description of error found during execution (or nil otherwise)
 */
 func GetTeam(teamID int) (m.Team, error) {
	var team m.Team
	
	err := db.QueryRow("SELECT teamId, name, logoUrl FROM tbl_team WHERE teamId = ?", teamID).Scan(&team.TeamID, &team.Name, &team.LogoURL)
	if err != nil {
		utils.HandleError("Team", "GetTeam", err)
		return team, err
	}

	return team, nil
}