package figma

type teamProjectsResponse struct {
	Projects []TeamProject `json:"projects"`
}

// TeamProject is a project which belongs to a team.
type TeamProject struct {
	// The ID of the project
	ID string `json:"id"`

	// The Name of the project
	Name string `json:"name"`
}
