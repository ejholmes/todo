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

// Complete marks the Todo as complete.
func (t *Todo) Complete() {
	now := time.Now()
	t.CompletedAt = &now
}

// Uncomplete clears the completed status.
func (t *Todo) Uncomplete() {
	t.CompletedAt = nil
}

// TodosService provides methods for CRUD'ing todos.
type TodosService struct {
	todos map[string]*Todo
}

// NewTodosService returns a new TodosService instance.
func NewTodosService() *TodosService {
	return &TodosService{todos: make(map[string]*Todo)}
}

// All returns all Todos.
func (s *TodosService) All() ([]*Todo, error) {
	todos := make([]*Todo, 0, len(s.todos))

	for _, t := range s.todos {
		todos = append(todos, t)
	}

	return todos, nil
}

// Find finds a single Todo by id.
func (s *TodosService) Find(id string) (*Todo, error) {
	t := s.todos[id]
	return t, nil
}

// Delete delets a Todo by id.
func (s *TodosService) Delete(id string) (*Todo, error) {
	t, err := s.Find(id)
	if err != nil {
		return nil, err
	}

	delete(s.todos, t.ID)
	return t, nil
}

// Insert inserts a Todo.
func (s *TodosService) Insert(t *Todo) (*Todo, error) {
	t.ID = uuid.New()
	s.todos[t.ID] = t
	return t, nil
}

// Create initializes a new Todo and inserts it.
func (s *TodosService) Create(text string) (*Todo, error) {
	return s.Insert(&Todo{Text: text})
}
