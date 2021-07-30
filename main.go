package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle(
		"/",
		Handlers(),
	)
	log.Fatal(http.ListenAndServe(":"+"8020", nil))
}
