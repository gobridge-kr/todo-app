package controller

import (
	"net/http"
)

type Controller interface {
	GetOne(w http.ResponseWriter, r *http.Request, id string)
	GetAll(w http.ResponseWriter, r *http.Request)
	PostOne(w http.ResponseWriter, r *http.Request, id string)
	PostAll(w http.ResponseWriter, r *http.Request)
	PatchOne(w http.ResponseWriter, r *http.Request, id string)
	PatchAll(w http.ResponseWriter, r *http.Request)
	DeleteOne(w http.ResponseWriter, r *http.Request, id string)
	DeleteAll(w http.ResponseWriter, r *http.Request)
	Options(w http.ResponseWriter, r *http.Request)
}
