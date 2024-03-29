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
	router.HandleFunc("/tokens/{token}/valid", h.IsResetTokenValid).Methods("GET")
	router.HandleFunc("/reset-password", h.ResetPassword).Methods("POST")

	// League methods
	router.HandleFunc("/leagues", h.GetAllLeagues).Methods("GET")
	router.HandleFunc("/leagues/{id}/fixtures", h.GetLeagueFixtures).Methods("GET")

	// Team methods
	router.HandleFunc("/teams/{id}/fixtures", h.GetTeamFixtures).Methods("GET")

	// Fixture methods
	router.HandleFunc("/fixtures/by-date/{date}", h.GetDateFixtures).Methods("GET")
	router.HandleFunc("/fixtures/last-update", h.GetLastFixtureUpdate).Methods("GET")

	// User and User/Fixture methods
	router.HandleFunc("/users/current", h.LoggedInUser).Methods("GET")
	router.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}/fixtures", h.GetUserFixtures).Methods("GET")
	router.HandleFunc("/users/username-existence", h.CheckUsernameExistence).Methods("POST")
	router.HandleFunc("/users/email-existence", h.CheckEmailExistence).Methods("POST")
	router.HandleFunc("/users/fixtures", h.CreateUserFixture).Methods("POST")
	router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}/change-password", h.ChangePassword).Methods("PUT")
	router.HandleFunc("/user-fixtures/{id}", h.DeleteUserFixture).Methods("DELETE")

	// Contact methods
	router.HandleFunc("/contact", h.SendEmail).Methods("POST")
	router.HandleFunc("/reset-password-email", h.SendResetPasswordEmail).Methods("POST")

	return router
}
