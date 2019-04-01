package handlers

import (
	"encoding/json"
	"net/http"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
	"github.com/gorilla/mux"
)

/*GetUserByUsername queries for a User row in the database.
 * Receives:
 * username (string) - User's username
 *
 * Returns:
 * User - User struct found
 * err - Description of error found during execution (or nil otherwise)
 */
func GetUserByUsername(username string) (m.User, error) {
	var u m.User

	err := db.QueryRow("SELECT userId, username, password, email FROM tbl_user WHERE username = ?", username).Scan(&u.UserID, &u.Username, &u.Password, &u.Email)
	if err != nil {
		utils.HandleError("User", "GetUserByUsername", err)
		return u, err
	}

	return u, nil
}

/*GetUser queries the database for a user's information.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func GetUser(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		user         m.User
		responseBody m.HTTPResponse
		userID       string
	)

	// get user ID from URL
	userID = mux.Vars(r)["id"]

	err := db.QueryRow("SELECT userId, username, email, createTime, spoilerMode FROM tbl_user WHERE userId = "+userID).Scan(&user.UserID, &user.Username, &user.Email, &user.CreateTime, &user.SpoilerMode)
	if err != nil {
		utils.HandleError("User", "GetUser", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    user,
		Rows:    1,
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*UpdateUser updates a user's information.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: PUT
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		user         m.User
		responseBody m.HTTPResponse
		userID       string
	)

	// get user ID from URL
	userID = mux.Vars(r)["id"]

	stmtUpd, err := db.Prepare(`UPDATE tbl_user
	SET email = ?, updateTime = ?, spoilerMode = ?
	WHERE userId = ?`)
	if err != nil {
		utils.HandleError("User", "UpdateUser", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
	}
	defer stmtUpd.Close()

	// fetch request body and decode into new User struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError("User", "UpdateUser", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
	}

	// get current date time for user's UpdateTime field
	user.UpdateTime = utils.GetCurrentDateTime()

	_, err = stmtUpd.Exec(user.Email, user.UpdateTime, user.SpoilerMode, userID)
	if err != nil {
		utils.HandleError("User", "UpdateUser", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Rows:    1,
	}

	SetResponse(w, http.StatusOK, responseBody)
}
