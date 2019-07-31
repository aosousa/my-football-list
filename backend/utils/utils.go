package utils

import (
	"crypto/rand"
	"fmt"
	"net/smtp"
	"time"

	"github.com/aosousa/my-football-list/logger"
)

/*HandleError logs an error that occurred in the application

Receives:
	* controller (string) - Name of the handler/controller calling a method that lead to an error
	* method (string) - Name of the method that lead to an error
	* err (error) - Error that occurred
*/
func HandleError(controller, method string, err error) {
	logger.Error(nil, logger.SetData("method", method), err)
}

/*CheckPasswordResetTokenValidity checks if a user's password reset token is still valid

Receives:
	* tokenValidity (string) - Token expiry date in Y-m-d H:i:s format

Returns:
	* bool - Whether or not a user's password reset token is still valid
*/
func CheckPasswordResetTokenValidity(tokenValidity string) bool {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	currentTimeParsed, _ := time.Parse("2006-01-02 15:04:05", currentTime)
	passwordResetTokenValidity, _ := time.Parse("2006-01-02 15:04:05", tokenValidity)
	anHourBeforePasswordTokenValidity := passwordResetTokenValidity.Add(-1 * time.Hour)

	return currentTimeParsed.Before(passwordResetTokenValidity) && currentTimeParsed.After(anHourBeforePasswordTokenValidity)
}

/*GenerateRandomToken generates a token for password reset

Returns:
	* string - Token used for password reset
*/
func GenerateRandomToken() string {
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

/*SendContactEmail sends an e-mail to the provided e-mail address from the contact page

Receives:
	* to (string) - Recipient e-mail address
	* contactType (string) - Type of contact (Bug, Suggestion, etc.)
	* subject (string) - E-mail subject
	* message (string) - E-mail message body

Returns:
	* error - Error if an error occurred while sending the message, nil otherwise
*/
func SendContactEmail(to, contactType, subject, message string) error {
	from := "footballtracker01@gmail.com"
	pass := "rryzzunnjwsplwmn"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + contactType + " - " + subject + "\n\n" +
		message

	err := smtp.SendMail("smtp.gmail.com:587",
	smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
	from, []string{to}, []byte(msg))
	return err
}

/*SendResetPasswordEmail sends an e-mail to the provided e-mail address with further instructions in order to reset a user's password

Receives:
	* to (string) - User's e-mail address
	* token (string) - Password reset token

Returns:
	* error - Error if an error occurred while sending the message, nil otherwise
*/
func SendResetPasswordEmail(to, token string) error {
	from := "footballtracker01@gmail.com"
	pass := "rryzzunnjwsplwmn"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
	msg := "Subject: Football Tracker - Password Reset\n" +
		mime +
		"Hello!<br><br>" +
		"You have requested to reset the password of your Football Tracker account.<br><br>" +
		"If someone other than you was the one requesting this change, you can ignore this e-mail. Otherwise, click the link below to complete the process. The link below will work for one hour, after that you must request a password reset again.<br><br>" +
		"<a href=\"http://localhost:4200/password/" + token + "\">Reset Password</a>"

	// TODO: change to final link

	err := smtp.SendMail("smtp.gmail.com:587",
	smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
	from, []string{to}, []byte(msg))
	return err
}