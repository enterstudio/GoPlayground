package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// Table Name
	tableName = "customer_trial"
)

var (
	// Fields in the Table
	tableFields     = [...]string{"cID", "cName", "cPoints"}
	tableFieldsForm = [...]string{"cid", "cname", "cpoints"}
)

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

func connectDatabase() error {
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
		return err
	}
	err = db.Ping()
	check(err)
	if err == nil {
		log.Println("Database connection Successful")
	}
	return err
}

func disconnectDatabase() {
	err := db.Close()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Database connection stopped Successfully")
	}
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}

func db_formFieldMap(r *http.Request) map[string]formField {

	fields := make(map[string]formField, 3)

	for i, _ := range tableFields {
		fields[tableFields[i]] = checkField(tableFieldsForm[i], r)
	}

	return fields
}

func db_searchStmt(fields map[string]formField) string {

	stm := `SELECT * FROM ` + tableName

	if fields["cID"].Present {
		stm += " WHERE cID=? "
	}

	if fields["cName"].Present {
		if !fields["cID"].Present {
			stm += " WHERE cName=?"
		} else {
			stm += " AND cName=?"
		}
	}

	if fields["cPoints"].Present {
		if !fields["cID"].Present && !fields["cName"].Present {
			stm += " WHERE cPoints=?"
		} else {
			stm += " AND cPoints=?"
		}
	}

	return stm
}

func db_insertrStmt(fields map[string]formField) string {

	stm := `INSERT INTO ` + tableName

	if fields["cName"].Present && fields["cPoints"].Present {
		stm += " (cName, cPoints) VALUES (?, ?)"
	} else {
		return ""
	}

	return stm
}

func db_exeQuery(stm string, params ...interface{}) ([][]string, error) {

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
	if err == nil {
		log.Printf(" [db] Query Successful - %s", stm)
	}
	return arr, nil
}

func db_exeCmd(stm string, params ...interface{}) (int64, error) {
	stmt, err := db.Prepare(stm)
	check(err)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(params...)
	check(err)
	if err != nil {
		return 0, err
	}

	n, err := result.RowsAffected()
	check(err)

	if err == nil {
		log.Printf(" [db] Statement Executed (for %d) - %s", n, stm)
	}
	return n, err
}

func dbCreateTable() error {
	_, err := db_exeCmd(`CREATE TABLE ` + tableName + `(
		cID int NOT NULL AUTO_INCREMENT PRIMARY KEY,
		cName text NOT NULL,
		cPoints int NOT NULL
	)`)
	if err == nil {
		log.Println(" [db] Table Created Successfully")
	}
	return err
}

func dbDropTable() error {
	_, err := db_exeCmd(`DROP TABLE ` + tableName)
	if err == nil {
		log.Println(" [db] Removed Table Successfully")
	}
	return err
}

func dbSearch(r *http.Request) (pageData, error) {
	var pdData pageData
	pdData.Finfo = db_formFieldMap(r)
	stm := db_searchStmt(pdData.Finfo)
	fields := pdData.Finfo

	var err error

	if fields["cID"].Present && fields["cName"].Present && fields["cPoints"].Present {
		pdData.Recs, err = db_exeQuery(stm, fields["cID"].Value, fields["cName"].Value, fields["cPoints"].Value)
	} else if fields["cID"].Present && fields["cName"].Present {
		pdData.Recs, err = db_exeQuery(stm, fields["cID"].Value, fields["cName"].Value)
	} else if fields["cName"].Present && fields["cPoints"].Present {
		pdData.Recs, err = db_exeQuery(stm, fields["cName"].Value, fields["cPoints"].Value)
	} else if fields["cID"].Present && fields["cPoints"].Present {
		pdData.Recs, err = db_exeQuery(stm, fields["cID"].Value, fields["cPoints"].Value)
	} else if fields["cID"].Present {
		pdData.Recs, err = db_exeQuery(stm, fields["cID"].Value)
	} else if fields["cName"].Present {
		pdData.Recs, err = db_exeQuery(stm, fields["cName"].Value)
	} else if fields["cPoints"].Present {
		pdData.Recs, err = db_exeQuery(stm, fields["cPoints"].Value)
	} else {
		pdData.Recs, err = db_exeQuery(stm)
	}

	return pdData, err
}

func dbReadAll(r *http.Request) ([][]string, error) {
	fields := db_formFieldMap(r)
	stm := db_searchStmt(fields)

	var arr [][]string
	var err error
	arr, err = db_exeQuery(stm)
	if err == nil {
		log.Println(" [db] Reading Full Table Successful")
	}
	return arr, err
}

func dbAddRecord(r *http.Request) error {
	fields := db_formFieldMap(r)
	stm := db_insertrStmt(fields)

	if len(stm) == 0 {
		return errors.New(" Missing Parameters")
	}

	n, err := db_exeCmd(stm, fields["cName"].Value, fields["cPoints"].Value)
	if n == 1 && err == nil {
		log.Println(" [db] Added Record Sucessfully")
	}

	return err
}
