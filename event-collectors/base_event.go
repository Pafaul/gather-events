package eventcollectors

import (
	"gather-events/db"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

var (
	insertBaseEvent = `
        INSERT INTO base_event (unique_hash, contract_id, block_number) 
        VALUES ($1, $2, $3)
    `
)

func logBaseEvent(contractId uint64, event types.Log) string {
	hash := getUniqueEventHash(event).Hex()
	if db.DBConn == nil {
		log.Println(logToString(contractId, event))
		return hash
	}

	_, err := db.DBConn.Exec(
		insertBaseEvent,
		hash,
		contractId,
		event.BlockNumber,
	)

	if !db.IsUniqueViolation(err) {
		failOnUnknownError(err)
	}

	return hash
}
