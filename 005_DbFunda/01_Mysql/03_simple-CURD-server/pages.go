package main

import (
	"fmt"
	"net/http"
	"strings"
)

type formField struct {
	Present bool
	Value   string
}

func (f formField) UpdatePresent(p bool) formField {
	f.Present = p
	return f
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
	err := dbCreateTable()
	if err == nil {
		tableCreated = true
		pgData := newPageDataMin("Table Creation", "Table Successfully Created !")
		tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
	} else if strings.Contains(err.Error(), "Error 1050") {
		tableCreated = true
		pgData := newPageDataMin("Table Creation", "Table Already Exists !")
		tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
	} else {
		check(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

func add(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	pgData := newPageDataMin("Add Records", "New Record")

	err := dbAddRecord(r)
	if err == nil {
		pgData.Heading = " Record Created Successfully"
	}
	tmpl.ExecuteTemplate(w, "add.gohtml", pgData)
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
	var pgData pageData
	wasFind := r.FormValue("submit") == "Find it"

	if wasFind {
		pgData, err = dbSearch(r)
	}
	if err != nil {
		check(err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
	updatePageData(&pgData, "Find Records", "Finding Records of Interest")

	tmpl.ExecuteTemplate(w, "find.gohtml", pgData)
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	var err error
	var pgData pageData
	wasGet := r.FormValue("submit") == "Get"
	//wasUpdate := r.FormValue("submit") == "Update"
	// Execute the Update
	n, pgData, err := dbUpdateRecord(r)
	updatePageData(&pgData, "Record Update", "Updating Records")

	if err != nil {
		pgData.Heading += " -" + err.Error()
	} else {
		if n > 0 {
			if n > 1 {
				pgData.Heading += fmt.Sprintf("- Found %d records", n)
			} else {
				if wasGet {
					pgData.Heading += "- Found one"
				} else {
					pgData.Heading += "- Updated Record"
				}
			}
		} else {
			pgData.Heading = "Updating Records - Nothing Found"
		}
	}
	tmpl.ExecuteTemplate(w, "update.gohtml", pgData)
}

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}
	var err error
	var pgData pageData
	wasGet := r.FormValue("submit") == "Get"
	pgData, err = dbDeleteRecord(r)

	updatePageData(&pgData, "Record Deletion", "Delete Record")
	if err != nil {
		pgData.Heading += " -" + err.Error()
	} else {
		if wasGet {
			pgData.Heading += " - Found one"
		} else {
			pgData.Heading += " - Successful"
		}
	}
	tmpl.ExecuteTemplate(w, "delete.gohtml", pgData)
}

func dropTable(w http.ResponseWriter, r *http.Request) {
	if !existingTable(w, r) {
		return
	}

	// Call Drop Table
	err := dbDropTable()
	if err == nil {
		tableCreated = false
		pgData := newPageDataMin("Remove Table", "Table Deleted Succesfully !")
		tmpl.ExecuteTemplate(w, "index.gohtml", pgData)
	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
