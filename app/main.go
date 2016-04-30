package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	host := flag.String("host", "", "MySQL server host")
	port := flag.String("port", "", "server port")
	database := flag.String("database", "", "default database")
	query := flag.String("query", "", "a SQL query")
	query_file := flag.String("query_file", "", "a SQL query file")
	user := flag.String("user", "", "username")
	password := flag.String("password", "", "password")
	flag.Parse()

	c, err := sql.Open("mysql", *user+":"+*password+"@tcp("+*host+":"+*port+")/"+*database+"?allowOldPasswords=1")
	if err != nil {
		panic(err)
	}
	defer c.Close()
}

func getDatai(c *sql.DB, query string) (error, data string) {

}
