package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/pumahawk/skilweb/controllers"
	"github.com/pumahawk/skilweb/server"
	"github.com/pumahawk/skilweb/services"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("Start skilweb")
	conf, err := LoadFlags()
	if err != nil {
		log.Fatalf("main: Invalid flags. %v", err)
	}
	db, err := GetDB(conf)
	if err != nil {
		log.Fatalf("main: Unable get DB object. %v", err)
	}
	err = StartHttpServer(conf, db)
	if err != nil {
		log.Fatalf("main: Unable start HttpServer. %v", err)
	}
}

type HttpChain = func(http.HandlerFunc) http.HandlerFunc

func LoadFlags() (*Conf, error) {
	var conf Conf
	flag.StringVar(&conf.Address, "address", ":8000", "Http server address")
	flag.StringVar(&conf.DBConn, "db", "", "Database access string")
	flag.Parse()
	return &conf, nil
}

func StartHttpServer(conf *Conf, db *sql.DB) error {
	err := LoadServerControllers(db)
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
	DBConn  string
}

func LoadServerControllers(db *sql.DB) error {
	http.HandleFunc("/hello", BaseChain(db, server.ControllerViewHandler(controllers.HelloWorld)))
	http.HandleFunc("POST /projects", BaseChain(db, server.ControllerViewHandler(controllers.ProjectCreate)))
	http.HandleFunc("GET /projects", BaseChain(db, server.ControllerViewHandler(controllers.ProjectCreateForm)))
	http.HandleFunc("GET /projects/search", BaseChain(db, server.ControllerViewHandler(controllers.ProjectSearch)))
	http.HandleFunc("GET /projects/{id}", BaseChain(db, server.ControllerViewHandler(controllers.ProjectDetails)))
	http.HandleFunc("/", BaseChain(db, server.ControllerViewHandler(controllers.NotFound)))
	return nil
}

type LogResponseWriter struct {
	http.ResponseWriter
}

func LogHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("http_request: %s", r.URL.Path)
		handler(NewLogResponseWriter(w), r)
	}
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

func BaseChain(db *sql.DB, handler http.HandlerFunc) http.HandlerFunc {
	return Chain(handler, LogHandler, DatabaseChain(db))
}

func DatabaseChain(db *sql.DB) HttpChain {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			conn, err := db.Conn(r.Context())
			if err != nil {
				log.Printf("main db handler: Unable to open request db connection. %v", err)
			}
			rctx := context.WithValue(r.Context(), services.DBConnK, conn)
			rc := r.WithContext(rctx)
			handler(w, rc)
		}
	}
}

func GetDB(conf *Conf) (*sql.DB, error) {
	db, err := sql.Open("postgres", conf.DBConn)
	if err != nil {
		return nil, fmt.Errorf("main: Unable to init db object. %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("main: Unable to ping database. %w", err)
	}

	return db, nil
}
