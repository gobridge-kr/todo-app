package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port = 8080
var baseURL = fmt.Sprintf("http://localhost:%d", port)

var id = 0
var todos = make([]todo, 0)

func main() {
	http.HandleFunc("/", router)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-methods", "GET, POST, PATCH, DELETE")
	w.Header().Set("access-control-allow-headers", "accept, content-type")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id := r.URL.Path[1:]
	hasID := len(id) > 0
	log.Printf("%s %s\n", r.Method, r.URL.Path)
	if hasID {
		log.Printf("ID: %s\n", id)
	}

	switch r.Method {
	case "GET":
		if hasID {
			getOne(w, r, id)
		} else {
			getAll(w, r)
		}
	case "POST":
		if hasID {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			log.Printf("Cannot %s to this route\n", r.Method)
		} else {
			req := todoCreateRequest{}
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			create(w, r, req.Title, req.Order)
		}
	case "PATCH":
		if hasID {
			req := todoUpdateRequest{}
			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
			update(w, r, id, req.Title, req.Completed, req.Order)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			log.Printf("Cannot %s to this route\n", r.Method)
		}
	case "DELETE":
		if hasID {
			deleteOne(w, r, id)
		} else {
			deleteAll(w, r)
		}
	case "OPTIONS":
		fmt.Fprint(w, "")
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		log.Printf("No such method: %s\n", r.Method)
	}
}

func getOne(w http.ResponseWriter, r *http.Request, id string) {
	for _, value := range todos {
		if value.ID == id {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func create(w http.ResponseWriter, r *http.Request, title string, order int) {
	id++
	nextID := id
	newTodo := todo{
		ID:        fmt.Sprintf("%d", nextID),
		Title:     title,
		Completed: false,
		Order:     order,
		URL:       fmt.Sprintf("%s/%d", baseURL, nextID),
	}
	todos = append(todos, newTodo)
	json.NewEncoder(w).Encode(newTodo)
}

func update(w http.ResponseWriter, r *http.Request, id string, title string, completed bool, order int) {
	for index, value := range todos {
		if value.ID == id {
			todo := value
			todo.Title = title
			todo.Completed = completed
			todo.Order = order
			todos[index] = todo

			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func deleteOne(w http.ResponseWriter, r *http.Request, id string) {
	for index, value := range todos {
		if value.ID == id {
			todos = append(todos[:index], todos[index+1:]...)

			json.NewEncoder(w).Encode(value)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func deleteAll(w http.ResponseWriter, r *http.Request) {
	todos = make([]todo, 0)
	json.NewEncoder(w).Encode(todos)
}

type todoCreateRequest struct {
	Title string `json:"title"`
	Order int    `json:"order"`
}

type todoUpdateRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
}

type todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	URL       string `json:"url"`
}
