package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gobridge-kr/todo-app/server/database"
)

type TodoController struct {
	database *database.Database
}

func (t *TodoController) GetOne(w http.ResponseWriter, r *http.Request, id string) {
	todo, err := t.database.GetTodo(id)
	if err != nil {
		if err == database.ErrItemNotFound {
			http.Error(w, "Item Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (t *TodoController) GetAll(w http.ResponseWriter, r *http.Request) {
	todos := t.database.GetTodos()
	json.NewEncoder(w).Encode(todos)
}

func (t *TodoController) PostOne(w http.ResponseWriter, r *http.Request, id string) {
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
}

func (t *TodoController) PostAll(w http.ResponseWriter, r *http.Request) {
	var params map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	todo, err := t.database.AddTodo(params)
	if err != nil {
		if err == database.ErrBadRequest {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (t *TodoController) PatchOne(w http.ResponseWriter, r *http.Request, id string) {
	var params map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	todo, err := t.database.UpdateTodo(id, params)
	if err != nil {
		if err == database.ErrItemNotFound {
			http.Error(w, "Item Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (t *TodoController) PatchAll(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
}

func (t *TodoController) DeleteOne(w http.ResponseWriter, r *http.Request, id string) {
	todo, err := t.database.DeleteTodo(id)
	if err != nil {
		if err == database.ErrItemNotFound {
			http.Error(w, "Item Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func (t *TodoController) DeleteAll(w http.ResponseWriter, r *http.Request) {
	t.database.DeleteTodos()
}

func (t *TodoController) Options(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "")
}

func Todo(database *database.Database) *TodoController {
	return &TodoController{
		database: database,
	}
}
