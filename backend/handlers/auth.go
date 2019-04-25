package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
)

var (
	key   = []byte("a1XvFXA0A2booqYcsyEDl2AtmlpAiLu8")
	store = sessions.NewCookieStore(key)
)

/* Checks the the validity of a password hash.
 *
 * Receives:
 * password (string) - Password hash saved in the database
 * hash (string) - Password hash generated from the password in plain-text received on login
 *
 * Returns:
 * bool - Whether or not the password is valid
 */
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/* Hashes a user's password using bcrypt.
 *
 * Receives:
 * password (string) - Password in plain text
 *
 * Returns:
 * string - Password hash
 * error - Error in case one occurred during execution
 */
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*Signup is the function that handles a POST /signup HTTP request.
 * Creates the user in the platform and automatically logs the user in afterwards.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func Signup(w http.ResponseWriter, r *http.Request) {
	var (
		user           m.User
		hashedPassword string
		responseBody   m.HTTPResponse
	)

	// get request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError("Auth", "Signup", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// check if required fields are empty
	if utils.IsEmpty(user.Username) || utils.IsEmpty(user.Password) {
		err := errors.New("Required field is empty")
		utils.HandleError("Auth", "Signup", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// hash user's password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		utils.HandleError("Auth", "Signup", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	user.Password = hashedPassword

	stmtIns, err := db.Prepare("INSERT INTO tbl_user (username, password, email, spoilerMode) VALUES (?, ?, ?, ?)")
	if err != nil {
		utils.HandleError("Auth", "Signup", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	result, err := stmtIns.Exec(user.Username, user.Password, user.Email, user.SpoilerMode)
	if err != nil {
		utils.HandleError("Auth", "Signup", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// get user id
	userID, err := result.LastInsertId()
	if err != nil {
		utils.HandleError("Auth", "Signup", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// create user session
	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "Login", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	session.Values["authenticated"] = true
	session.Values["userID"] = userID
	session.Values["username"] = user.Username
	session.Save(r, w)

	returnUser := m.User{
		UserID:   int(userID),
		Username: user.Username,
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    returnUser,
		Rows:    1,
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*Login is the function that handles a POST /login HTTP request.
 * Logs the user in the platform if the credentials are correct.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func Login(w http.ResponseWriter, r *http.Request) {
	var (
		user, returnUser m.User
		loginSuccess     bool
		statusCode       int
		responseBody     m.HTTPResponse
	)

	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "Login", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError("Auth", "Login", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	compareUser, err := GetUserByUsername(user.Username)
	if err != nil {
		utils.HandleError("Auth", "Login", err)
		SetResponse(w, http.StatusOK, responseBody)
		return
	}

	passwordHashMatch := checkPasswordHash(user.Password, compareUser.Password)
	if passwordHashMatch {
		loginSuccess, statusCode = true, 200
		session.Values["authenticated"] = true
		session.Values["userID"] = compareUser.UserID
		session.Values["username"] = compareUser.Username
		session.Save(r, w)

		returnUser = m.User{
			UserID:   compareUser.UserID,
			Username: compareUser.Username,
		}
	} else {
		loginSuccess, statusCode = false, 401
	}

	responseBody = m.HTTPResponse{
		Success: loginSuccess,
		Data:    returnUser,
	}

	SetResponse(w, statusCode, responseBody)
}

/*Logout is the function that handles a POST /logout HTTP request.
 * Logs the user out of the platform by destroying the user session.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func Logout(w http.ResponseWriter, r *http.Request) {
	var responseBody m.HTTPResponse

	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "Logout", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// revoke user's authentication
	session.Values["authenticated"] = false
	session.Values["userID"] = 0
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		utils.HandleError("Auth", "Logout", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    true,
		Rows:    0,
	}

	SetResponse(w, http.StatusOK, responseBody)
}

func LoggedInUser(w http.ResponseWriter, r *http.Request) {
	// check user's authentication before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		responseBody m.HTTPResponse
		user         m.User
	)

	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "LoggedInUser", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	user = m.User{
		UserID:   session.Values["userID"].(int),
		Username: session.Values["username"].(string),
	}

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    user,
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*IsResetTokenValid checks if the reset password token sent is still valid for use.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: GET
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func IsResetTokenValid(w http.ResponseWriter, r *http.Request) {
	var (
		user               m.User
		responseBody       m.HTTPResponse
		passwordResetToken string
		isValid            bool
	)

	// get token from URL
	passwordResetToken = mux.Vars(r)["token"]

	// find user with that token
	err := db.QueryRow("SELECT userId, username, passwordResetToken, passwordResetTokenValidity, email FROM tbl_user WHERE passwordResetToken = '"+passwordResetToken+"'").Scan(&user.UserID, &user.Username, &user.PasswordResetToken, &user.PasswordResetTokenValidity, &user.Email)
	if err != nil {
		utils.HandleError("Auth", "IsResetTokenValid", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// check if a valid token already exists
	isValid = utils.CheckPasswordResetTokenValidity(user.PasswordResetTokenValidity.String)

	responseBody = m.HTTPResponse{
		Success: true,
		Data:    isValid,
	}

	SetResponse(w, http.StatusOK, responseBody)
}

/*ResetPassword is used to reset a user's password.
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var (
		resetPasswordStruct m.ResetPassword
		responseBody m.HTTPResponse
		tokenIsValid bool
	)

	if err := json.NewDecoder(r.Body).Decode(&resetPasswordStruct); err != nil {
		utils.HandleError("Auth", "ResetPassword", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// find user with token received
	user, err := GetUserByToken(resetPasswordStruct.Token)
	if err != nil {
		utils.HandleError("Auth", "ResetPassword", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// check if token received is still valid
	tokenIsValid = utils.CheckPasswordResetTokenValidity(user.PasswordResetTokenValidity.String)

	if !tokenIsValid {
		err = errors.New("Reset password token is no longer valid. Create a new reset password request to proceed.")

		responseBody.Error = err.Error()
		utils.HandleError("Auth", "ResetPassword", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// hash user's password
	hashedPassword, err := hashPassword(resetPasswordStruct.Password)
	if err != nil {
		utils.HandleError("Auth", "ResetPassword", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// save new password, and save passwordResetToken and passwordResetTokenValidity as nulls again
	currentTime := utils.GetCurrentDateTime()

	stmtUpd, err := db.Prepare(`UPDATE tbl_user
	SET password = ?, passwordResetToken = ?, passwordResetTokenValidity = ?, updateTime = ? WHERE passwordResetToken = ?`)
	if err != nil {
		utils.HandleError("Auth", "ResetPassword", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	_, err = stmtUpd.Exec(hashedPassword, sql.NullString{}, sql.NullString{}, currentTime, resetPasswordStruct.Token)
	if err != nil {
		utils.HandleError("Auth", "ResetPassword", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// TODO: send email confirming password change

	// set response body
	responseBody = m.HTTPResponse{
		Success: true,
	}

	SetResponse(w, http.StatusOK, responseBody)
	return
}

/*Gets the ID of the current logged in user.
 *
 * Receives: http.ResponseWriter and http.Request
 *
 * Returns:
 * int - ID of the current logged in user
 * error - Error in case one occurred (nil otherwise)
 */
func getUserIDFromSession(r *http.Request) (string, error) {
	var userID string

	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "GetUserIDFromSession", err)
		return userID, err
	}

	userID = strconv.Itoa(session.Values["userID"].(int))

	return userID, nil
}

/*CheckAuthStatus is the function that checks whether or not the user is logged in the platform.
 * It is called at the start of most handler functions (auth methods being an exception)
 *
 * Receives: http.ResponseWriter and http.Request
 *
 * Returns: bool, bool - Status of the user's authentication
 */
func checkAuthStatus(w http.ResponseWriter, r *http.Request) (bool, bool) {
	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "CheckAuthStatus", err)
		return false, false
	}

	// check if the user is authenticated
	auth, ok := session.Values["authenticated"].(bool)

	return auth, ok
}
