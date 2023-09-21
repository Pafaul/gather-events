package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	pgDBS = []dbQueryWithMsg{
		{pgContractInfoTable, "base contract info"},
		{pgBaseEventTable, "base event"},
		{pgERC20TransfersTable, "erc20 transfer"},
		{pgERC20BalanceChangesTable, "erc20 user balance"},
		{pgBridgeClaimsTable, "claims"},
		{pgBridgeTeleportsTable, "bridge teleleports"},
		{pgBridgeRefundStatusTable, "bridge refund status"},
	}
)

func openPGConnection(connectionString string) *sql.DB {
	var err error
	DBConn, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	err = DBConn.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	createDBS(DBConn, pgDBS)

	return DBConn
}
