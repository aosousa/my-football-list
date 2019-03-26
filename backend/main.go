package main

import (
	"fmt"
	"net/http"
	"os"

	h "github.com/aosousa/my-football-list/handlers"
	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

func init() {
	// set up Config struct before performing any queries
	fmt.Println("Configuration file: Loading")
	h.InitConfig()
	fmt.Printf("Configuration file: OK\n\n")

	// initialize database through Config struct
	fmt.Println("System database: Checking")
	h.InitDatabase()
	fmt.Printf("System database: OK\n\n")
}

func main() {
	// start cron job that will update the database periodically
	// according to the refresh timer set in the configuration file
	go func() {
		h.StartCronJob()
		// h.SaveTeams("96")
	}()

	// start HTTP server that will handle all API requests
	router := NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Cache-Control", "Pragma", "Origin", "Authorization", "Content-Type", "X-Requested-With", "Expiry"},
	})

	routerHandler := c.Handler(router)

	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, routerHandler))
}
