package controllers

import (
	"log"
	"net/http"
)

type HelloMessage struct {
	Title   string
	Message string
}

func HelloWorld(r *http.Request) (string, any) {
	log.Println("controller helloworld: Incoming request")
	return "hello", HelloMessage{"Hello page", "Hello, World!"}
}

func NotFound(r *http.Request) (string, any) {
	return "404", nil
}
