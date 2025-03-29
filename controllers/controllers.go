package controllers

import (
	"log"
	"net/http"

	"github.com/pumahawk/skilweb/views"
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
	return 404, "generic", views.NotFoundData("Page not found")
}
