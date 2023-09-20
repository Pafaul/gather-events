package db

import (
	"log"
	"os"
)

func OpenDBConnection() {
	pgConnStr, pgProvided := os.LookupEnv("PG_URL")
	if pgProvided {
		openPGConnection(pgConnStr)
		return
	}

	sqlite, sqliteProvided := os.LookupEnv("SQLITE_URL")
	if sqliteProvided {
		openSQLiteConnection(sqlite)
		return
	}

	log.Fatalln("No database url provided (SQLITE_URL and PG_URL envs are empty)")
}
