package todo

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

// A Todo represents a thing you gotta do.
type Todo struct {
	ID string `json:"id"`

	// A text description of the thing you gotta do.
	Text string `json:"text"`

	// When it was completed.
	CompletedAt *time.Time `json:"completed_at"`
}

// TodosService provides methods for CRUD'ing todos.
type TodosService struct {
	todos []*Todo
}

// NewTodosService returns a new TodosService instance.
func NewTodosService() *TodosService {
	return &TodosService{todos: make([]*Todo, 0)}
}

// All returns all Todos.
func (s *TodosService) All() ([]*Todo, error) {
	return s.todos, nil
}

// Insert inserts a Todo.
func (s *TodosService) Insert(t *Todo) (*Todo, error) {
	t.ID = uuid.New()
	s.todos = append(s.todos, t)
	return t, nil
}

// Create initializes a new Todo and inserts it.
func (s *TodosService) Create(text string) (*Todo, error) {
	return s.Insert(&Todo{Text: text})
}
