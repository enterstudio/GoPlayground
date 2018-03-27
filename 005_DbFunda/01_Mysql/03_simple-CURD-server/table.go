package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
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
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

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

func formFieldMap(r *http.Request) map[string]formField {

	fields := make(map[string]formField, 3)

	fields["cID"] = checkField("cid", r)
	fields["cName"] = checkField("cname", r)
	fields["cPoints"] = checkField("cpoints", r)

	return fields
}

func searchStmt(fields map[string]formField) string {

	stm := `SELECT * FROM ` + tableName

	if fields["cID"].present {
		stm += " WHERE cID=? "
	}

	if fields["cName"].present {
		if !fields["cID"].present {
			stm += " WHERE cName=?"
		} else {
			stm += " AND cName=?"
		}
	}

	if fields["cPoints"].present {
		if !fields["cID"].present && !fields["cName"].present {
			stm += " WHERE cPoints=?"
		} else {
			stm += " AND cPoints=?"
		}
	}

	return stm
}

func dbExeQuery(stm string, params ...interface{}) ([][]string, error) {

	stmt, err := db.Prepare(stm)
	check(err)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows *sql.Rows

	rows, err = stmt.Query(params...)
	check(err)
	if err != nil {
		return nil, err
	}

	arr := make([][]string, 0, 10)
	for rows.Next() {
		var r record
		err = rows.Scan(&r.cID, &r.cName, &r.cPoints)
		check(err)
		log.Printf(" [db] Record # %d\t%s\t%d", r.cID, r.cName, r.cPoints)
		arr = append(arr, r.get())
	}

	log.Println(" [db] Query Successful")

	return arr, nil
}

func dbsearchProcess(r *http.Request) ([][]string, error) {

	fields := formFieldMap(r)
	stm := searchStmt(fields)

	log.Println(" Executing statement -", stm)
	var arr [][]string
	var err error

	if fields["cID"].present && fields["cName"].present && fields["cPoints"].present {
		arr, err = dbExeQuery(stm, fields["cID"].value, fields["cName"].value, fields["cPoints"].value)
	} else if fields["cID"].present && fields["cName"].present {
		arr, err = dbExeQuery(stm, fields["cID"].value, fields["cName"].value)
	} else if fields["cName"].present && fields["cPoints"].present {
		arr, err = dbExeQuery(stm, fields["cName"].value, fields["cPoints"].value)
	} else if fields["cID"].present && fields["cPoints"].present {
		arr, err = dbExeQuery(stm, fields["cID"].value, fields["cPoints"].value)
	} else if fields["cID"].present {
		arr, err = dbExeQuery(stm, fields["cID"].value)
	} else if fields["cName"].present {
		arr, err = dbExeQuery(stm, fields["cName"].value)
	} else if fields["cPoints"].present {
		arr, err = dbExeQuery(stm, fields["cPoints"].value)
	} else {
		arr, err = dbExeQuery(stm)
	}

	return arr, err
}
