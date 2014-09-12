package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/ejholmes/todo"
)

func main() {
	var port = flag.String("port", "3000", "The port to run the server on.")
	flag.Parse()

	c := todo.New()
	s := todo.NewServer(c)

	fmt.Printf("Listening on %s\n", *port)
	http.ListenAndServe(":"+*port, s)
}
