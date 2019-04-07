package handlers

import (
	"database/sql"
	"net/http"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
	"github.com/gorilla/mux"
)

/*GetAllLeagues queries the database for all the leagues available.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func GetAllLeagues(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		league       m.League
		leagues      m.Leagues
		responseBody m.HTTPResponse
	)

	rows, err := db.Query("SELECT * FROM tbl_league ORDER BY leagueId ASC")
	if err != nil {
		utils.HandleError("League", "GetAllLeagues", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
	}

	for rows.Next() {
		var (
			id, season                      int
			name, country, logoURL, flagURL string
		)

		err = rows.Scan(&id, &name, &country, &season, &logoURL, &flagURL)
		if err != nil {
			utils.HandleError("League", "GetAllLeagues", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
		}

		league = m.League{
			LeagueID: id,
			Name:     name,
			Country:  country,
			Season:   season,
			LogoURL:  logoURL,
			FlagURL:  flagURL,
		}

		leagues = append(leagues, league)
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    leagues,
		Rows:    len(leagues),
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*GetLeagueFixtures queries the database for a league's fixtures.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func GetLeagueFixtures(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		fixture      m.Fixture
		fixtures     m.Fixtures
		responseBody m.HTTPResponse
		leagueID     string
	)

	// get user ID from session
	userID, err := getUserIDFromSession(r)
	if err != nil {
		utils.HandleError("League", "GetLeagueFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// get league ID from URL
	leagueID = mux.Vars(r)["id"]

	rows, err := db.Query(`SELECT tbl_fixture.fixtureId, apiFixtureId, date, league, tbl_league.name, tbl_league.country, tbl_league.season, 
	tbl_league.logoUrl, flagUrl, round, homeTeam, homeTeam.name, homeTeam.logoUrl, homeTeamGoals, awayTeam, awayTeam.name, awayTeam.logoUrl, 
	awayTeamGoals, tbl_fixture.status, elapsed, tbl_user_fixture.status 
	FROM tbl_fixture 
	INNER JOIN tbl_league ON tbl_fixture.league = tbl_league.leagueId
	INNER JOIN tbl_team AS homeTeam ON tbl_fixture.homeTeam = homeTeam.teamId
	INNER JOIN tbl_team AS awayTeam ON tbl_fixture.awayTeam = awayTeam.teamId
	LEFT JOIN tbl_user_fixture ON tbl_user_fixture.fixtureId = tbl_fixture.fixtureId
	WHERE league = ` + leagueID + ` AND (userId = ` + userID + ` OR userId IS NULL)
	ORDER BY apiFixtureId ASC`)
	if err != nil {
		utils.HandleError("League", "GetLeagueFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	for rows.Next() {
		var (
			league                                                                               m.League
			homeTeam, awayTeam                                                                   m.Team
			fixtureID, apiFixtureID, homeTeamGoals, awayTeamGoals, elapsed, userFixtureStatusInt int
			date, round, status                                                                  string
			userFixtureStatus                                                                    sql.NullInt64
		)

		err = rows.Scan(&fixtureID, &apiFixtureID, &date, &league.LeagueID, &league.Name, &league.Country, &league.Season, &league.LogoURL, &league.FlagURL, &round, &homeTeam.TeamID, &homeTeam.Name, &homeTeam.LogoURL, &homeTeamGoals, &awayTeam.TeamID, &awayTeam.Name, &awayTeam.LogoURL, &awayTeamGoals, &status, &elapsed, &userFixtureStatus)
		if err != nil {
			utils.HandleError("League", "GetLeagueFixtures", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
			return
		}

		if userFixtureStatus.Valid {
			userFixtureStatusInt = int(userFixtureStatus.Int64)
		}

		fixture = m.Fixture{
			FixtureID:         fixtureID,
			APIFixtureID:      apiFixtureID,
			Date:              date,
			League:            league,
			Round:             round,
			HomeTeam:          homeTeam,
			HomeTeamGoals:     homeTeamGoals,
			AwayTeam:          awayTeam,
			AwayTeamGoals:     awayTeamGoals,
			Status:            status,
			Elapsed:           elapsed,
			UserFixtureStatus: userFixtureStatusInt,
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
