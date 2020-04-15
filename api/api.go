package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func (a *api) Router() http.Handler {
	return a.router
}

func New() Server {
	a := &api{}
	r := mux.NewRouter()

	r.HandleFunc("/projects", a.listProjects).Methods(http.MethodGet)
	r.HandleFunc("/projects/:id", a.getProject).Methods(http.MethodGet)

	a.router = r
	return a
}
