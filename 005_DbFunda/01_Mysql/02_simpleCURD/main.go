package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Database Object
var db *sql.DB

const tableName = "customer_trial"

var dataInsert = []struct {
	Name   string
	Points int
}{{"Ram", 100}, {"Dev", 40}, {"Pavan", 5}, {"Amit", 20}}

func init() {
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
	critialerror(err)
	err = db.Ping()
	critialerror(err)
	log.Println("Database connection Successful")
}

func main() {

	// Remember to Close the Db once we are done
	defer db.Close()

	// Create Table
	createTable()

	// Add Record to Table
	addRecords()

	// Read Record from Table
	readAll()

	// Find Record and get cID
	name := "Dev"
	cid, _, points := findRecord(name)
	log.Printf(" Found Customer cID : %d for Name = %s", cid, name)

	// Modify Record in Table
	modifyRecordByID(cid, "Dev Kiran", points)
	readAll()

	// Delete Record in Table
	name = "Pavan"
	cid, _, _ = findRecord(name)
	log.Printf(" Found Customer cID : %d for Name = %s", cid, name)
	deleteRecordByID(cid)
	readAll()

	// Drop Table
	dropTable()
}

func critialerror(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func createTable() {
	stmt, err := db.Prepare(`CREATE TABLE ` + tableName + `(
		cID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
		cName text NOT NULL,
		cPoints int NOT NULL
	)`)
	critialerror(err)
	defer stmt.Close()
	_, err = stmt.Exec()
	critialerror(err)

	log.Println(" Table Created Successfully")
}

func addRecords() {
	stmt, err := db.Prepare(`INSERT INTO ` + tableName +
		` (cName, cPoints) VALUES (?, ?)`)
	check(err)
	defer stmt.Close()

	for _, d := range dataInsert {
		_, err = stmt.Exec(d.Name, d.Points)
		check(err)
	}

	log.Println(" Values Inserted Successfully")
}

func readAll() {
	stmt, err := db.Prepare(`SELECT * FROM ` + tableName)
	check(err)
	rows, err := stmt.Query()
	critialerror(err)
	stmt.Close()
	for rows.Next() {
		var cID int
		var cName string
		var cPoints int

		err = rows.Scan(&cID, &cName, &cPoints)
		check(err)
		log.Printf(" Record # %d\t%s\t%d", cID, cName, cPoints)
	}
	log.Println(" Reading Full Table Successful")
}

func dropTable() {
	stmt, err := db.Prepare(`DROP TABLE ` + tableName)
	check(err)
	_, err = stmt.Exec()
	critialerror(err)
	log.Println("< Removed Table Successfully")
}

func findRecord(name string) (cid int, cname string, cpoints int) {
	stmt, err := db.Prepare(`SELECT * FROM ` + tableName +
		` WHERE cName=?`)
	check(err)
	defer stmt.Close()

	err = stmt.QueryRow(name).Scan(&cid, &cname, &cpoints)
	critialerror(err)
	return
}

func modifyRecordByID(cid int, name string, points int) {
	/*
	   UPDATE `customer_trial` SET
	   `cID` = '2',
	   `cName` = 'Dev Kiran',
	   `cPoints` = '40'
	   WHERE `cID` = '2';
	*/
	stmt, err := db.Prepare(`UPDATE ` + tableName +
		` SET cID=?,cName=?,cPoints=? WHERE cID=?`)
	check(err)
	defer stmt.Close()
	result, err := stmt.Exec(cid, name, points, cid)
	critialerror(err)
	n, err := result.RowsAffected()
	check(err)
	log.Printf(" Records Update %d", n)
}

func deleteRecordByID(cid int) {
	/*
		DELETE FROM `customer_trial`
		WHERE ((`cID` = '3'));
	*/
	stmt, err := db.Prepare(`DELETE FROM ` + tableName +
		` WHERE cID=?`)
	check(err)
	defer stmt.Close()
	result, err := stmt.Exec(cid)
	critialerror(err)
	n, err := result.RowsAffected()
	check(err)
	log.Printf(" Records Deleted %d", n)
}
