package entity

type Response struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Error   []ErrorDetails `json:"error"`
}

type ErrorDetails struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
	Type    string   `json:"type"`
	Context Context  `json:"context"`
}

type Context struct {
	Label string `json:"label"`
	Key   string `json:"key"`
}

type Request struct {
	Name string `json:"name" validate:"required"`
}
