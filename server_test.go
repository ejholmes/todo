package todo

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func init() {
	GenID = func() string {
		return "1234"
	}
}

func MustRequest(method, path string) *http.Request {
	req, _ := http.NewRequest(method, path, nil)
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
			Request: MustRequest("GET", "/todos"),
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
			Request: MustRequest("GET", "/todos"),
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
