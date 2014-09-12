package todo

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	GenID = func() string {
		return "1234"
	}
}

func MustRequest(method, path string, r io.Reader) *http.Request {
	req, _ := http.NewRequest(method, path, r)
	return req
}

type RequestTest struct {
	Before  func(*Client)
	Request *http.Request
	Check   func(*httptest.ResponseRecorder)
}

func TestTodosAll(t *testing.T) {
	tests := []RequestTest{
		{
			Before:  func(c *Client) {},
			Request: MustRequest("GET", "/todos", nil),
			Check: func(resp *httptest.ResponseRecorder) {
				body := "[]\n"

				if resp.Body.String() != body {
					t.Errorf("Body => %s; want %s", resp.Body.String(), body)
				}
			},
		},

		{
			Before: func(c *Client) {
				c.Todos.Create("Hello")
			},
			Request: MustRequest("GET", "/todos", nil),
			Check: func(resp *httptest.ResponseRecorder) {
				body := `[{"id":"1234","text":"Hello","completed_at":null}]` + "\n"

				if resp.Body.String() != body {
					t.Errorf("Body => %s; want %s", resp.Body.String(), body)
				}
			},
		},
	}

	for _, tt := range tests {
		c := New()
		s := NewServer(c)

		resp := httptest.NewRecorder()
		req := tt.Request

		tt.Before(c)
		s.ServeHTTP(resp, req)
		tt.Check(resp)
	}
}

func TestTodosCreate(t *testing.T) {
	tests := []RequestTest{
		{
			Before:  func(c *Client) {},
			Request: MustRequest("POST", "/todos", strings.NewReader(`{"text":"Hello"}`)),
			Check: func(resp *httptest.ResponseRecorder) {
				body := `{"id":"1234","text":"Hello","completed_at":null}` + "\n"

				if resp.Body.String() != body {
					t.Errorf("Body => %s; want %s", resp.Body.String(), body)
				}
			},
		},
	}

	for _, tt := range tests {
		c := New()
		s := NewServer(c)

		resp := httptest.NewRecorder()
		req := tt.Request

		tt.Before(c)
		s.ServeHTTP(resp, req)
		tt.Check(resp)
	}
}

func TestTodosDelete(t *testing.T) {
	tests := []RequestTest{
		{
			Before: func(c *Client) {
				c.Todos.Create("Hello")
			},
			Request: MustRequest("DELETE", "/todos/1234", strings.NewReader(`{"text":"Hello"}`)),
			Check: func(resp *httptest.ResponseRecorder) {
				body := `{"id":"1234","text":"Hello","completed_at":null}` + "\n"

				if resp.Body.String() != body {
					t.Errorf("Body => %s; want %s", resp.Body.String(), body)
				}
			},
		},
	}

	for _, tt := range tests {
		c := New()
		s := NewServer(c)

		resp := httptest.NewRecorder()
		req := tt.Request

		tt.Before(c)
		s.ServeHTTP(resp, req)
		tt.Check(resp)
	}
}

func TestTodosUpdate(t *testing.T) {
	tests := []RequestTest{
		{
			Before: func(c *Client) {
				c.Todos.Create("Hello")
			},
			Request: MustRequest("PATCH", "/todos/1234", strings.NewReader(`{"text":"Hello World"}`)),
			Check: func(resp *httptest.ResponseRecorder) {
				body := `{"id":"1234","text":"Hello World","completed_at":null}` + "\n"

				if resp.Body.String() != body {
					t.Errorf("Body => %s; want %s", resp.Body.String(), body)
				}
			},
		},
	}

	for _, tt := range tests {
		c := New()
		s := NewServer(c)

		resp := httptest.NewRecorder()
		req := tt.Request

		tt.Before(c)
		s.ServeHTTP(resp, req)
		tt.Check(resp)
	}
}

func TestPath(t *testing.T) {
	tests := []struct {
		in     string
		method string
		path   string
	}{
		{"GET   /todos", "GET", "/todos"},
		{"POST  /todos", "POST", "/todos"},
		{"PATCH /todos/{id}", "PATCH", "/todos/{id}"},
	}

	for _, tt := range tests {
		m, p := path(tt.in).extract()

		if m != tt.method {
			t.Errorf("Method => %s; want %s", m, tt.method)
		}

		if p != tt.path {
			t.Errorf("Path => %s; want %s", p, tt.path)
		}
	}
}
