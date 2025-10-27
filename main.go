package main

import (
	"context"
	"log"
	"net/http"
)

func main() {

	client = connectMongo()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("error disconnecting mongo: %v", err)
		}
	}()
    
	mux := http.NewServeMux()
	mux.HandleFunc("/register", registerHandler)
	mux.HandleFunc("/login", loginHandler)

	mux.Handle("/profile", authMiddleware(http.HandlerFunc(profileHandler)))

	log.Println("server running on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}