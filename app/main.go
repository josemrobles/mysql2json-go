package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	host := flag.String("host", "", "MySQL server host")
	port := flag.String("port", "", "server port")
	datbase := flag.String("database", "", "default database")
	query := flag.String("query", "", "a SQL query")
	query_file := flag.String("query_file", "", "a SQL query file")
	user := flag.String("user", "", "username")
	password := flag.String("password", "", "password")
	flag.Parse()

	con, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?allowOldPasswords=1")
	if err != nil {
		panic(err)
	}
	defer con.Close()
}
