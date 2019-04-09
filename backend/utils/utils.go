package utils

import (
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
