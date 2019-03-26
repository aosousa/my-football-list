package handlers

import (
	"encoding/json"
	"net/http"

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
	)

	// get request body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError("Auth", "Signup", err)
	}

	// hash user's password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		utils.HandleError("Auth", "Signup", err)
	}

	user.Password = hashedPassword

	stmtIns, err := db.Prepare("INSERT INTO tbl_user (username, password, email) VALUES (?, ?, ?)")
	if err != nil {
		utils.HandleError("Auth", "Signup", err)
	}

	_, err = stmtIns.Exec(user.Username, user.Password, user.Email)
	if err != nil {
		utils.HandleError("Auth", "Signup", err)
	}

	responseBody := m.HTTPResponse{
		Success: true,
		Data:    true,
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
		user         m.User
		loginSuccess bool
		statusCode   int
	)

	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "Login", err)
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.HandleError("Auth", "Login", err)
	}

	compareUser := GetUserByUsername(user.Username)

	passwordHashMatch := checkPasswordHash(user.Password, compareUser.Password)
	if passwordHashMatch {
		loginSuccess, statusCode = true, 200
		session.Values["authenticated"] = true
		session.Save(r, w)
	} else {
		loginSuccess, statusCode = false, 401
	}

	responseBody := m.HTTPResponse{
		Success: loginSuccess,
		Data:    loginSuccess,
		Rows:    1,
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
	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "Logout", err)
	}

	// revoke user's authentication
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1

	err = session.Save(r, w)
	if err != nil {
		utils.HandleError("Auth", "Logout", err)
	}

	body := m.HTTPResponse{
		Success: true,
		Data:    true,
		Rows:    0,
	}

	SetResponse(w, http.StatusOK, body)
}

/*CheckAuthStatus is the function that checks whether or not the user is logged in the platform.
 * It is called at the start of most handler functions (auth methods being an exception)
 *
 * Receives: http.ResponseWriter and http.Request
 *
 * Returns: bool - Status of the user's authentication
 */
func CheckAuthStatus(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "session-token")
	if err != nil {
		utils.HandleError("Auth", "CheckAuthStatus", err)
	}

	// check if the user is authenticated
	return session.Values["authenticated"].(bool)
}
