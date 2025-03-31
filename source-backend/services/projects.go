package services

import (
	"context"
	"fmt"
)

var NotFound = fmt.Errorf("Missinc record")

type ProjectDetails struct {
	Id          string
	Name        string
	Description string
}

type Project struct {
	Id          string
	Name        string
	Description string
}

type ProjectCreateData struct {
	Name        string
	Description string
}

func ProjectSerach(ctx context.Context) ([]Project, error) {
	data := []Project{
		{
			Id:          "001",
			Name:        "first-project",
			Description: "My first project",
		},
		{
			Id:          "002",
			Name:        "second-project",
			Description: "My second project",
		},
	}
	return data, nil
}

func ProjectDetailsById(ctx context.Context, id string) (*ProjectDetails, error) {
	return &ProjectDetails{
		Id:          "001",
		Name:        "Project details",
		Description: "Testing project details",
	}, nil
}

func ProjectCreate(ctx context.Context, project ProjectCreateData) (string, error) {
	return "project-created-id", nil
}
