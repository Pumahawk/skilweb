package server

import (
	"bytes"
	"log"
	"net/http"
	"regexp"

	"github.com/pumahawk/skilweb/views"
)

type Controller = func(r *http.Request) (int, string, any)

func ControllerViewHandler(controller Controller) http.HandlerFunc {
	vs := views.LoadViews(views.LinksFuncMap())
	return func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			log.Printf("main controller view: Request context closed. %v", r.Context().Err())
		default:
			code, name, data := controller(r)
			regx := regexp.MustCompile("^redirect:(..*)")
			if regx.MatchString(name) {
				path := regx.FindStringSubmatch(name)
				if len(path) != 2 {
					log.Printf("Inavalid redirect value: %s", name)
					w.WriteHeader(500)
					return
				}

				url := path[1]
				http.Redirect(w, r, url, http.StatusSeeOther)
				return
			}
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
