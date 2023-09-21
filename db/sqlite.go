package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	sqliteDBS = []dbQueryWithMsg{
		{sqliteContractInfoTable, "base contract info"},
		{sqliteBaseEventTable, "base event"},
		{sqliteERC20TransfersTable, "erc20 transfer"},
		{sqliteERC20BalanceChangesTable, "erc20 user balance"},
		{sqliteBridgeClaimsTable, "claims"},
		{sqliteBridgeTeleportsTable, "bridge teleleports"},
		{sqliteBridgeRefundStatusTable, "bridge refund status"},
	}
)

func openSQLiteConnection(connectionString string) *sql.DB {
	var err error
	DBConn, err = sql.Open("sqlite3", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	err = DBConn.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	createDBS(DBConn, sqliteDBS)

	return DBConn
}
