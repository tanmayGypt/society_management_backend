package main

import (
	"fmt"
	"log"
	"net/http"

	"society_management_backend/prisma/db"
)

// homePage function will handle requests to the root URL.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// handleRequests function will set up the routes and start the server.
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// main function is the entry point of the program.
func main() {
	if err := run(); err != nil {
		log.Fatalf("An error occurred: %v", err)
	}
}

func run() error {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	log.Println("Database connection established.")
	handleRequests()
	return nil
}
