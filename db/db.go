package db

import (
	"database/sql"
	"fmt"
	// Dont need the functions really so jus tgoing through it
	_ "github.com/mattn/go-sqlite3"
)

// Database which is to be used in the functions file
var Database *sql.DB
var err error

//The init
func init(){
	Database, err = sql.Open("sqlite3", "./tasks.db")

	if err != nil {
		fmt.Println(err)
	}
}