package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
	"github.com/gorilla/mux"
)

/*GetTeamFixtures queries the database for a team's fixtures.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func GetTeamFixtures(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		fixture      m.Fixture
		fixtures     m.TeamFixtures
		responseBody m.HTTPResponse
		teamID       string
	)

	// get team ID from session
	userID, err := getUserIDFromSession(r)
	if err != nil {
		utils.HandleError("Fixture", "GetTeamFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// get team ID from URL
	teamID = mux.Vars(r)["id"]
	intTeamID, err := strconv.Atoi(teamID)
	if err != nil {
		err = errors.New("Error in getting team fixtures. Please try again later.")
		responseBody.Error = err.Error()

		utils.HandleError("Fixture", "GetTeamFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// get team through ID received in URL
	team, err := GetTeam(intTeamID)
	if err != nil {
		err = errors.New("Team does not exist in the platform")
		responseBody.Error = err.Error()

		utils.HandleError("Fixture", "GetTeamFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	fixtures.Team = team

	rows, err := db.Query(`SELECT tbl_fixture.fixtureId, apiFixtureId, date, league, tbl_league.name, tbl_league.country, tbl_league.season, 
	tbl_league.logoUrl, flagUrl, round, homeTeam, homeTeam.name, homeTeam.logoUrl, homeTeamGoals, awayTeam, awayTeam.name, awayTeam.logoUrl, 
	awayTeamGoals, tbl_fixture.status, elapsed, tbl_user_fixture.status, userFixtureId 
	FROM tbl_fixture 
	INNER JOIN tbl_league ON tbl_fixture.league = tbl_league.leagueId
	INNER JOIN tbl_team AS homeTeam ON tbl_fixture.homeTeam = homeTeam.teamId
	INNER JOIN tbl_team AS awayTeam ON tbl_fixture.awayTeam = awayTeam.teamId
	LEFT JOIN tbl_user_fixture ON tbl_user_fixture.fixtureId = tbl_fixture.fixtureId
	WHERE homeTeam = ` + teamID + ` OR awayTeam = ` + teamID + ` 
	AND (userId = ` + userID + ` OR userId IS NULL)
	ORDER BY date DESC`)
	if err != nil {
		utils.HandleError("Fixture", "GetTeamFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	for rows.Next() {
		var (
			league                                                                               m.League
			homeTeam, awayTeam                                                                   m.Team
			fixtureID, apiFixtureID, homeTeamGoals, awayTeamGoals, elapsed, userFixtureStatusInt, userFixtureIDInt int
			date, round, status                                                                  string
			userFixtureStatus, userFixtureID                                                                    sql.NullInt64
		)

		err = rows.Scan(&fixtureID, &apiFixtureID, &date, &league.LeagueID, &league.Name, &league.Country, &league.Season, &league.LogoURL, &league.FlagURL, &round, &homeTeam.TeamID, &homeTeam.Name, &homeTeam.LogoURL, &homeTeamGoals, &awayTeam.TeamID, &awayTeam.Name, &awayTeam.LogoURL, &awayTeamGoals, &status, &elapsed, &userFixtureStatus, &userFixtureID)
		if err != nil {
			utils.HandleError("Fixture", "GetTeamFixtures", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
			return
		}

		if userFixtureStatus.Valid {
			userFixtureStatusInt = int(userFixtureStatus.Int64)
		}

		if userFixtureID.Valid {
			userFixtureIDInt = int(userFixtureID.Int64)
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
			UserFixtureID: userFixtureIDInt,
		}

		fixtures.Fixtures = append(fixtures.Fixtures, fixture)
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    fixtures,
		Rows:    len(fixtures.Fixtures),
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*GetDateFixtures queries the database for fixtures in a given date.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func GetDateFixtures(w http.ResponseWriter, r *http.Request) {
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
		date         string
	)

	// get user ID from session
	userID, err := getUserIDFromSession(r)
	if err != nil {
		utils.HandleError("Fixture", "GetDateFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// get date from URL
	date = mux.Vars(r)["date"]

	rows, err := db.Query(`SELECT tbl_fixture.fixtureId, apiFixtureId, date, league,
	tbl_league.name, tbl_league.country, tbl_league.season,
	tbl_league.logoUrl, flagUrl, round, homeTeam, homeTeam.name, homeTeam.logoUrl,
	homeTeamGoals, awayTeam, awayTeam.name, awayTeam.logoUrl,
	awayTeamGoals, tbl_fixture.status, elapsed, tbl_user_fixture.status, userFixtureId
	FROM tbl_fixture
	INNER JOIN tbl_league ON tbl_fixture.league = tbl_league.leagueId
	INNER JOIN tbl_team AS homeTeam ON tbl_fixture.homeTeam = homeTeam.teamId
	INNER JOIN tbl_team AS awayTeam ON tbl_fixture.awayTeam = awayTeam.teamId
	LEFT JOIN tbl_user_fixture ON tbl_user_fixture.fixtureId = tbl_fixture.fixtureId
	WHERE date LIKE '` + date + `%' AND (userId = ` + userID + ` OR userId IS NULL)
	ORDER BY apiFixtureId ASC`)
	if err != nil {
		utils.HandleError("Fixture", "GetDateFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	for rows.Next() {
		var (
			league                                                                               m.League
			homeTeam, awayTeam                                                                   m.Team
			fixtureID, apiFixtureID, homeTeamGoals, awayTeamGoals, elapsed, userFixtureStatusInt, userFixtureIDInt int
			date, round, status                                                                  string
			userFixtureStatus, userFixtureID                                                                    sql.NullInt64
		)

		err = rows.Scan(&fixtureID, &apiFixtureID, &date, &league.LeagueID, &league.Name, &league.Country, &league.Season, &league.LogoURL, &league.FlagURL, &round, &homeTeam.TeamID, &homeTeam.Name, &homeTeam.LogoURL, &homeTeamGoals, &awayTeam.TeamID, &awayTeam.Name, &awayTeam.LogoURL, &awayTeamGoals, &status, &elapsed, &userFixtureStatus, &userFixtureID)
		if err != nil {
			utils.HandleError("Fixture", "GetTeamFixtures", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
			return
		}

		if userFixtureStatus.Valid {
			userFixtureStatusInt = int(userFixtureStatus.Int64)
		}

		if userFixtureID.Valid {
			userFixtureIDInt = int(userFixtureID.Int64)
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
			UserFixtureID: userFixtureIDInt,
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
