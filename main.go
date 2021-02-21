package main

import (
	"fmt"
	"net/http"
	"context"
	"os"
	"log"

	"github.com/gorilla/mux"
	"cloud.google.com/go/storage"
)

// fmt.Printf("%+v\n", p)

var (
	ctx context.Context
	client *storage.Client
)

type Endpoint struct {
	Title string `json:"title"`
	Endpoints []string `json:"endpoints"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to aimi, stay enjoy !")
}

func Handler() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", MainHandler)

	// path
	r.HandleFunc("/users/", HomeUsersHandler)
	r.HandleFunc("/users/avatar/{id}/{hash}", UserHandler)

	// Test
	r.HandleFunc("/users/test", UserTestHandler)

	return r
}

func main() {
	defer client.Close()

	http.Handle("/", Handler())

	port := os.Getenv("PORT")

    if port == "" {
        port = "8080"
        log.Printf("defaulting to port %s\n", port)
	}

	http.ListenAndServe(":" + port, nil)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}