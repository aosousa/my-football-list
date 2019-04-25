package utils

import (
	"crypto/rand"
	"fmt"
	"net/smtp"
	"time"

	"github.com/aosousa/my-football-list/logger"
)

/*HandleError logs an error that occurred in the application
 * Receives:
 * controller (string) - Name of the handler/controller calling a method that lead to an error
 * method (string) - Name of the method that lead to an error
 * err (error) - Error that occurred
 */
func HandleError(controller, method string, err error) {
	logger.Error(nil, logger.SetData("method", method), err)
}

/*GetCurrentDateTime returns the current date/time
 * Returns: string - Current date/time in YYYY-mm-dd HH:ii:ss format
 */
func GetCurrentDateTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

// IsEmpty checks if a required field is empty
func IsEmpty(field string) bool {
	return field == ""
}

// CheckPasswordResetTokenValidity checks if a user's password reset token is still valid
func CheckPasswordResetTokenValidity(tokenValidity string) bool {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	currentTimeParsed, _ := time.Parse("2006-01-02 15:04:05", currentTime)
	passwordResetTokenValidity, _ := time.Parse("2006-01-02 15:04:05", tokenValidity)
	anHourBeforePasswordTokenValidity := passwordResetTokenValidity.Add(-1 * time.Hour)

	return currentTimeParsed.Before(passwordResetTokenValidity) && currentTimeParsed.After(anHourBeforePasswordTokenValidity)
}

// GenerateRandomToken generates a token for password reset
func GenerateRandomToken() string {
	b := make([]byte, 12)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// SendEmail sends an e-mail to the provided e-mail address
func SendEmail(to, contactType, subject, message string) error {
	from := ""
	pass := ""
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + contactType + " - " + subject + "\n\n" +
		message

	err := smtp.SendMail("smtp.gmail.com:587",
	smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
	from, []string{to}, []byte(msg))
	return err
}