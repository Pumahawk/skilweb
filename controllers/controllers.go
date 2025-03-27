package controllers

import "net/http"

type HelloMessage struct {
	Message string
}

func HelloWorld(r *http.Request) (string, any) {
	return "hello.html", HelloMessage{"Hello, Wolrd!"}
}
