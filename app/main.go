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

	jsonData, err := getData(db, *query)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(jsonData)
	}
}

func getData(db *sql.DB, query string) (string, error) {

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Get the column names
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)

	// Get the data and convert to json
	tableData := make([]map[string]interface{}, 0)
	data := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &data[i]
		}
		rows.Scan(valuePtrs...)
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
		tableData = append(tableData, entry)
	}
	json, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(json), nil
}
