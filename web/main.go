package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./ui/public")))

	log.Println("Starting application on :8080")
	http.ListenAndServe(":8080", nil)
}
