package api

import (
	"github.com/gorilla/mux"
	"github.com/mariogarcia/kpia/projects"
	"github.com/mariogarcia/kpia/utils"
	"net/http"
)

// Lists all Piweek's projects of all times
func (api *API) listProjects(response http.ResponseWriter, request *http.Request) {
	pagination := utils.GetPagination(request.URL.Query())
	service := projects.Service{DB: api.DB}

	renderJSON(response, service.ListProjects(pagination))
}

// Gets a project by its id
func (api *API) getProject(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	ID := params["project"]
	service := projects.Service{DB: api.DB}

	renderJSON(response, service.FindProjectByID(ID))
}

func (api *API) listProjectMembers(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	ID := params["project"]
	service := projects.Service{DB: api.DB}

	renderJSON(response, service.FindAllMembersByProjectID(ID))
}

func (api *API) getProjectMember(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	projectID := params["project"]
	memberID := params["member"]
	service := projects.Service{DB: api.DB}

	renderJSON(response, service.FindProjectMemberByID(projectID, memberID))
}
