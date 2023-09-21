package eventcollectors

import (
	"gather-events/contracts"
	"gather-events/db"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type (
	ERC721EventCollector struct {
		eventCollector
	}
)

var (
	erc721Contract            *contracts.ERC721
	erc721TransferHash        = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	insertERC721TransferQuery = `
        INSERT INTO erc721_transfers (
            base_event_id,
            from_address,
            to_address,
            token_id
        )
        VALUES ($1, $2, $3, $4)
    `

	checkIfTokenExistsQuery = `
        SELECT COUNT (*) 
        FROM erc721_balances
        WHERE contract_id = $1 AND token_id = $2
    `

	updateERC721BalanceQuery = `
        UPDATE erc721_balances
        SET user_address = $1
        WHERE contract_id = $2 AND token_id = $3;
    `

	insertERC721BalanceQuery = `
        INSERT INTO erc721_balances (
            contract_id,
            token_id,
            user_address
        ) 
        VALUES ($1, $2, $3)
    `
)

func (erc721 *ERC721EventCollector) StartEventCollector(stopChannel <-chan bool, errorChannel chan<- error) {
	for {
		select {
		case <-stopChannel:
			return
		case txLog := <-erc721.logsReceiver:
			if txLog.Address.Hex() == erc721.Address().Hex() {
				var logErr error

				switch txLog.Topics[0].Hex() {
				case erc721TransferHash.Hex():
					transfer, err := erc721Contract.ParseTransfer(txLog)
					if err != nil {
						errorChannel <- formatECError("ERC721Transfer", &erc721.eventCollector, &txLog, err)
					}
					logErr = logERC721Transfer(erc721.contractDBId, transfer)
				}

				if logErr != nil {
					errorChannel <- logErr
				}

				erc721.lastKnownBlock = txLog.BlockNumber
			}
		}
	}
}

func logERC721Transfer(contractId uint64, erc721Transfer *contracts.ERC721Transfer) error {
	if db.DBConn == nil {
		log.Print(erc721TransferToString(erc721Transfer))
	}

	baseEventId := logBaseEvent(contractId, erc721Transfer.Raw)

	_, err := db.DBConn.Exec(
		insertERC721TransferQuery,
		baseEventId,
		erc721Transfer.From.Hex(),
		erc721Transfer.To.Hex(),
		erc721Transfer.TokenId.String(),
	)

	if !checkUniqueViolation(err, "ERC721 transfer") {
		return err
	}

	return updateERC721Balance(contractId, erc721Transfer)
}

func updateERC721Balance(contractId uint64, erc721Transfer *contracts.ERC721Transfer) error {
	res := db.DBConn.QueryRow(
		checkIfTokenExistsQuery,
		contractId,
		erc721Transfer.TokenId.String(),
	)

	if res.Err() != nil {
		return res.Err()
	}
	var count int64
	scanErr := res.Scan(&count)
	if scanErr != nil {
		return scanErr
	}

	if count == 0 {
		_, err := db.DBConn.Exec(
			insertERC721BalanceQuery,
			erc721Transfer.TokenId.String(),
			erc721Transfer.To.String(),
		)
		if !checkUniqueViolation(err, "ERC721 balances") {
			return err
		}

		return nil
	}

	_, erc721BalanceUpdateErr := db.DBConn.Exec(
		updateERC721BalanceQuery,
		erc721Transfer.To.Hex(),
		contractId,
		erc721Transfer.TokenId.String(),
	)

	return erc721BalanceUpdateErr
}

func NewERC721EventCollector(chainId uint64, address common.Address, lastKnownBlock uint64) EventCollector {
	if erc721Contract == nil {
		erc721Contract, _ = contracts.NewERC721(address, nil)
	}

	erc721EC := &ERC721EventCollector{
		eventCollector: eventCollector{
			chainId:        chainId,
			address:        address,
			lastKnownBlock: lastKnownBlock,
		},
	}

	return erc721EC
}

var _ (EventCollector) = (*ERC721EventCollector)(nil)
