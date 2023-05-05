package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func createDB() *sql.DB {
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "_ABcd149126",
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    "pets_manage",
		ParseTime: true,
		Loc: time.Local}
	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Capture connection properties.
