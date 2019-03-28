package handlers

import (
	"net/http"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
	"github.com/gorilla/mux"
)

/*GetTeamFixtures queries the databse for a team's fixtures.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func GetTeamFixtures(w http.ResponseWriter, r *http.Request) {
	// check user's authentication before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		fixture      m.Fixture
		fixtures     m.Fixtures
		responseBody m.HTTPResponse
	)

	// get team ID from URL
	teamID := mux.Vars(r)["id"]

	rows, err := db.Query(`SELECT fixtureId, apiFixtureId, date, league, tbl_league.name, tbl_league.country, tbl_league.season, 
	tbl_league.logoUrl, flagUrl, round, homeTeam, homeTeam.name, homeTeam.logoUrl, homeTeamGoals, awayTeam, awayTeam.name, awayTeam.logoUrl, 
	awayTeamGoals, status, elapsed 
	FROM tbl_fixture 
	INNER JOIN tbl_league ON tbl_fixture.league = tbl_league.leagueId
	INNER JOIN tbl_team AS homeTeam ON tbl_fixture.homeTeam = homeTeam.teamId
	INNER JOIN tbl_team AS awayTeam ON tbl_fixture.awayTeam = awayTeam.teamId
	WHERE homeTeam = ` + teamID + " OR awayTeam = " + teamID)
	if err != nil {
		utils.HandleError("Fixture", "GetTeamFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
	}

	for rows.Next() {
		var (
			league                                                         m.League
			homeTeam, awayTeam                                             m.Team
			fixtureID, apiFixtureID, homeTeamGoals, awayTeamGoals, elapsed int
			date, round, status                                            string
		)

		err = rows.Scan(&fixtureID, &apiFixtureID, &date, &league.LeagueID, &league.Name, &league.Country, &league.Season, &league.LogoURL, &league.FlagURL, &round, &homeTeam.TeamID, &homeTeam.Name, &homeTeam.LogoURL, &homeTeamGoals, &awayTeam.TeamID, &awayTeam.Name, &awayTeam.LogoURL, &awayTeamGoals, &status, &elapsed)
		if err != nil {
			utils.HandleError("Fixture", "GetTeamFixtures", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
		}

		fixture = m.Fixture{
			FixtureID:     fixtureID,
			APIFixtureID:  apiFixtureID,
			Date:          date,
			League:        league,
			Round:         round,
			HomeTeam:      homeTeam,
			HomeTeamGoals: homeTeamGoals,
			AwayTeam:      awayTeam,
			AwayTeamGoals: awayTeamGoals,
			Status:        status,
			Elapsed:       elapsed,
		}

		fixtures = append(fixtures, fixture)
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    fixtures,
		Rows:    len(fixtures),
	}

	SetResponse(w, http.StatusOK, responseBody)
}
