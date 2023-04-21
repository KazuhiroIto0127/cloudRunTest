package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type User struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"e-mail"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")
	u := User{ Id : 1, Name : "Thome", Email : "thome@example.com"}

    json, err := json.Marshal(u)
    if err != nil {
        log.Fatal(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    w.Write(json)
}

func main() {
	log.Print("helloworld: starting server...")
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
