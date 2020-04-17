package projects

import (
	"database/sql"
	"github.com/mariogarcia/kpia/utils"
)

// Service represents busines logic for projects
type Service struct {
	DB *sql.DB
}

// ListProjects lists all Piweek projects of all times
func (service *Service) ListProjects(pagination utils.Pagination) utils.PagedResult {
	results := []Project{}
	rows, _ := service.DB.Query(`SELECT * FROM projects`)

	// closing connection when logic has finished
	defer rows.Close()

	// looping results
	for rows.Next() {
		results = append(results, mapToProject(rows))
	}

	return utils.PagedResult{}
}

func mapToProject(rows *sql.Rows) Project {
	pro := Project{}
	rows.Scan(&pro.ID, &pro.Name, &pro.Description)
	return pro
}

// FindAllMembersByProjectID list all members of a given project
func (service *Service) FindAllMembersByProjectID(projectID string) utils.PagedResult {
	rows, err := service.DB.Query(`SELECT * FROM projects_members WHERE project_id = $1`, projectID)

	return utils.PagedResult{}
}

// FindProjectByID is
func (service *Service) FindProjectByID(projectID string) Project {
	row, err := service.DB.QueryRow(`SELECT * FROM project WHERE id = $1`, projectID)

	return Project{}
}

// FindProjectMemberByID is
func (service *Service) FindProjectMemberByID(projectID, memberID string) ProjectMember {
	row, err := service.DB.QueryRow(`SELECT * FROM project_member WHERE project = $1 AND member = $2`, projectID, memberID)

	return ProjectMember{}
}
