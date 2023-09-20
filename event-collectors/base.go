package eventcollectors

import (
	"context"
	"database/sql"
	"fmt"
	"gather-events/db"
	logsource "gather-events/log-source"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type (
	EventCollector interface {
		Address() common.Address
		ChainId() uint64
		LastKnownBlock() uint64

		LoadContractInfo() error
		SubscribeToLogs(logsource.LogBroadcaster)
		StartEventCollector(<-chan bool, chan<- error)
		SaveLastKnownBlock() error
	}

	eventCollector struct {
		logsReceiver   <-chan types.Log
		address        common.Address
		chainId        uint64
		lastKnownBlock uint64
		contractDBId   uint64
	}

	eventLoggerInterface interface {
		LogEvents(chainId uint64)
	}

	eventLogger struct {
		ctx        context.Context
		logChannel chan *types.Log
	}
)

var (
	ZERO_ADDRESS = "0x0000000000000000000000000000000000000000"

	getContractInfo = `
        SELECT contract_id, last_known_block
        FROM contracts
        WHERE 
            contract_address = $1 AND chain_id = $2;
    `

	saveLastKnownBlock = `
        INSERT INTO contracts (
            contract_address,
            chain_id,
            last_known_block
        )
        VALUES ($1, $2, $3)
		RETURNING contract_id;
    `

	updateLastKnownBlock = `
        UPDATE contracts 
        SET last_known_block = $1
        WHERE 
            contract_address = $2 AND chain_id = $3
    `
)

func (ec *eventCollector) Address() common.Address {
	return ec.address
}

func (ec *eventCollector) LastKnownBlock() uint64 {
	return ec.lastKnownBlock
}

func (ec *eventCollector) ChainId() uint64 {
	return ec.chainId
}

func (ec *eventCollector) SubscribeToLogs(ls logsource.LogBroadcaster) {
	ec.logsReceiver = ls.Subscribe()
}

func (ec *eventCollector) LoadContractInfo() error {
	if db.DBConn == nil {
		return nil
	}

	res := db.DBConn.QueryRow(getContractInfo, ec.address.Hex(), ec.chainId)
	if res.Err() != nil && res.Err() != sql.ErrNoRows {
		return res.Err()
	}

	contractInfoFromDB := new(struct {
		ContractId     int64
		LastKnownBlock int64
	})
	scanErr := res.Scan(&contractInfoFromDB.ContractId, &contractInfoFromDB.LastKnownBlock)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			id, err := ec.createDBRecord()
			if err != nil {
				return err
			}

			contractInfoFromDB.ContractId = int64(id)
			contractInfoFromDB.LastKnownBlock = int64(ec.lastKnownBlock)
		}
	}

	ec.contractDBId = uint64(contractInfoFromDB.ContractId)
	ec.lastKnownBlock = uint64(contractInfoFromDB.LastKnownBlock)

	return nil
}

func (ec *eventCollector) createDBRecord() (uint64, error) {
	contractId := uint64(0)
	err := db.DBConn.QueryRow(
		saveLastKnownBlock,
		ec.Address().Hex(),
		ec.ChainId(),
		ec.LastKnownBlock(),
	).Scan(&contractId)

	if checkUniqueViolation(err, "contract info") {
		return ec.contractDBId, nil
	}

	return contractId, err
}

func (ec *eventCollector) SaveLastKnownBlock() error {
	_, err := db.DBConn.Exec(
		updateLastKnownBlock,
		ec.LastKnownBlock(),
		ec.Address().Hex(),
		ec.ChainId(),
	)

	return err
}

func formatECError(eventName string, ec *eventCollector, log *types.Log, err error) error {
	return fmt.Errorf(
		"Error parsing %v event for contract %v\nTxHash: %v, LogIndex: %v\nError: %v",
		eventName,
		ec.address.Hex(),
		log.TxHash.Hex(),
		log.Index,
		err,
	)
}

func getUniqueEventHash(event types.Log) *common.Hash {
	keccak := crypto.NewKeccakState()
	keccak.Write(event.TxHash[:])
	keccak.Write([]byte(strconv.Itoa(int(event.Index))))
	uniqueHash := new(common.Hash)
	_, _ = keccak.Read(uniqueHash[:])

	return uniqueHash
}

func checkUniqueViolation(err error, dataName string) bool {
	if err != nil {
		if db.IsUniqueViolation(err) {
			log.Printf("Skipping duplicate %v data\n", dataName)
			return true
		}
		return false
	}

	return false
}

func failOnUnknownError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
