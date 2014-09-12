package todo

import (
	"time"

	"code.google.com/p/go-uuid/uuid"
)

// TodosService is our interface for CRUD'ing Todo's.
type TodosService interface {
	// All returns a slice of all the todos in the store.
	All() ([]*Todo, error)

	// Create creates a new Todo and inserts it.
	Create(text string) (*Todo, error)

	// Find finds a single Todo by id.
	Find(id string) (*Todo, error)

	// Delete removes a single Todo by id.
	Delete(id string) (*Todo, error)

	// Insert inserts the Todo into the store.
	Insert(*Todo) (*Todo, error)
}

// GenID is used to generate a unique id.
var GenID = func() string {
	return uuid.New()
}

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

// todosService is an implementation of the TodosService interface
// that stores Todos in memory.
type todosService struct {
	todos map[string]*Todo
}

// NewTodosService returns a new TodosService instance.
func NewTodosService() TodosService {
	return &todosService{todos: make(map[string]*Todo)}
}

// All returns all Todos.
func (s *todosService) All() ([]*Todo, error) {
	todos := make([]*Todo, 0, len(s.todos))

	for _, t := range s.todos {
		todos = append(todos, t)
	}

	return todos, nil
}

// Find finds a single Todo by id.
func (s *todosService) Find(id string) (*Todo, error) {
	t := s.todos[id]
	return t, nil
}

// Delete delets a Todo by id.
func (s *todosService) Delete(id string) (*Todo, error) {
	t, err := s.Find(id)
	if err != nil {
		return nil, err
	}

	delete(s.todos, t.ID)
	return t, nil
}

// Insert inserts a Todo.
func (s *todosService) Insert(t *Todo) (*Todo, error) {
	t.ID = GenID()
	s.todos[t.ID] = t
	return t, nil
}

// Create initializes a new Todo and inserts it.
func (s *todosService) Create(text string) (*Todo, error) {
	return s.Insert(&Todo{Text: text})
}
