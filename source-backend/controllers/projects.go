package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pumahawk/skilweb/server"
	"github.com/pumahawk/skilweb/services"
)

func ProjectSearch(r *http.Request) server.ControllerResponse[any] {
	projects, err := services.ProjectSearch(r.Context())
	if err != nil {
		log.Printf("controller project search: Unable search projects. %v", err)
		return server.MessageResponse(500, "Search problems")
	}

	var response []ProjectDTO
	for _, pr := range projects {
		response = append(response, ProjectDTO{
			Id: pr.Id,
			Name: pr.Name,
			Description: pr.Description,
		})
	}
	return server.NewResponse(200, response)
}

func ProjectDetails(r *http.Request) server.ControllerResponse[any] {

	id := r.PathValue("id")
	if id == "" {
		return server.MessageResponse(400, "Mandatory parameter id")
	}

	project, err := services.ProjectDetailsById(r.Context(), id)
	if err != nil {
		if err == services.NotFound {
			return server.MessageResponse(404, "Project not found")
		}
		log.Printf("controller project search: Unable search projects. %v", err)
		return server.MessageResponse(500,"Search problems")
	}

	response := ProjectDetailsDTO{
		Id: project.Id,
		Name: project.Name,
		Description: project.Description,
	}
	return server.NewResponse(200, response)
}

func ProjectCreate(r *http.Request) server.ControllerResponse[any] {
	var body ProjectCreateRequestDTO
	server.ReadBody(r.Body, body)

	if body.Name == "" {
		return server.MessageResponse(400, "Mandatory parameter name")
	}
	if body.Description == "" {
		return server.MessageResponse(400, "Mandatory parameter description")
	}

	project := services.ProjectCreateData{
		Name:        body.Name,
		Description: body.Description,
	}

	id, err := services.ProjectCreate(r.Context(), project)
	if err != nil {
		return server.MessageResponse(500, "project service: Project creation Fail")
	}
	return server.MessageResponse(201, fmt.Sprintf("New project: %s", id))
}

