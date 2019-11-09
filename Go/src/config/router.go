package config

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
	Server *Server
}

func InitializeRouter(s *Server) *Router {
	r := Router{Router: mux.NewRouter(), Server: s}

	r.Router.HandleFunc("/{id}", r.usersShow).Methods("GET")
	r.Router.HandleFunc("/", r.usersIndex).Methods("GET")
	r.Router.HandleFunc("/", r.usersCreate).Methods("POST")

	return &r
}

func (r *Router) HandleJSONResponse(w http.ResponseWriter, d interface{}, statusCode int) {
	js, err := json.Marshal(d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (r *Router) Abort(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
}
