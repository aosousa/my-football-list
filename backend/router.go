package main

import (
	"net/http"

	"github.com/gorilla/mux"

	h "github.com/aosousa/my-football-list/handlers"
	m "github.com/aosousa/my-football-list/models"
)

/*NewRouter creates a new mux Router with the routes defined
 * in the method below
 */
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/test", test).Methods("GET")
	router.HandleFunc("/signup", h.Signup).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/logout", h.Logout).Methods("POST")

	return router
}

func test(w http.ResponseWriter, r *http.Request) {
	body := m.HTTPResponse{
		Success: true,
		Data:    "test",
		Rows:    0,
	}

	h.SetResponse(w, http.StatusOK, body)
}
