package main

import (
	"net/http"
)

type formField struct {
	Present bool
	Value   string
}

func checkField(key string, r *http.Request) formField {
	var f formField
	f.Value = r.FormValue(key)
	if len(f.Value) != 0 {
		f.Present = true
	} else {
		f.Present = false
	}
	return f
}

type pageDataMin struct {
	TableCreated bool
	Title        string
	Heading      string
}

func newPageDataMin(title, heading string) pageDataMin {
	return pageDataMin{tableCreated, title, heading}
}

type pageData struct {
	pageDataMin
	Recs  [][]string
	Finfo map[string]formField
}

func updatePageData(p *pageData, title, heading string) {
	p.pageDataMin = newPageDataMin(title, heading)
}

func root(w http.ResponseWriter, r *http.Request) {
	pgData := newPageDataMin("Home Page", "Welcome to the CURD Server")
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
		pgData := newPageDataMin("Table Creation", "Table Successfully Created !")
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
	var pgData pageData
	var err error
	// Actually Read all data
	pgData.Recs, err = dbReadAll(r)
	if err == nil {
		updatePageData(&pgData, "Reading All Records", "All Records")
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

	var err error
	pgData, err := dbSearch(r)

	if err == nil {
		updatePageData(&pgData, "Find Records", "Finding Records of Interest")
		tmpl.ExecuteTemplate(w, "find.gohtml", pgData)
	} else {
		check(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	pgData := newPageDataMin("Record Update", "Updating Records")
	tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	pgData := newPageDataMin("Record Deletion", "Delete Record")
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
		pgData := newPageDataMin("Remove Table", "Table Deleted Succesfully !")
		tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
