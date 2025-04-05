package controllers

type ProjectDetailsDTO struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectDTO struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectCreateData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectCreateRequestDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

