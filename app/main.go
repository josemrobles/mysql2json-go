package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	// Flags
	host := flag.String("host", "", "MySQL server host")
	port := flag.String("port", "", "server port")
	database := flag.String("database", "", "default database")
	query := flag.String("query", "", "a SQL query")
	//query_file := flag.String("query_file", "", "a SQL query file")
	user := flag.String("user", "", "username")
	password := flag.String("password", "", "password")
	flag.Parse()

	// Connect to mysql database
	db, err := sql.Open("mysql", *user+":"+*password+"@tcp("+*host+":"+*port+")/"+*database+"?allowOldPasswords=1")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	// Query database
	jsonData, err := getData(db, *query)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(jsonData)
	}
}

func getData(db *sql.DB, query string) (string, error) {

	// Execute the query
	records, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer records.Close()

	// Get the column names
	cols, err := records.Columns()
	if err != nil {
		return "", err
	}
	count := len(cols)
	td := make([]map[string]interface{}, 0)
	data := make([]interface{}, count)
	vals := make([]interface{}, count)
	
	for records.Next() {
		for i := 0; i < count; i++ {
			vals[i] = &data[i]
		}
		records.Scan(vals...)
		entry := make(map[string]interface{})
		for i, column := range columns {
			var v interface{}
			datum := data[i]
			b, ok := datum.([]byte)
			if ok {
				v = string(b)
			} else {
				v = datum
			}
			entry[column] = v
		}
		td = append(td, entry)
	}
	json, err := json.Marshal(td)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
