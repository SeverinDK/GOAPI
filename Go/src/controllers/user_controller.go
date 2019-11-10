package controllers

import (
	"net/http"
	"server/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (router *Router) usersIndex(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAll(router.Server.Env.Connection)

	if err != nil {
		router.Error(w, err, http.StatusConflict)
	}

	router.JSONResponse(w, users, http.StatusOK)
}

func (router *Router) usersShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		router.Error(w, err, http.StatusInternalServerError)
		return
	}

	user, err := models.GetByKey(router.Server.Env.Connection, id)

	if err != nil {
		router.Error(w, err, http.StatusInternalServerError)
		return
	}

	var statusCode int

	if user != nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusNotFound
	}

	router.JSONResponse(w, user, statusCode)
}

func (router *Router) usersCreate(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	user, err := models.Create(router.Server.Env.Connection, username)

	if err != nil {
		router.Error(w, err, http.StatusConflict)
		return
	}

	router.JSONResponse(w, user, http.StatusCreated)
}

func (router *Router) usersDestroy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	success, err := models.Delete(router.Server.Env.Connection, id)

	if err != nil {
		router.Error(w, err, http.StatusInternalServerError)
		return
	}

	var statusCode int

	if success == true {
		statusCode = http.StatusNoContent
	} else {
		statusCode = http.StatusNotFound
	}

	router.JSONResponse(w, success, statusCode)
}
