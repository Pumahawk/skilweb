package controllers

import (
	"log"
	"net/http"

	"github.com/pumahawk/skilweb/services"
)

type HelloMessage struct {
	Title   string
	Message string
}

func HelloWorld(r *http.Request) (int, string, any) {
	log.Println("controller helloworld: Incoming request")
	return 200, "hello", HelloMessage{"Hello page", "Hello, World!"}
}

func NotFound(r *http.Request) (int, string, any) {
	return 404, "404", nil
}

func ProjectsSearch(r *http.Request) (int, string, any) {
	var data ProjectsSearchResponse

	projects, err := services.ProjectSerach()
	if err != nil {
		log.Printf("controller project search: Unable serach projects. %v", err)
		return 500, "500", nil
	}

	data = ProjectsSearchResponse{
		Data: projects,
	}
	return 200, "projects-search", data
}
