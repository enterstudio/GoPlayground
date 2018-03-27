package main

import "net/http"

type pageData struct {
	TableCreated bool
	Title        string
	Heading      string
	Rec          []string
	Recs         [][]string
}

func root(w http.ResponseWriter, r *http.Request) {
	pgData := pageData{tableCreated,
		"Home Page", "Welcome to the CURD Server", nil, nil}
	tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
}

func existingTable(w http.ResponseWriter, r *http.Request) bool {
	if !tableCreated {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return false
	}
	return tableCreated
}

func createTable(w http.ResponseWriter, r *http.Request) {

	if tableCreated {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	// Call the Creation function
	err := dbTableCreate()
	if err == nil {
		tableCreated = true
		pgData := pageData{tableCreated,
			"Table Creation", "Table Successfully Created !", nil, nil}
		tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
	} else {
		check(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func readAll(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	pgData := pageData{tableCreated,
		"Reading All Records", "All Records", nil, nil}
	var err error
	// Actually Read all data
	pgData.Recs, err = dbTableReadAll()
	if err == nil {
		pgData.Rec = record{}.fields()
		tmpl.ExecuteTemplate(w, "readall.gohtml", pgData)
	} else {
		check(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func findRecord(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	pgData := pageData{tableCreated,
		"Find Records", "Finding Records of Interest", nil, nil}
	tmpl.ExecuteTemplate(w, "find.gohtml", pgData)
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	pgData := pageData{tableCreated,
		"Record Update", "Updating Records", nil, nil}
	tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	pgData := pageData{tableCreated,
		"Record Deletion", "Delete Record", nil, nil}
	tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
}

func dropTable(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}

	// Call Drop Table
	err := dbTableDrop()
	if err == nil {
		tableCreated = false
		pgData := pageData{tableCreated,
			"Remove Table", "Table Deleted Succesfully !", nil, nil}
		tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
