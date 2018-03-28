package main

import (
	"database/sql"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/boseji/mserver"
)

var (
	// Global Template Storage
	tmpl *template.Template

	// Flags for parsing
	port            = flag.Int("port", 8080, "http port to listen on")
	shutdownTimeout = flag.Duration("shutdown-timeout", 10*time.Second,
		"shutdown timeout (5s,5m,5h) before connections are cancelled")

	// Storage for Databse
	db *sql.DB

	// Storage for Table Creation status
	tableCreated = false

	// Web Server
	server *mserver.Mserver
)

func init() {

	// Parse All Templates
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	// Register Default Handlers
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", root)
	http.HandleFunc("/create", createTable)
	http.HandleFunc("/readall", readAll)
	http.HandleFunc("/find", findRecord)
	http.HandleFunc("/update", updateRecord)
	http.HandleFunc("/delete", deleteRecord)
	http.HandleFunc("/drop", dropTable)

	// Parse the Command-line
	flag.Parse()

	// Database
	err := connectDatabase()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	// Stop Database Connection at the End
	defer disconnectDatabase()

	// Start the Server
	server = mserver.
		NewMserver(fmt.Sprintf(":%d", *port), *shutdownTimeout)

	server.GracefulStop(true)

	log.Printf(" Server down.")
}
