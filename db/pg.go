package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var (
	createContractInfoTable = `
        CREATE TABLE IF NOT EXISTS contracts (
            contract_id serial PRIMARY KEY,
            contract_address VARCHAR (42) NOT NULL,
            chain_id INT NOT NULL,
            last_known_block BIGINT NOT NULL,
            UNIQUE (contract_address, chain_id)
        )
    `

	createBaseEventTable = `
        CREATE TABLE IF NOT EXISTS base_event (
            unique_hash VARCHAR (66) PRIMARY KEY,
            contract_id INT NOT NULL,
            block_number BIGINT NOT NULL,
            CONSTRAINT base_event_contract
                FOREIGN KEY (contract_id)
                REFERENCES contracts (contract_id)
        )
    `

	createERC20TransfersTable = `
        CREATE TABLE IF NOT EXISTS erc20_transfers (
            transfer_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            CONSTRAINT base_event_erc20_transfer
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	createBridgeTeleportsTable = `
        CREATE TABLE IF NOT EXISTS bridge_teleports (
            teleport_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            nonce NUMERIC (80) NOT NULL,
            refund_status INT NOT NULL,
            CONSTRAINT base_event_bridge_telepots
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	createBridgeRefundStatusTable = `
        CREATE TABLE IF NOT EXISTS bridge_teleport_refund_status (
            refund_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            teleport_id INT NOT NULL,
            status smallint NOT NULL,
            timestamp bigint NOT NULL,
            CONSTRAINT base_event_refund_status
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            CONSTRAINT refund_status_teleport
                FOREIGN KEY (teleport_id)
                REFERENCES bridge_teleports (teleport_id),
            UNIQUE (base_event_id)
        )
    `

	createBridgeClaimsTable = `
        CREATE TABLE IF NOT EXISTS bridge_claims (
            claim_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            nonce NUMERIC (80) NOT NULL,
            CONSTRAINT base_event_bridge_claims
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	createERC20BalanceChangesTable = `
        CREATE TABLE IF NOT EXISTS user_erc20_balance (
            base_event_id VARCHAR (66) NOT NULL, 
            user_address VARCHAR (42) NOT NULL,
            balance_change NUMERIC (80) NOT NULL,
            PRIMARY KEY (base_event_id, user_address),
            UNIQUE (base_event_id, user_address)
        )
    `

	createERC721Transfers = `
        CREATE TABLE IF NOT EXISTS erc721_transfers (
            transfer_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            token_id NUMERIC (80) NOT NULL,
            CONSTRAINT base_event_erc721_transfer
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	createERC721Balances = `
        CREATE TABLE IF NOT EXISTS erc721_balances (
            contract_id INT NOT NULL, 
            token_id NUMERIC (80) NOT NULL,
            user_address VARCHAR (42) NOT NULL,
            CONSTRAINT erc721_balances_contract_id
                FOREIGN KEY (contract_id)
                REFERENCES contracts (contract_id),
            PRIMARY KEY (contract_id, token_id)
        )
    `

	DBConn *sql.DB
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

	dbs := []struct {
		creationQuery string
		msg           string
	}{
		{createContractInfoTable, "base contract info"},
		{createBaseEventTable, "base event"},
		{createERC20TransfersTable, "erc20 transfer"},
		{createERC20BalanceChangesTable, "erc20 user balance"},
		{createBridgeClaimsTable, "claims"},
		{createBridgeTeleportsTable, "bridge teleleports"},
		{createBridgeRefundStatusTable, "bridge refund status"},
	}

	for _, dbToCreate := range dbs {
		log.Println("Creating", dbToCreate.msg, "table")
		_, createErr := DBConn.Exec(createContractInfoTable)
		if createErr != nil {
			log.Fatalln(createErr)
		}
	}

	return DBConn
}

func IsUniqueViolation(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		return err.Code == "23505"
	}

	return false
}

func IsNoRowsErr(err error) bool {
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("Error is pg no rows")
		return err.Code == "02000"
	}

	return err == sql.ErrNoRows
}
