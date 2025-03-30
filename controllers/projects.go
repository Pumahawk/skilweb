package controllers

import (
	"log"
	"net/http"

	"github.com/pumahawk/skilweb/services"
	"github.com/pumahawk/skilweb/views"
)

func ProjectSearch(r *http.Request) (int, string, any) {
	var data views.ProjectsSearchResponse

	projects, err := services.ProjectSerach()
	if err != nil {
		log.Printf("controller project search: Unable search projects. %v", err)
		return 500, "generic", views.ServerErrorData("Search problems")
	}

	data = views.ProjectsSearchResponse{
		Title: "Projects",
		Data:  projects,
	}
	return 200, "projects-search", data
}

func ProjectDetails(r *http.Request) (int, string, any) {
	var data views.ProjectsDetailsResponse

	id := r.PathValue("id")
	if id == "" {
		return 400, "generic", views.NewGenericViewData("Bad request", "Mandatory parameter id")
	}

	project, err := services.ProjectDetailsById(id)
	if err != nil {
		if err == services.NotFound {
			return 404, "generic", views.NotFoundData("Project not found")
		}
		log.Printf("controller project search: Unable search projects. %v", err)
		return 500, "generic", views.ServerErrorData("Search problems")
	}

	data = views.ProjectsDetailsResponse{
		Title: "Projects",
		Data:  *project,
	}
	return 200, "projects-details", data
}
