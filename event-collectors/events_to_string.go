package eventcollectors

import (
	"fmt"
	"gather-events/contracts"

	"github.com/ethereum/go-ethereum/core/types"
)

func logToString(chainId uint64, log types.Log) string {
	return fmt.Sprintf(
		"LOG\nChain id: %d\nContract: %v\nTimestamp: %d",
		chainId,
		log.Address.Hex(),
		log.BlockNumber,
	)
}

func erc20TransferToString(transfer *contracts.ERC20Transfer) string {
	return fmt.Sprintf(
		"ERC20 TRANSFER\nFrom: %v\nTo: %v\nAmount: %v",
		transfer.From.Hex(),
		transfer.To.Hex(),
		transfer.Value.String(),
	)
}

func teleportToString(teleport *contracts.BridgeBaseTeleport) string {
	return fmt.Sprintf(
		"TELEPORT\nFrom: %v\nTo: %v\nAmount: %v\nNonce: %v",
		teleport.From.Hex(),
		teleport.To.Hex(),
		teleport.Amount.String(),
		teleport.Nonce.String(),
	)
}

func claimToString(claim *contracts.BridgeBaseClaimed) string {
	return fmt.Sprintf(
		"CLAIM\nFrom: %v\nTo: %v\nAmount: %v\nNonce: %v",
		claim.From.Hex(),
		claim.To.Hex(),
		claim.Amount.String(),
		claim.Nonce.String(),
	)
}

func refundToString(refund *contracts.BridgeBaseRefundStatusChanged) string {
	return fmt.Sprintf(
		"REFUND\nFrom: %v\nNonce: %v\nStatus: %d\nTimestamp: %d",
		refund.From.Hex(),
		refund.Nonce.String(),
		refund.Status,
		refund.Timestamp,
	)
}

func erc721TransferToString(transfer *contracts.ERC721Transfer) string {
	return fmt.Sprintf(
		"ERC721 TRANSFER\nFrom: %v\nTo: %v\nTokenId: %v",
		transfer.From.Hex(),
		transfer.To.Hex(),
		transfer.TokenId.String(),
	)
}
