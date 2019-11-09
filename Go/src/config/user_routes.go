package config

import (
	"net/http"
	"server/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (router *Router) usersIndex(w http.ResponseWriter, r *http.Request) {
	users := models.GetAll(router.Server.Env.Connection)

	router.HandleJSONResponse(w, users, http.StatusOK)
}

func (router *Router) usersShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		return
	}

	user := models.GetByKey(router.Server.Env.Connection, id)
	var statusCode int

	if user != nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusNotFound
	}

	router.HandleJSONResponse(w, user, statusCode)
}

func (router *Router) usersCreate(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	user, err := models.Create(router.Server.Env.Connection, username)

	if err != nil {
		router.Abort(w, http.StatusConflict)
		return
	}

	router.HandleJSONResponse(w, user, http.StatusCreated)
}
