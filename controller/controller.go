package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const (
	DBUsername = "root"
	DBPassword = "2003@society"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "societydb"
)

// connect with mongoDB
func init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName) // Data source name

	// open database
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// assign database connection to global variable
	db = database

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	insert, err := db.Query("INSERT INTO `societydb`.`studentprofile` (`EnrollmentNo`, `FirstName`, `LastName`) VALUES ('2','Ben','Ten');")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	fmt.Println("Connected to database!")
}

func ServeHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	w.Write([]byte("<h1>Welcome to SOCIETY</h1>"))
}
