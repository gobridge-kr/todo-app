package database

import (
	"errors"
	"fmt"

	"github.com/gobridge-kr/todo-app/server/model"
)

var (
	// ErrItemNotFound means the requested item couldn't be found
	ErrItemNotFound = errors.New("Item Not Found")
	// ErrBadRequest means request data was ill-formed, or invalid
	ErrBadRequest = errors.New("Bad Request")
)

// Database represents current database state
type Database struct {
	idCounter int
	todos     []model.Todo
	config    Config
}

// Config is configuration data for a database
type Config struct {
	BaseURL string
}

// New creates a Database with given Config
func New(config Config) *Database {
	return &Database{
		idCounter: 0,
		todos:     make([]model.Todo, 0),
		config:    config,
	}
}

// GetTodos returns all todos
func (d *Database) GetTodos() []model.Todo {
	return d.todos
}

// GetTodo retrieves a todo by given ID
func (d *Database) GetTodo(id string) (model.Todo, error) {
	for _, value := range d.todos {
		if value.ID == id {
			return value, nil
		}
	}
	return model.Todo{}, ErrItemNotFound
}

// AddTodo creates a todo with params map
func (d *Database) AddTodo(params map[string]interface{}) (model.Todo, error) {
	d.idCounter++
	title, ok := params["title"].(string)
	if !ok {
		return model.Todo{}, ErrBadRequest
	}
	order, ok := params["order"].(float64)
	if !ok {
		order = 0
	}
	todo := model.Todo{
		ID:        fmt.Sprint(d.idCounter),
		Title:     title,
		Completed: false,
		Order:     int64(order),
		URL:       fmt.Sprintf("%s/%d", d.config.BaseURL, d.idCounter),
	}
	d.todos = append(d.todos, todo)
	return todo, nil
}

// UpdateTodo updates a todo with given ID
func (d *Database) UpdateTodo(id string, params map[string]interface{}) (model.Todo, error) {
	for index, value := range d.todos {
		if value.ID == id {
			todo := value
			if title, ok := params["title"].(string); ok {
				todo.Title = title
			}
			if completed, ok := params["completed"].(bool); ok {
				todo.Completed = completed
			}
			if order, ok := params["order"].(float64); ok {
				todo.Order = int64(order)
			}
			d.todos[index] = todo
			return todo, nil
		}
	}
	return model.Todo{}, ErrItemNotFound
}

// DeleteTodo deletes a todo with given ID
func (d *Database) DeleteTodo(id string) (model.Todo, error) {
	for index, value := range d.todos {
		if value.ID == id {
			d.todos = append(d.todos[:index], d.todos[index+1:]...)
			return value, nil
		}
	}
	return model.Todo{}, ErrItemNotFound
}

// DeleteTodos deletes all todos
func (d *Database) DeleteTodos() {
	d.todos = make([]model.Todo, 0)
}
