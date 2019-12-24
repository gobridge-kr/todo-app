package main

import (
	"fmt"
	"net/http"

	"github.com/gobridge-kr/todo-app/server"
	"github.com/gobridge-kr/todo-app/server/controller"
	"github.com/gobridge-kr/todo-app/server/database"
	"github.com/gobridge-kr/todo-app/server/middleware"
)

const port = 8080

var baseURL = fmt.Sprintf("http://localhost:%d", port)

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
