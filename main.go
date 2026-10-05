package main

import (
	"fmt"
	"net/http"
	"notesAPI/handlers"
)

func main() {
	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/echo", handlers.EchoHandler)
	http.HandleFunc("/notes", handlers.NotesHandlers)

	fmt.Println("server start on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("failed start server:", err)
	}

	fmt.Println("server cancelled")
}
