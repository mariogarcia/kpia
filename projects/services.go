package projects

var projects []Project = []Project{
	Project{Name: "Taiga"},
}

var members []Member = []Member{
	Member{Name: "John"},
}

// ListProjects lists all Piweek projects of all times
func ListProjects() []Project {
	return projects
}

// ListProjectMembers list all members of a given project
func ListProjectMembers(projectID int16) []Member {
	return members
}
