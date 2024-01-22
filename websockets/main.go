package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	setupAPI()

	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServeTLS(":8080", "certificates/server.crt","certificates/server.key", nil))
}

func setupAPI() {
	ctx := context.Background()

	manager := NewManager(ctx)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
	http.HandleFunc("/login", manager.loginHandler)
}
