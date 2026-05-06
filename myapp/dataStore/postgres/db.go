package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// db details
const (
	postgres_host     = "localhost"
	postgres_port     = 5432
	postgres_user     = "postgres"
	postgres_password = "postgres"
	postgres_dbname   = "postgres"
)

// create pointer variable Db which points to sql driver
var Db *sql.DB

// init() is always called before main() by Go compiler
func init() {
	// creating a database connection string
	db_info := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgres_host,
		postgres_port,
		postgres_user,
		postgres_password,
		postgres_dbname,
	)

	// fmt.Println(db_info)
	var err error

	// establish a connection to postgresql server using the driver
	Db, err = sql.Open("postgres", db_info)

	// handle error
	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully connected")
	}
}
