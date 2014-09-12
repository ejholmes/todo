package todo

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	// ErrBadRequest represents a 400 error.
	ErrBadRequest = errors.New(http.StatusText(400))

	// ErrNotFound represents a 404 error.
	ErrNotFound = errors.New(http.StatusText(404))

	// ErrInternalServer represents a 500 error.
	ErrInternalServer = errors.New(http.StatusText(500))
)

// ErrorResource represents an error response.
type ErrorResource struct {
	Message string `json:"message"`
}

// HandlerFunc defines our handler function signature.
type HandlerFunc func(*Client, *ResponseWriter, *Request)

// ResponseWriter wraps an http.ResponseWriter
type ResponseWriter struct {
	http.ResponseWriter
}

// Encode encodes v into the response as json.
func (w *ResponseWriter) Encode(v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

// Error response with an error.
func (w *ResponseWriter) Error(code int, err error) error {
	w.WriteHeader(code)
	return w.Encode(&ErrorResource{Message: err.Error()})
}

// Request wraps an http.Request.
type Request struct {
	*http.Request
}

// Decode json decodes the request body into v.
func (r *Request) Decode(v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// Server handles http requests for CRUD'ing Todos.
type Server struct {
	*Client
	mux *mux.Router
}

// NewServer returns a new Server instance.
func NewServer(c *Client) *Server {
	s := &Server{Client: c, mux: mux.NewRouter()}

	s.Handle("GET", "/todos", TodosList)
	s.Handle("POST", "/todos", TodosCreate)
	s.Handle("POST", "/todos/{id}/complete", TodosComplete)
	s.Handle("DELETE", "/todos/{id}/complete", TodosUncomplete)

	return s
}

// Handle adds a handle to the router.
func (s *Server) Handle(method, path string, fn HandlerFunc) {
	s.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS")
		fn(s.Client, &ResponseWriter{w}, &Request{r})
	}).Methods(method)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// TodosList returns all Todos.
func TodosList(c *Client, w *ResponseWriter, r *Request) {
	todos, err := c.Todos.All()
	if err != nil {
		w.Error(500, ErrInternalServer)
	}

	w.Encode(todos)
}

// TodosCreate creates a new Todo.
func TodosCreate(c *Client, w *ResponseWriter, r *Request) {
	var t Todo
	if err := r.Decode(&t); err != nil {
		w.Error(400, ErrBadRequest)
		return
	}

	if _, err := c.Todos.Insert(&t); err != nil {
		w.Error(400, ErrBadRequest)
		return
	}

	w.Encode(t)
}

// TodosComplete marks a Todo as complete.
func TodosComplete(c *Client, w *ResponseWriter, r *Request) {
	withTodo(c, w, r, func(t *Todo) {
		t.Complete()
	})
}

// TodosUncomplete marks a Todo as uncomplete
func TodosUncomplete(c *Client, w *ResponseWriter, r *Request) {
	withTodo(c, w, r, func(t *Todo) {
		t.Uncomplete()
	})
}

func withTodo(c *Client, w *ResponseWriter, r *Request, fn func(*Todo)) {
	vars := mux.Vars(r.Request)

	t, err := c.Todos.Find(vars["id"])
	if err != nil {
		w.Error(400, ErrBadRequest)
		return
	}

	if t == nil {
		w.Error(404, ErrNotFound)
		return
	}

	fn(t)

	w.Encode(t)
}
