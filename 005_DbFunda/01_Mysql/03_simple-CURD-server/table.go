package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Table Name
const tableName = "customer_trial"

// Data Record
type record struct {
	cID     int
	cName   string
	cPoints int
}

func (p record) get() []string {
	m := make([]string, 0, 3)
	m = append(m, fmt.Sprintf("%d", p.cID))
	m = append(m, p.cName)
	m = append(m, fmt.Sprintf("%d", p.cPoints))
	return m
}

func (p record) fields() []string {
	m := make([]string, 0, 3)
	m = append(m, "cID")
	m = append(m, "cName")
	m = append(m, "cPoints")
	return m
}

func connectDatabase() {
	var err error
	connstr := os.Getenv("DB_CONNECTION")
	if len(connstr) == 0 {
		log.Println(`Wrong Environment
 
Need to Setup Environment Variable 'DB_CONNECTION'

  In Windows:
	set DB_CONNECTION="<User>:<Password>@tcp(<IP Address>:3306)/<DB Name>"

  In Linux:
	export DB_CONNECTION="<User>:<Password>@tcp(<IP Address>:3306)/<DB Name>"
	
	`)
		log.Fatalln(" Can't continue without these")
	}
	db, err = sql.Open("mysql", connstr)
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	check(err)
	log.Println("Database connection Successful")
}

func disconnectDatabase() {
	err := db.Close()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Database connection stopped Successfully")
	}
}

func dbTableCreate() error {
	stmt, err := db.Prepare(`CREATE TABLE ` + tableName + `(
		cID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
		cName text NOT NULL,
		cPoints int NOT NULL
	)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println(" [db] Table Created Successfully")
	return nil
}

func dbTableDrop() error {
	stmt, err := db.Prepare(`DROP TABLE ` + tableName)
	check(err)

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	log.Println(" [db] Removed Table Successfully")
	return nil
}

func dbTableReadAll() ([][]string, error) {
	stmt, err := db.Prepare(`SELECT * FROM ` + tableName)
	check(err)
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	stmt.Close()
	arr := make([][]string, 0, 10)
	for rows.Next() {
		var r record
		err = rows.Scan(&r.cID, &r.cName, &r.cPoints)
		check(err)
		log.Printf(" [db] Record # %d\t%s\t%d", r.cID, r.cName, r.cPoints)
		arr = append(arr, r.get())
	}
	log.Println(" [db] Reading Full Table Successful")
	return arr, nil
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
