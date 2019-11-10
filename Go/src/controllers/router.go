package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"server/config"
	"strconv"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
	Server *config.Server
}

// InitializeRouter initializes routes for the api
func InitializeRouter(s *config.Server) {
	r := &Router{
		Router: mux.NewRouter(),
		Server: s,
	}

	initializeUserRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r.Router))
}

func initializeUserRoutes(r *Router) {
	r.Router.HandleFunc("/{id}", r.usersShow).Methods("GET")
	r.Router.HandleFunc("/", r.usersIndex).Methods("GET")
	r.Router.HandleFunc("/", r.usersCreate).Methods("POST")
	r.Router.HandleFunc("/{id}", r.usersDestroy).Methods("DELETE")
}

// JSONResponse handels API JSON responses
func (r *Router) JSONResponse(w http.ResponseWriter, d interface{}, statusCode int) {
	js, err := json.Marshal(d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// JSONResponse handels API error responses
func (r *Router) Error(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}

func (router *Router) getIDFromRequest(r *http.Request) (*int64, error) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		return nil, err
	}

	return &id, nil
}
