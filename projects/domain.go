package projects

// Member represents a member of a Project
type Member struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// Project represents a Piweek project
type Project struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Members     []Member `json:"members"`
}

// ProjectMember represents a given Member working inside a Project
type ProjectMember struct {
	Project Project
	Member  Member
}
