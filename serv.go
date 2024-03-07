package main

import (
	"fmt"
	"net/http"
)

var serverPort = 8080

func main() {
	// Register handler
	http.Handle("/", http.FileServer(http.Dir("./web")))

	// Start HTTP server
	fmt.Printf("Listen on port '%d'.\n", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
	if err != nil {
		fmt.Printf("Could not start server! (Error: %s)\n", err)
	}
}
