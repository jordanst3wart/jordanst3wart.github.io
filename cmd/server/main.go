package main

import (
	"log"
	"net/http"
)

func main() {
	// 1. Define the directory to serve
	fileServer := http.FileServer(http.Dir("./../../public"))

	// 2. Register the handler for all requests
	http.Handle("/", fileServer)

	// 3. Start the server on port 8080
	log.Println("Server starting on :8090...")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
