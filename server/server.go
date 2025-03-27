package server

import "net/http"

type Controller = func(r *http.Request) (string, any)
