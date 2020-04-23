package api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lib/pq" // PosgreSQL driver
	"log"
	"net/http"
)

// API represents the whole REST API application
type API struct {
	router *mux.Router
	DB     *sql.DB
}

// InitDB initializes database
func (api *API) InitDB(url string) {
	log.Println(fmt.Sprintf("initializing database: %s", url))
	db, err := sql.Open("postgres", url)
	api.DB = db

	if err == nil {
		LoadFixtures(api.DB)
	} else {
		error := err.(*pq.Error)
		log.Fatal(error.Code.Name())
	}
}

// InitHandlers mappings between paths and handlers
func (api *API) InitHandlers() {
	log.Println("initializing handlers")
	api.router = mux.NewRouter()
	api.router.HandleFunc("/projects", api.listProjects).Methods(http.MethodGet)
	api.router.HandleFunc("/projects/{project}", api.getProject).Methods(http.MethodGet)
	api.router.HandleFunc("/projects/{project}/members", api.listProjectMembers).Methods(http.MethodGet)
	api.router.HandleFunc("/projects/{project}/members/{member}", api.getProjectMember).Methods(http.MethodGet)
}

// Startup initializes API
func (api *API) Startup(addr string) {
	log.Println("starting app...")
	log.Fatal(http.ListenAndServe(addr, api.router))
}
