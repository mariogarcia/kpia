package api

import (
	"github.com/mariogarcia/kpia/projects"
	"net/http"
)

// Lists all Piweek's projects of all times
func (a *api) listProjects(response http.ResponseWriter, request *http.Request) {
	projects := projects.ListProjects()

	renderJSON(response, projects)
}

// Gets a project by its id
func (a *api) getProject(response http.ResponseWriter, request *http.Request) {
	project := projects.ListProjects()[0]

	renderJSON(response, project)
}
