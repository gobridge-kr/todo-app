package main

import (
	"net/http"
	"os"

	"github.com/gobridge-kr/todo-app/server"
	"github.com/gobridge-kr/todo-app/server/controller"
	"github.com/gobridge-kr/todo-app/server/database"
	"github.com/gobridge-kr/todo-app/server/middleware"
)

var (
	port    = "8080"
	baseURL = "http://localhost:" + port
)

func init() {
	if env := os.Getenv("PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("BASE_URL"); env != "" {
		baseURL = env
	}
}

func main() {
	dbConfig := database.Config{
		BaseURL: baseURL,
	}
	db := database.New(dbConfig)
	c := controller.Todo(db)
	s := server.New(baseURL)

	s.Middleware(func(w http.ResponseWriter, r *http.Request) { middleware.Cors(w) })

	s.Route("/", c)

	s.Serve(port)
}
