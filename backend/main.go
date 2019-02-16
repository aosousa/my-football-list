package main

import (
	"fmt"
	"os"
)

func init() {
	// set up Config struct before performing any queries
	fmt.Println("Configuration file: Loading")
	_, err := initConfig()
	if err != nil {
		fmt.Println("Failed to load configuration file. Check the logs for more information.")
		os.Exit(1)
	}
	fmt.Println("Configuration file: OK")

	// TODO: initialize database through Config struct
}

func main() {

}
