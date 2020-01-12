package server

import (
	"fmt"
	"net/http"

	"github.com/gobridge-kr/todo-app/server/controller"
)

// Server represents current server status
type Server struct {
	baseURL     string
	middlewares []func(w http.ResponseWriter, r *http.Request)
}

// New creates a new Server with given URL
func New(baseURL string) *Server {
	return &Server{
		baseURL: baseURL,
	}
}

// Middleware configures middleware to process requests
func (s *Server) Middleware(middleware func(w http.ResponseWriter, r *http.Request)) {
	s.middlewares = append(s.middlewares, middleware)
}

// Route configures routing map
func (s *Server) Route(path string, controller controller.Controller) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		for _, middleware := range s.middlewares {
			middleware(w, r)
		}

		id := r.URL.Path[1:]
		hasID := len(id) > 0

		switch r.Method {
		case "GET":
			if hasID {
				controller.GetOne(w, r, id)
			} else {
				controller.GetAll(w, r)
			}
		case "POST":
			if hasID {
				controller.PostOne(w, r, id)
			} else {
				controller.PostAll(w, r)
			}
		case "PATCH":
			if hasID {
				controller.PatchOne(w, r, id)
			} else {
				controller.PatchAll(w, r)
			}
		case "DELETE":
			if hasID {
				controller.DeleteOne(w, r, id)
			} else {
				controller.DeleteAll(w, r)
			}
		case "OPTIONS":
			controller.Options(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

// Serve starts the actual serving job
func (s *Server) Serve(port int) {
	http.ListenAndServe(":"+fmt.Sprint(port), nil)
}
