package utils

import (
	"fmt"
	"os"
)

/*HandleError logs an error that occurred in the application
 * Receives:
 * controller (string) - Name of the handler/controller calling a method that lead to an error
 * method (string) - Name of the method that lead to an error
 * err (error) - Error that occurred
 */
func HandleError(controller, method string, err error) {
	fmt.Printf("ERROR | [%s - %s] %s\n", controller, method, err)
	os.Exit(1)
}
