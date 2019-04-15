package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/smtp"

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

	from := ""
	pass := ""
	to := ""
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + contact.Type + " - " + contact.Subject + "\n\n" +
		contact.Message

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
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
