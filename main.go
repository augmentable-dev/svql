package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/sqltocsv"
	"github.com/mattn/go-sqlite3"
)

var file string
var query string

func init() {
	flag.StringVar(&file, "file", "", "csv file")
	flag.StringVar(&query, "query", "", "query")
	flag.Parse()
}

func main() {
	sql.Register("sqlite3_with_extensions",
		&sqlite3.SQLiteDriver{
			Extensions: []string{
				"csv",
			},
		})
	db, err := sql.Open("sqlite3_with_extensions", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE VIRTUAL TABLE temp.csv USING csv(filename='%s', header=true);", file))
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	err = sqltocsv.Write(os.Stdout, rows)
	if err != nil {
		log.Fatal(err)
	}
}
