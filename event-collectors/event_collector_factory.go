package eventcollectors

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
)

type (
	ContractType         string
	EventCollectorConfig struct {
		Name           ContractType
		Address        common.Address
		ChainId        uint64
		LastKnownBlock uint64
	}
)

var (
	ERC20  ContractType = "ERC20"
	ERC721 ContractType = "ERC721"
	BRIDGE ContractType = "Bridge"
)

func CreateEventCollector(ecConfig EventCollectorConfig) EventCollector {
	var ec EventCollector
	switch ecConfig.Name {
	case ERC20:
		ec = NewERC20EventCollector(
			ecConfig.ChainId,
			ecConfig.Address,
			ecConfig.LastKnownBlock,
		)

	case ERC721:
		ec = NewERC721EventCollector(
			ecConfig.ChainId,
			ecConfig.Address,
			ecConfig.LastKnownBlock,
		)
	case BRIDGE:
		ec = NewBridgeEventCollector(
			ecConfig.ChainId,
			ecConfig.Address,
			ecConfig.LastKnownBlock,
		)
	default:
		log.Fatalln("Unknown collector name:", ecConfig.Name)
	}

	return ec
}
