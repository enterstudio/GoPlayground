package main

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Page Info Structure
type Page struct {
	Title   string
	Body    string
	Labels  [5]string
	Data    [5]float32
	records int
}

// Templates Storage
var templates = template.New("")

func init() {
	// Load the Template into memory
	_, err := templates.ParseFiles("templates\\display.html")
	if err != nil {
		log.Fatalln(err)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	// Get the URI
	loc := html.EscapeString(r.URL.Path)
	if loc == "/view" {
		http.NotFound(w, r)
		return
	}
	// Strip the Prefix
	loc = strings.TrimLeft(loc, "/view/")
	// Reject the No Names and multiple slash inputs
	if len(loc) == 0 || strings.Index(loc, "/") != -1 {
		http.NotFound(w, r)
		return
	}

	// If all is ok then Begin porcessing

	s := fmt.Sprintf("Hello, %q", loc)

	// Create the Page info
	info1 := Page{
		Title: loc,
		Body:  s,
		/*
			Labels: [5]string{
				"P1", "P2", "P3", "P4", "P5",
			},
			Data: [5]float32{65, 59, 80, 81, 56},
		*/
	}
	// Generate the Random Data
	rand.Seed(int64(len(loc)) + time.Now().Unix())
	for i := 0; i < 5; i++ {
		info1.Data[i] = float32(rand.Intn(100))
		info1.Labels[i] = fmt.Sprintf("P%d", i)
	}
	log.Printf("VIEW URI: %q\n", r.URL.Path)
	//log.Printf(" %s\n", loc)

	// Execute the Template
	err := templates.ExecuteTemplate(w, "display.html", &info1)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/view/", ViewHandler)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.Handle("/", http.FileServer(http.Dir("www")))
	log.Println("Starting Server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
