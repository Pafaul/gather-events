package db

import (
	"log"
	"os"

	"database/sql"
)

type (
	dbQueryWithMsg struct {
		creationQuery string
		msg           string
	}
)

var (
	DBConn *sql.DB
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

func createDBS(db *sql.DB, databases []dbQueryWithMsg) {
	for _, dbToCreate := range databases {
		log.Println("Creating", dbToCreate.msg, "table")
		_, createErr := db.Exec(dbToCreate.creationQuery)
		if createErr != nil {
			log.Fatalln(createErr)
		}
	}
}
