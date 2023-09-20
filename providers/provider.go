package providers

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type (
	Provider struct {
		Name    string
		ChainId uint64
		Client  *ethclient.Client
	}
)
