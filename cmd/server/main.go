package main

import (
	"flag"
	"net/http"

	"github.com/remind101/todo"
)

func main() {
	var port = flag.String("port", "3000", "The port to run the server on.")

	c := todo.New()
	s := todo.NewServer(c)

	http.ListenAndServe(":"+*port, s)
}
