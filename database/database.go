package database

import (
	"database/sql"
	"fmt"
)

func Connect() (*sql.DB, error) {

	var username string = "root"
	var password string = ""
	var host string = "localhost"
	var database string = "laporin"

	db, err := sql.Open("mysql", fmt.Sprintf("%s@%stcp(%s:3306)/%s", username, password, host, database))

	if err != nil {
		return nil, err
	}

	return db, nil

}
