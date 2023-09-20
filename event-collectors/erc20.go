package eventcollectors

import (
	"fmt"
	"gather-events/contracts"
	"gather-events/db"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type (
	ERC20EventCollector struct {
		eventCollector
	}
)

var (
	erc20Contract            *contracts.ERC20
	transferERC20Hash        = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	insertTransferERC20Query = `
        INSERT INTO erc20_transfers (
            base_event_id, 
            from_address, 
            to_address, 
            amount
        ) 
        VALUES ($1, $2, $3, $4)
    `
	addUserBalanceChange = `
        INSERT INTO user_erc20_balance (
            base_event_id, 
            user_address, 
            balance_change
        )
        VALUES ($1, $2, $3) 
    `
)

func (erc20 *ERC20EventCollector) StartEventCollector(stopChannel <-chan bool, errorChannel chan<- error) {
	for {
		select {
		case <-stopChannel:
			return
		case txLog := <-erc20.logsReceiver:
			if txLog.Address.Hex() == erc20.address.Hex() {
				var logErr error

				switch txLog.Topics[0].Hex() {
				case transferERC20Hash.Hex():
					transferEvent, err := erc20Contract.ParseTransfer(txLog)
					if err != nil {
						errorChannel <- fmt.Errorf("Error parsing transfers for contract: %v\nError: %v\n", erc20.address.Hex(), err)
					}
					logErr = logERC20Transfer(erc20.contractDBId, transferEvent)
				}

				if logErr != nil {
					errorChannel <- logErr
				}

			}
			erc20.lastKnownBlock = txLog.BlockNumber - 1
		}
	}
}

func NewERC20EventCollector(chainId uint64, address common.Address, startBlock uint64) EventCollector {
	if erc20Contract == nil {
		erc20Contract, _ = contracts.NewERC20(address, nil)
	}

	ec := &ERC20EventCollector{
		eventCollector: eventCollector{
			chainId:        chainId,
			address:        address,
			lastKnownBlock: startBlock,
		},
	}

	return ec
}

func logERC20Transfer(contractId uint64, transfer *contracts.ERC20Transfer) error {
	if db.DBConn == nil {
		log.Println(erc20TransferToString(transfer))
		return nil
	}

	baseEventId := logBaseEvent(contractId, transfer.Raw)

	_, err := db.DBConn.Exec(
		insertTransferERC20Query,
		baseEventId,
		transfer.From.Hex(),
		transfer.To.Hex(),
		transfer.Value.String(),
	)

	if checkUniqueViolation(err, "Transfer") {
		return nil
	}
	if err != nil {
		return err
	}

	err = changeUserBalance(contractId, baseEventId, transfer)
	return err
}

func changeUserBalance(
	contractId uint64,
	baseEventId string,
	transfer *contracts.ERC20Transfer,
) error {
	if transfer.From.Hex() != ZERO_ADDRESS {
		negativeValue := big.NewInt(0)
		negativeValue.Mul(transfer.Value, big.NewInt(-1))

		_, err := db.DBConn.Exec(
			addUserBalanceChange,
			baseEventId,
			transfer.From.Hex(),
			negativeValue.String(),
		)

		if !checkUniqueViolation(err, "change balance from") {
			return err
		}
	}

	_, err := db.DBConn.Exec(
		addUserBalanceChange,
		baseEventId,
		transfer.To.Hex(),
		transfer.Value.String(),
	)
	if checkUniqueViolation(err, "change balance to") {
		return nil
	}
	return err
}

var _ EventCollector = (*ERC20EventCollector)(nil)
