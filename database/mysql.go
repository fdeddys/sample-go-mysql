package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB

func init() {
	var err error
	dbConn, err = openConnection()
	if err != nil {

		panic("Failed open to database" + err.Error())
	}
}

func openConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/gromart")

	// if there is an error opening the connection, handle it
	if err != nil {
		return nil, err
	}

	// defer the close till after the main function has finished
	// executing
	// defer db.Close()

	fmt.Println("Database connected!")
	return db, nil
}

func GetConnection() *sql.DB {
	return dbConn
}
