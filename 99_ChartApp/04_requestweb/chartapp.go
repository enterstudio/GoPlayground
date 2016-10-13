package main

import (
	"crypto/sha1"
	"fmt"
	"html"
	"html/template"
	"io"
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

const (
	MaxApiKey = 40
	MaxGroups = 4
)

// Templates Storage
var templates = template.New("")

// Map for Data Types
var DataSampleTypes = map[string]int{
	"string":  0,
	"number":  1,
	"boolean": 2,
}

// Hash Groups
var hash_arr = make(map[string]string, 4)

func init() {
	// Load the Template into memory
	_, err := templates.ParseFiles("templates\\display.html")
	if err != nil {
		log.Fatalln(err)
	}
	// Create Hashes for API Keys
	fmt.Println("\n\n ... Generating API Keys for groups ...\n")
	for i := 1; i < (MaxGroups + 1); i++ {
		h := sha1.New()
		s := fmt.Sprintf("Group %d", i)
		io.WriteString(h, s)
		c := fmt.Sprintf("%x", string(h.Sum(nil)))
		hash_arr[c] = s
		fmt.Printf(" %s = %s\n", s, c)
	}
	fmt.Println("\n VV All Done ! VV \n\n")
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	// Get the URI
	loc := html.EscapeString(r.URL.Path)
	if loc == "/view" {
		http.Error(w, " invlid request", http.StatusUnauthorized)
		return
	}
	// Strip the Prefix
	loc = strings.TrimLeft(loc, "/view/")
	// Reject the No Names and multiple slash inputs
	if len(loc) == 0 || strings.Index(loc, "/") != -1 {
		log.Println("view: Error - Incorrect ID = ", loc)
		http.Error(w, " invlid request", http.StatusNotAcceptable)
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
		log.Println("view: Error -", err)
		http.Error(w, "Only POST request Supported", http.StatusInternalServerError)
	}
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {

	// Accept no method other than Post
	if r.Method != http.MethodPost {
		log.Println("update: Non POST request : ", r.Method)
		http.Error(w, "Only POST request Supported", http.StatusMethodNotAllowed)
		return
	}

	// Check if Key was supplied
	key := string(r.FormValue("api_key"))
	log.Println("update: received key ", key)
	if len(key) != MaxApiKey {
		log.Println("update: No API KEY suppiled ", key)
		http.Error(w, "No API KEY supplied", http.StatusBadRequest)
		return
	}

	// Check for the Presence of the Keys
	group, prs := hash_arr[key]
	if !prs {
		log.Println("update: Invalid API KEY suppiled ", key)
		http.Error(w, "Invalid API KEY supplied", http.StatusBadRequest)
		return
	}

	// Check if Data Type was supplied
	dtype := string(r.FormValue("data_type"))
	log.Println("update: received Data Type as ", dtype)
	if len(dtype) == 0 {
		log.Println("update: No Data Type suppiled ", dtype)
		http.Error(w, "No Data Type supplied", http.StatusBadRequest)
		return
	}

	// Check the Type
	_, prs = DataSampleTypes[dtype]
	if !prs {
		log.Println("update: Invalid Data Type suppiled ", dtype)
		http.Error(w, "Invalid Data Type supplied", http.StatusBadRequest)
		return
	}

	// Log the Received Key Request
	log.Println("update: Name =", group, " Type =", dtype)
	fmt.Fprintf(w, "updated %s", group)
}

func main() {
	http.HandleFunc("/view/", ViewHandler)
	http.HandleFunc("/update", UpdateHandler)
	http.HandleFunc("/favicon.ico", http.NotFound)
	http.Handle("/", http.FileServer(http.Dir("www")))
	log.Println("Starting Server at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
