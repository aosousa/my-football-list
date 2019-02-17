package main

import (
	"fmt"
	"net/http"
	"os"

	h "./handlers"
)

func init() {
	// set up Config struct before performing any queries
	fmt.Println("Configuration file: Loading")
	_, err := h.InitConfig()
	if err != nil {
		fmt.Println("Failed to load configuration file. Check the logs for more information.")
		os.Exit(1)
	}
	fmt.Println("Configuration file: OK")

	// TODO: initialize database through Config struct
}

func main() {
	// start cron job that will update the database periodically
	// according to the refresh timer set in the configuration file
	go func() {
		h.StartCronJob()
	}()

	// start HTTP server that will handle all API requests
	router := NewRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", router)
}
