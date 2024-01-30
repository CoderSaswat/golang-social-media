package config

import (
	"database/sql"
	"fmt"
	//_ "github.com/denisenkom/go-mssqldb"
	"log"
)

var SqlServerClient *sql.DB

func init() {
	server := "localhost"
	port := 1433
	user := "sa"
	password := "root"
	database := "social-media"

	// Build connection string
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)

	// Open a database connection
	db, err := sql.Open("mssql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	SqlServerClient = db
	//defer db.Close()

	// Try to ping the database to check if the connection is successful
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal(err)
	//}

}

func GetMsSqlClient() *sql.DB {
	return SqlServerClient
}
