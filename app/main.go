package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	host := flag.String("host", "", "MySQL server host")
	port := flag.String("port", "", "server port")
	database := flag.String("database", "", "default database")
	query := flag.String("query", "", "a SQL query")
	//query_file := flag.String("query_file", "", "a SQL query file")
	user := flag.String("user", "", "username")
	password := flag.String("password", "", "password")
	flag.Parse()

	db, err := sql.Open("mysql", *user+":"+*password+"@tcp("+*host+":"+*port+")/"+*database+"?allowOldPasswords=1")
	if err != nil {
		log.Println("one")
		log.Fatal(err)
	}
	defer db.Close()

	getData(db, *query)
}

func getData(db *sql.DB, query string) {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		// do stuff with data
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
