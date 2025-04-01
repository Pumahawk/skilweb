package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Controller = func(r *http.Request) ControllerResponse[any]

type ControllerResponse[T any] struct {
	Code int
	Data T
}

type MessageData struct {
	Message string
}

func ControllerViewHandler(controller Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			log.Printf("main controller view: Request context closed. %v", r.Context().Err())
		default:
			res := controller(r)
			w.WriteHeader(res.Code)
			if err := json.NewEncoder(w).Encode(res.Data); err != nil {
				log.Printf("server controller handler: Invalid json response. %v", err)
			}
		}
	}
}

func NewResponse(code int, data any) ControllerResponse[any] {
	return ControllerResponse[any]{
		Code: code,
		Data: data,
	}
}

func MessageResponse(code int, message string) ControllerResponse[any]  {
	return ControllerResponse[any]{
		Code: code,
		Data: MessageData{
			Message: message,
		},
	}
}

func ReadBody[T any](r io.Reader, v T) error {
	return json.NewDecoder(r).Decode(v)
}
