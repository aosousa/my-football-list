package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
	"github.com/gorilla/mux"
)

/*GetUserFixtures queries the database for the list of fixtures watched by a user.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func GetUserFixtures(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		userFixture  m.UserFixture
		userFixtures m.UserFixtures
		responseBody m.HTTPResponse
		userID       string
	)

	// get user ID from URL
	userID = mux.Vars(r)["id"]

	rows, err := db.Query(`SELECT userFixtureId, userId, fixture.fixtureId, tbl_user_fixture.status, date, league, round, homeTeam, homeTeam.name, homeTeam.logoUrl, homeTeamGoals,
	awayTeam, awayTeam.name, awayTeam.logoUrl, awayTeamGoals, fixture.status, elapsed, tbl_league.name, tbl_league.country, tbl_league.logoUrl, tbl_league.flagUrl
	FROM footballtracker.tbl_user_fixture
	INNER JOIN tbl_fixture AS fixture ON fixture.fixtureId = tbl_user_fixture.fixtureId
	INNER JOIN tbl_league ON fixture.league = tbl_league.leagueId
	INNER JOIN tbl_team AS homeTeam ON fixture.homeTeam = homeTeam.teamId
	INNER JOIN tbl_team AS awayTeam ON fixture.awayTeam = awayTeam.teamId
	WHERE userId = ` + userID + `
	ORDER BY fixtureId ASC`)
	if err != nil {
		utils.HandleError("UserFixture", "GetUserFixtures", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	for rows.Next() {
		var (
			league                                                                                     m.League
			fixture                                                                                    m.Fixture
			homeTeam, awayTeam                                                                         m.Team
			userFixtureID, userID, fixtureID, userFixtureStatus, homeTeamGoals, awayTeamGoals, elapsed int
			date, round, fixtureStatus                                                                 string
		)

		err = rows.Scan(&userFixtureID, &userID, &fixtureID, &userFixtureStatus, &date, &league.LeagueID, &round, &homeTeam.TeamID, &homeTeam.Name, &homeTeam.LogoURL, &homeTeamGoals, &awayTeam.TeamID, &awayTeam.Name, &awayTeam.LogoURL, &awayTeamGoals, &fixtureStatus, &elapsed, &league.Name, &league.Country, &league.LogoURL, &league.FlagURL)
		if err != nil {
			utils.HandleError("UserFixture", "GetUserFixtures", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
			return
		}

		fixture = m.Fixture{
			FixtureID:         fixtureID,
			Date:              date,
			League:            league,
			Round:             round,
			HomeTeam:          homeTeam,
			HomeTeamGoals:     homeTeamGoals,
			AwayTeam:          awayTeam,
			AwayTeamGoals:     awayTeamGoals,
			Status:            fixtureStatus,
			Elapsed:           elapsed,
			UserFixtureStatus: userFixtureStatus,
		}

		userFixture = m.UserFixture{
			UserFixtureID: userFixtureID,
			UserID:        userID,
			Fixture:       fixture,
			Status:        userFixtureStatus,
		}

		userFixtures = append(userFixtures, userFixture)
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    userFixtures,
		Rows:    len(userFixtures),
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*CreateUserFixture is used to create or update a tbl_user_fixture row in the database.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func CreateUserFixture(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		userFixture  m.UserFixtureRequest
		responseBody m.HTTPResponse
	)

	// get user ID from session
	userID, err := getUserIDFromSession(r)
	if err != nil {
		utils.HandleError("UserFixture", "CreateUserFixture", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	stmtIns, err := db.Prepare("INSERT INTO tbl_user_fixture (userId, fixtureId, status) VALUES (?, ?, ?)")
	if err != nil {
		utils.HandleError("UserFixture", "CreateUserFixture", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}
	defer stmtIns.Close()

	stmtUpd, err := db.Prepare(`UPDATE tbl_user_fixture
	SET status = ?
	WHERE userId = ? AND fixtureId = ?`)
	if err != nil {
		utils.HandleError("UserFixture", "CreateUserFixture", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}
	defer stmtUpd.Close()

	// fetch request body and decode into new User struct
	if err := json.NewDecoder(r.Body).Decode(&userFixture); err != nil {
		utils.HandleError("UserFixture", "CreateUserFixture", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// if UserFixtureID is valid, then update row. Create new one otherwise
	if userFixture.UserFixtureID != 0 {
		_, err := stmtUpd.Exec(userFixture.Status, userID, userFixture.Fixture)
		if err != nil {
			utils.HandleError("UserFixture", "CreateUserFixture", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
			return
		}
	} else {
		res, err := stmtIns.Exec(userID, userFixture.Fixture, userFixture.Status)
		if err != nil {
			utils.HandleError("UserFixture", "CreateUserFixture", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
			return
		}

		id, err := res.LastInsertId()
        if err != nil {
			utils.HandleError("UserFixture", "CreateUserFixture", err)
			SetResponse(w, http.StatusInternalServerError, responseBody)
			return
		}
		
		userFixture.UserFixtureID = int(id)
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    userFixture,
		Rows:    1,
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*DeleteUserFixture is used to delete a tbl_user_fixture row in the database.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: DELETE
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func DeleteUserFixture(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		responseBody  m.HTTPResponse
		fixtureUserID int
		userFixtureID string
	)

	// get user fixture ID from URL
	userFixtureID = mux.Vars(r)["id"]

	// get user ID from session and compare to the one in ID - user can only delete his own user fixture rows
	userID, err := getUserIDFromSession(r)
	if err != nil {
		utils.HandleError("UserFixture", "DeleteUserFixture", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		utils.HandleError("UserFixture", "DeleteUserFixture", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	row := db.QueryRow("SELECT userId FROM tbl_user_fixture WHERE userFixtureID = " + userFixtureID)
	row.Scan(&fixtureUserID)

	if fixtureUserID != intUserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	stmtDel, err := db.Prepare(`DELETE FROM tbl_user_fixture
	WHERE userFixtureId = ?`)
	if err != nil {
		utils.HandleError("UserFixture", "DeleteUserFixture", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	_, err = stmtDel.Exec(userFixtureID)
	if err != nil {
		utils.HandleError("UserFixture", "DeleteUserFixture", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    true,
		Rows:    1,
	}

	SetResponse(w, http.StatusOK, responseBody)
}
