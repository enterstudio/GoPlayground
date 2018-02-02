package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	// Open File
	fl, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer fl.Close()

	// Load Templates
	tpl, err := template.ParseGlob("*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Read the Whole File
	rdata, err := ioutil.ReadAll(fl)
	if err != nil {
		log.Fatalln(err)
	}
	fl.Close()

	// Create the CSV Reader
	csvReader := csv.NewReader(strings.NewReader(string(rdata)))
	rawRecords, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	var processedRecords []map[string]string

	// Get the Header used in the File
	headers := rawRecords[0]
	log.Println("Headers: \n", headers)

	// Convert the Whole Dataset into Records
	for i := 1; i < len(rawRecords); i++ {
		mp := make(map[string]string, len(headers))
		for j := 0; j < len(headers); j++ {
			mp[headers[j]] = rawRecords[i][j]
		}
		processedRecords = append(processedRecords, mp)
	}

	log.Println("Full Data:\n", processedRecords)

	passingData := struct {
		Headers []string
		Data    []map[string]string
	}{
		headers,
		processedRecords,
	}
	tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", passingData)
}
