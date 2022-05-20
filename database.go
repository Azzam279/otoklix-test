package main

import (
	"github.com/jinzhu/gorm"

	"log"
)

// initDb is database initialization
func initDb() *gorm.DB {
	// Connect to database
	db, err := gorm.Open("sqlite3", "./otoklix.db")
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
