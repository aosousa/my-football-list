package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	m "github.com/aosousa/my-football-list/models"
	"github.com/aosousa/my-football-list/utils"
)

/*SendEmail sends an email address to myself through the contact form
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func SendEmail(w http.ResponseWriter, r *http.Request) {
	// check user's authentication status before proceeding
	auth, ok := checkAuthStatus(w, r)
	if !auth || !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var (
		contact      m.Contact
		responseBody m.HTTPResponse
	)

	// get request body
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		utils.HandleError("Contact", "SendEmail", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// check if required fields are empty
	if utils.IsEmpty(contact.Type) || utils.IsEmpty(contact.Subject) || utils.IsEmpty(contact.Message) {
		err := errors.New("Required fields are empty")
		utils.HandleError("Contact", "SendEmail", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	to := ""
	err := utils.SendEmail(to, contact.Type, contact.Subject, contact.Message)
	if err != nil {
		utils.HandleError("Contact", "SendEmail", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	responseBody = m.HTTPResponse{
		Success: true,
	}

	SetResponse(w, http.StatusOK, responseBody)
	return
}

/*SendResetPasswordEmail sends an email address for 
 *
 * Receives: http.ResponseWriter and http.Request
 * Request method: POST
 *
 * Response
 * Content-Type: application/json
 * Body: m.HTTPResponse
 */
func SendResetPasswordEmail(w http.ResponseWriter, r *http.Request) {
	var (
		resetPasswordStruct m.ResetPassword
		responseBody m.HTTPResponse
		validTokenExists bool
	)

	if err := json.NewDecoder(r.Body).Decode(&resetPasswordStruct); err != nil {
		utils.HandleError("Contact", "SendResetPasswordEmail", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// find user with email received
	user, err := GetUserByEmail(resetPasswordStruct.Email)
	if err != nil {
		utils.HandleError("Contact", "SendResetPasswordEmail", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// check if a valid token already exists
	validTokenExists = utils.CheckPasswordResetTokenValidity(user.PasswordResetTokenValidity.String)

	if validTokenExists {
		err = errors.New("A reset password request for that e-mail has already been made recently. Check the e-mail address provided for further instructions on how to reset your password.")

		responseBody.Error = err.Error()
		utils.HandleError("Contact", "SendResetPasswordEmail", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// create UUID to use as reset token
	userPasswordResetToken := utils.GenerateRandomToken();

	// add validity of 1 hour to token
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	currentTimeParsed, _ := time.Parse("2006-01-02 15:04:05", currentTime)
	userPasswordResetTokenValidity := currentTimeParsed.Add(1*time.Hour)

	// update user in database - including updateTime field
	stmtUpd, err := db.Prepare(`UPDATE tbl_user
	SET passwordResetToken = ?, passwordResetTokenValidity = ?, updateTime = ?
	WHERE email = ?`)
	if err != nil {
		utils.HandleError("Contact", "SendResetPasswordEmail", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	_, err = stmtUpd.Exec(userPasswordResetToken, userPasswordResetTokenValidity, currentTime, resetPasswordStruct.Email)
	if err != nil {
		utils.HandleError("Contact", "SendResetPasswordEmail", err)
		SetResponse(w, http.StatusInternalServerError, responseBody)
		return
	}

	// TODO: send email to user

	// set response body
	responseBody = m.HTTPResponse{
		Success: true,
	}

	SetResponse(w, http.StatusOK, responseBody)
	return
}
