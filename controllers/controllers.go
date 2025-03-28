package controllers

import (
	"log"
	"net/http"
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
