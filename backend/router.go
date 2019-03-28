package main

import (
	"github.com/gorilla/mux"

	h "github.com/aosousa/my-football-list/handlers"
)

/*NewRouter creates a new mux Router with the routes defined
 * in the method below
 */
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Auth methods
	router.HandleFunc("/signup", h.Signup).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/logout", h.Logout).Methods("POST")

	// League methods
	router.HandleFunc("/leagues", h.GetAllLeagues).Methods("GET")
	router.HandleFunc("/leagues/{id}/fixtures", h.GetLeagueFixtures).Methods("GET")

	// Fixture methods
	router.HandleFunc("/teams/{id}/fixtures", h.GetTeamFixtures).Methods("GET")

	return router
}
