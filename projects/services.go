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
	return utils.PagedResult{}
}

// FindAllMembersByProjectID list all members of a given project
func (service *Service) FindAllMembersByProjectID(projectID string) utils.PagedResult {
	return utils.PagedResult{}
}

// FindProjectByID is
func (service *Service) FindProjectByID(projectID string) Project {
	return Project{}
}

// FindProjectMemberByID is
func (service *Service) FindProjectMemberByID(projectID, memberID string) ProjectMember {
	return ProjectMember{}
}
