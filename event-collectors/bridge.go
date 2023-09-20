package eventcollectors

import (
	"gather-events/contracts"
	"gather-events/db"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type (
	BridgeEventCollector struct {
		eventCollector
	}
)

var (
	bridgeContract      *contracts.BridgeBase
	teleportHash        = crypto.Keccak256Hash([]byte("Teleport(address,address,uint256,uint256)"))
	claimHash           = crypto.Keccak256Hash([]byte("Claimed(address,address,uint256,uint256)"))
	refundStatusHash    = crypto.Keccak256Hash([]byte("RefundStatusChanged(address,uint256,uint8,uint64)"))
	insertTeleportQuery = `
        INSERT INTO bridge_teleports (
            base_event_id, 
            from_address, 
            to_address, 
            amount, 
            nonce,
            refund_status
        )
        VALUES ($1, $2, $3, $4, $5, 0) 
    `
	insertClaimQuery = `
        INSERT INTO bridge_claims (
            base_event_id, 
            from_address, 
            to_address, 
            amount, 
            nonce
        )
        VALUES ($1, $2, $3, $4, $5)
    `
	insertRefundQuery = `
        INSERT INTO bridge_refund_status (
            base_event_id, 
            teleport_id, 
            status, 
            timestamp
        )
        VALUES ($1, $2, $3, $4)
    `
	updateTeleportRefundStatusQuery = `
        UPDATE bridge_teleports 
        SET refund_status = $1
        WHERE base_event_id = $2;
    `
)

func (bridge *BridgeEventCollector) StartEventCollector(stopChannel <-chan bool, errorChannel chan<- error) {
	for {
		select {
		case <-stopChannel:
			return
		case txLog := <-bridge.logsReceiver:
			if txLog.Address.Hex() == bridge.address.Hex() {
				var logErr error

				switch txLog.Topics[0].Hex() {
				case teleportHash.Hex():
					teleport, err := bridgeContract.ParseTeleport(txLog)
					if err != nil {
						errorChannel <- formatECError("Teleport", &bridge.eventCollector, &txLog, err)
						return
					}
					logErr = logTeleport(bridge.contractDBId, teleport)
				case claimHash.Hex():
					claim, err := bridgeContract.ParseClaimed(txLog)
					if err != nil {
						errorChannel <- formatECError("Claim", &bridge.eventCollector, &txLog, err)
					}
					logErr = logClaim(bridge.contractDBId, claim)
				case refundStatusHash.Hex():
					refundStatus, err := bridgeContract.ParseRefundStatusChanged(txLog)
					if err != nil {
						errorChannel <- formatECError("RefundStatusChanged", &bridge.eventCollector, &txLog, err)
					}
					logErr = logRefund(bridge.contractDBId, refundStatus)
				}

				if logErr != nil {
					errorChannel <- logErr
				}

			}
			bridge.lastKnownBlock = txLog.BlockNumber - 1
		}
	}
}

func logTeleport(contractId uint64, teleport *contracts.BridgeBaseTeleport) error {
	if db.DBConn == nil {
		log.Println(teleportToString(teleport))
		return nil
	}

	baseEventId := logBaseEvent(contractId, teleport.Raw)

	_, err := db.DBConn.Exec(
		insertTeleportQuery,
		baseEventId,
		teleport.From.Hex(),
		teleport.To.Hex(),
		teleport.Amount.String(),
		teleport.Nonce.String(),
	)

	if checkUniqueViolation(err, "Teleport") {
		return nil
	}

	return err
}

func logClaim(contractId uint64, claim *contracts.BridgeBaseClaimed) error {
	if db.DBConn == nil {
		log.Println(claimToString(claim))
		return nil
	}

	baseEventId := logBaseEvent(contractId, claim.Raw)

	_, err := db.DBConn.Exec(
		insertClaimQuery,
		baseEventId,
		claim.From.Hex(),
		claim.To.Hex(),
		claim.Amount.String(),
		claim.Nonce.String(),
	)

	if checkUniqueViolation(err, "Teleport") {
		return nil
	}
	return err
}

func logRefund(contractId uint64, refund *contracts.BridgeBaseRefundStatusChanged) error {
	if db.DBConn == nil {
		log.Print(refundToString(refund))
		return nil
	}

	baseEventId := logBaseEvent(contractId, refund.Raw)

	_, err := db.DBConn.Exec(
		insertRefundQuery,
		baseEventId,
		refund.Nonce.String(),
		refund.Status,
		refund.Timestamp,
	)

	if !checkUniqueViolation(err, "Refund status") {
		return err
	}

	_, err = db.DBConn.Exec(
		updateTeleportRefundStatusQuery,
		refund.Status,
		baseEventId,
	)

	return err
}

func NewBridgeEventCollector(chainId uint64, address common.Address, lastKnownBlock uint64) EventCollector {
	if bridgeContract == nil {
		bridgeContract, _ = contracts.NewBridgeBase(address, nil)
	}

	bridgeEc := &BridgeEventCollector{
		eventCollector: eventCollector{
			chainId:        chainId,
			address:        address,
			lastKnownBlock: lastKnownBlock,
		},
	}

	return bridgeEc
}

var _ EventCollector = (*BridgeEventCollector)(nil)
