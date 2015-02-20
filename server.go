package todo

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

var routes = map[path]HandlerFunc{
	"GET    /todos":               TodosList,
	"POST   /todos":               TodosCreate,
	"DELETE /todos/{id}":          TodosDelete,
	"PATCH  /todos/{id}":          TodosUpdate,
	"POST   /todos/{id}/complete": TodosComplete,
	"DELETE /todos/{id}/complete": TodosUncomplete,
}

type path string

func (p path) extract() (method, path string) {
	c := regexp.MustCompile(`\s+`).Split(string(p), -1)
	return c[0], c[1]
}

// ErrorResource represents an error response.
type ErrorResource struct {
	Code    int    `json:"code"`
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
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// Error response with an error.
func (w *ResponseWriter) Error(code int) error {
	w.WriteHeader(code)
	return w.Encode(
		&ErrorResource{
			Code:    code,
			Message: http.StatusText(code),
		},
	)
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

	for path, h := range routes {
		m, p := path.extract()
		s.Handle(m, p, h)
	}

	return s
}

// Handle adds a handle to the router.
func (s *Server) Handle(method, path string, fn HandlerFunc) {
	s.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fn(s.Client, &ResponseWriter{w}, &Request{r})
	}).Methods(method)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	} else {
		s.mux.ServeHTTP(w, r)
	}
}

// TodosList returns all Todos.
func TodosList(c *Client, w *ResponseWriter, r *Request) {
	todos, err := c.Todos.All()
	if err != nil {
		w.Error(http.StatusInternalServerError)
	}

	w.Encode(todos)
}

// TodosCreate creates a new Todo.
func TodosCreate(c *Client, w *ResponseWriter, r *Request) {
	var t Todo
	if err := r.Decode(&t); err != nil {
		w.Error(http.StatusBadRequest)
		return
	}

	if _, err := c.Todos.Insert(&t); err != nil {
		w.Error(http.StatusBadRequest)
		return
	}

	w.Encode(t)
}

// TodosDelete deletes a Todo.
func TodosDelete(c *Client, w *ResponseWriter, r *Request) {
	withTodo(c, w, r, func(t *Todo) {
		c.Todos.Delete(t.ID)
	})
}

// TodosUpdate updates a Todo.
func TodosUpdate(c *Client, w *ResponseWriter, r *Request) {
	var u Todo
	if err := r.Decode(&u); err != nil {
		w.Error(http.StatusBadRequest)
		return
	}

	withTodo(c, w, r, func(t *Todo) {
		t.Text = u.Text
	})
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
		w.Error(http.StatusBadRequest)
		return
	}

	if t == nil {
		w.Error(http.StatusNotFound)
		return
	}

	fn(t)

	w.Encode(t)
}
