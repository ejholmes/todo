package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ejholmes/todo"
)

func main() {
	var port = flag.String("port", "3000", "The port to run the server on.")
	flag.Parse()

	c := todo.New()
	s := todo.NewServer(c)

	log.Printf("Listening on %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, s))
}
