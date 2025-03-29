package server

import (
	"bytes"
	"log"
	"net/http"

	"github.com/pumahawk/skilweb/controllers"
	"github.com/pumahawk/skilweb/views"
)

type Controller = func(r *http.Request) (int, string, any)

func ControllerViewHandler(controller Controller) http.HandlerFunc {
	vs := views.LoadViews(controllers.LinksFuncMap())
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			log.Printf("main controller view: Request context closed. %v", r.Context().Err())
		default:
			code, name, data := controller(r)
			var bf bytes.Buffer
			err := views.Render(vs, &bf, name, data)
			if err != nil {
				log.Printf("main controller view: Unable rendering view, [Path=%s]. %v", r.URL.Path, err)
				w.WriteHeader(500)
				return
			}
			w.WriteHeader(code)
			_, err = bf.WriteTo(w)
			if err != nil {
				log.Printf("main controller view: Unable to write response, [Path=%s]. %v", r.URL.Path, err)
				return
			}
		}
	}
}
