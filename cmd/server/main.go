package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/ejholmes/todo"
)

func main() {
	var port = flag.String("port", stringVar("PORT", "3000"), "The port to run the server on.")
	flag.Parse()

	c := todo.New()
	s := todo.NewServer(c)

	log.Printf("Listening on %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, s))
}

func stringVar(key, fb string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}

	return fb
}
