package handlers

import (
	"encoding/json"
	"net/http"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
	"github.com/gorilla/mux"
)

/*GetUserByUsername queries for a User row in the database through username.
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

/*GetUserByEmail queries for a User row in the database through email.
 * Receives:
 * email (string) - User's email address
 *
 * Returns:
 * User - User struct
 * error - Description of error found during execution (or nil otherwise)
 */
func GetUserByEmail(email string) (m.User, error) {
	var u m.User

	err := db.QueryRow("SELECT userId, username, passwordResetToken, passwordResetTokenValidity, email FROM tbl_user WHERE email = ?", email).Scan(&u.UserID, &u.Username, &u.PasswordResetToken, &u.PasswordResetTokenValidity, &u.Email)
	if err != nil {
		utils.HandleError("User", "GetUserByEmail", err)
		return u, err
	}
	
	return u, nil
}

/*GetUserByToken queries for a User row in the database through password reset token.
 * Receives:
 * token (string) - User's password reset token
 *
 * Returns:
 * User - User struct
 * error - Description of error found during execution (or nil otherwise)
 */
func GetUserByToken(token string) (m.User, error) {
	var u m.User

	err := db.QueryRow("SELECT passwordResetToken, passwordResetTokenValidity FROM tbl_user WHERE passwordResetToken = ?", token).Scan(&u.PasswordResetToken, &u.PasswordResetTokenValidity)
	if err != nil {
		utils.HandleError("User", "GetUserByToken", err)
		return u, err
	}

	return u, nil
}

/*Queries for a User row in the database through username.
 * Receives:
 * username (string) - User's username
 *
 * Returns:
 * int8 - Number of Users with that email (0 or 1 max)
 * error - Description of error found during execution (or nil otherwise)
 */
func getUserByUsername(username string) (int8, error) {
	var userCount int8

	err := db.QueryRow("SELECT count(*) AS count FROM tbl_user WHERE username = ?", username).Scan(&userCount)
	if err != nil {
		utils.HandleError("User", "getUserByUsername", err)
		return userCount, err
	}

	return userCount, nil
}

/*Queries for a User row in the database through email.
 * Receives:
 * email (string) - User's email
 *
 * Returns:
 * int8 - Number of Users with that email (0 or 1 max)
 * error - Description of error found during execution (or nil otherwise)
 */
func getUserByEmail(email string) (int8, error) {
	var userCount int8

	err := db.QueryRow("SELECT count(*) AS count FROM tbl_user WHERE email = ?", email).Scan(&userCount)
	if err != nil {
		utils.HandleError("User", "getUserByEmail", err)
		return userCount, err
	}

	return userCount, nil
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
		return
	}

	// get current date time for user's UpdateTime field
	user.UpdateTime = utils.GetCurrentDateTime()

	_, err = stmtUpd.Exec(user.Email, user.UpdateTime, user.SpoilerMode, userID)
	if err != nil {
		utils.HandleError("User", "UpdateUser", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Rows:    1,
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*CheckUsernameExistence checks if a User with a certain username already exists in the database.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func CheckUsernameExistence(w http.ResponseWriter, r *http.Request) {
	var (
		userCount    int8
		user         m.User
		responseBody m.HTTPResponse
	)

	// fetch request body and decode into new User struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError("User", "CheckUsernameExistence", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	userCount, err := getUserByUsername(user.Username)
	if err != nil {
		utils.HandleError("User", "CheckUsernameExistence", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Rows:    int(userCount),
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*CheckEmailExistence checks if a User with a certain email already exists in the database.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func CheckEmailExistence(w http.ResponseWriter, r *http.Request) {
	var (
		userCount    int8
		user         m.User
		responseBody m.HTTPResponse
	)

	// fetch request body and decode into new User struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError("User", "CheckEmailExistence", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	userCount, err := getUserByEmail(user.Email)
	if err != nil {
		utils.HandleError("User", "CheckEmailExistence", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Rows:    int(userCount),
	}

	SetResponse(w, http.StatusOK, responseBody)
}
