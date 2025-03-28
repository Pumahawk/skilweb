package server

import "net/http"

type Controller = func(r *http.Request) (int, string, any)
