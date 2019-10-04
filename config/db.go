package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// TODO: Following variables should be moved to env variables or DB

const db_host = "moiadb.cm7fu3cututl.eu-central-1.rds.amazonaws.com"
const db_user = "dbuser"
const db_password = "dbpassword"
const db_name = "moiadb"
const db_port = "5432"

var DB *sql.DB

// initializes the database connection
// Then checks if we have urls tables, if not, creates it

func init() {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", db_host, db_port, db_user, db_name, db_password))
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database.")

	_, err = DB.Exec("SELECT 1 FROM urls LIMIT 1;")
	if err != nil {
		// Table does not exist so create it
		_, err = DB.Exec("create table urls " +
			"( id serial primary key not null," +
			"key varchar(6) unique not null," +
			"url text not null);")
		if err != nil {
			panic(err)
		}
	}
}
