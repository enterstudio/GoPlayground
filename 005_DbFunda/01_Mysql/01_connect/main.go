package main

import (
	"database/sql"
	"log"
	_ "net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

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

	defer db.Close()

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
