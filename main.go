package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/sahilchauhan0603/society_management_backend/router"
)


// main function is the entry point of the program.
func main() {
	fmt.Println("SQL DATABASE !!")

	r := router.Router()
	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("Listening at port 8000 ...")
}
