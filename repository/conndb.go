package repository

import (
	"flag"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"log"
)

// initDb is database initialization
func initDb() *gorm.DB {
	var dbPath string
	if flag.Lookup("test.v") == nil {
		// normal run
		dbPath = "database/otoklix.db"
	} else {
		// run under go test
		dbPath = "../database/otoklix.db"
	}
	// Connect to database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	checkErr(err)

	return db
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
