package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	m "./models"
)

/*NewRouter creates a new mux Router with the routes defined
 * in the method below
 */
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/test", test).Methods("GET")

	return router
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m.HTTPResponse{Success: true, Data: "test", Rows: 0})
	return
}
