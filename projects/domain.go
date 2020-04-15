package projects

// Member represents a member of a Project
type Member struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

// Project represents a Piweek project
type Project struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Members     []Member `json:"members"`
}
