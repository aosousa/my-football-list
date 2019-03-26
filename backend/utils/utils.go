package utils

import (
	"fmt"
	"time"
)

/*HandleError logs an error that occurred in the application
 * Receives:
 * controller (string) - Name of the handler/controller calling a method that lead to an error
 * method (string) - Name of the method that lead to an error
 * err (error) - Error that occurred
 */
func HandleError(controller, method string, err error) {
	currentTime := time.Now()
	currentTimeFmt := currentTime.Format("2006-01-02 15:04:05")

	fmt.Printf("%s ERROR in %s (%s)\n", currentTimeFmt, method, controller)
	fmt.Println(err)
}
