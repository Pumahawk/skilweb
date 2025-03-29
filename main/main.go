package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pumahawk/skilweb/controllers"
	"github.com/pumahawk/skilweb/server"
	"github.com/pumahawk/skilweb/views"
)

func main() {
	log.Println("Start skilweb")
	conf, err := LoadFlags()
	if err != nil {
		log.Fatalf("main: Invalid flags. %v", err)
	}
	err = StartHttpServer(conf)
	if err != nil {
		log.Fatalf("main: Unable start HttpServer. %v", err)
	}
}

type HttpChain = func(http.HandlerFunc) http.HandlerFunc

func LoadFlags() (*Conf, error) {
	var conf Conf
	flag.StringVar(&conf.Address, "address", ":8000", "Http server address")
	flag.Parse()
	return &conf, nil
}

func StartHttpServer(conf *Conf) error {
	err := LoadServerControllers()
	if err != nil {
		return fmt.Errorf("main server: Unable to load controllers. %w", err)
	}

	address := conf.Address
	log.Printf("main: Start web server, [Address=%s]", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		return fmt.Errorf("main server: Unable to startup http server. %w", err)
	}
	return nil
}

type Conf struct {
	Address string
}

func LoadServerControllers() error {
	http.HandleFunc("/hello", BaseChain(ControllerViewHandler(controllers.HelloWorld)))
	http.HandleFunc("/projects/search", BaseChain(ControllerViewHandler(controllers.ProjectsSearch)))
	http.HandleFunc("/", BaseChain(ControllerViewHandler(controllers.NotFound)))
	return nil
}

func ControllerViewHandler(controller server.Controller) http.HandlerFunc {
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

func LogHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("http_request: %s", r.URL.Path)
		handler(NewLogResponseWriter(w), r)
	}
}

type LogResponseWriter struct {
	http.ResponseWriter
}

func (lw *LogResponseWriter) WriteHeader(code int) {
	log.Printf("http_response: code %d", code)
	lw.ResponseWriter.WriteHeader(code)
}

func NewLogResponseWriter(w http.ResponseWriter) http.ResponseWriter {
	return &LogResponseWriter{w}
}

func Chain(handler http.HandlerFunc, chain ...HttpChain) http.HandlerFunc {
	if len(chain) > 0 {
		return chain[0](Chain(handler, chain[1:]...))
	} else {
		return func(w http.ResponseWriter, r *http.Request) {
			handler(w, r)
		}
	}
}

func BaseChain(handler http.HandlerFunc) http.HandlerFunc {
	return Chain(handler, LogHandler)
}
