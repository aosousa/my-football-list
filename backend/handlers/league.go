package handlers

import (
	"net/http"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
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
	auth, ok := CheckAuthStatus(w, r)
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
