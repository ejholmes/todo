package main

import (
	"net/http"

	"github.com/remind101/todo"
)

func main() {
	c := todo.New()
	s := todo.NewServer(c)

	http.ListenAndServe(":3000", s)
}
