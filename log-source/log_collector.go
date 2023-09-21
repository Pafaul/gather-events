package logsource

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type (
	LogsCollector interface {
		ChainId() uint64
		LogsChannel() <-chan types.Log
		ErrorChannel() <-chan error
		SetFilterAddresses([]common.Address)
		CollectLogs(context.Context, uint64) (<-chan types.Log, <-chan error)
	}

	logsCollector struct {
		client       *ethclient.Client
		chainId      uint64
		addresses    []common.Address
		logsChannel  chan types.Log
		errorChannel chan error
	}
)

const (
	MAX_BLOCK_DELTA = 500
)

func (l *logsCollector) ChainId() uint64 {
	return l.chainId
}

func (l *logsCollector) LogsChannel() <-chan types.Log {
	return l.logsChannel
}

func (l *logsCollector) ErrorChannel() <-chan error {
	return l.errorChannel
}

func (l *logsCollector) CollectLogs(ctx context.Context, startBlock uint64) (<-chan types.Log, <-chan error) {
	go l.startLogsCollection(ctx, startBlock)

	return l.logsChannel, l.errorChannel
}

func (l *logsCollector) SetFilterAddresses(addresses []common.Address) {
	l.addresses = addresses
}

func (l *logsCollector) startLogsCollection(ctx context.Context, startBlock uint64) {
	// if we exit we close the channels to stop possible data flow
	defer (func() {
		close(l.errorChannel)
		close(l.logsChannel)
	})()

	// Default call to check if client is available
	_, err := l.client.ChainID(ctx)
	if err != nil {
		l.errorChannel <- err
		return
	}

	blockNumber := startBlock
	endBlock := blockNumber + MAX_BLOCK_DELTA
	currentBlock, blockErr := l.client.BlockNumber(ctx)
	if blockErr != nil {
		l.errorChannel <- blockErr
		return
	}

	// Loop stop is controlled by provided ctx
	for {
		select {
		case <-ctx.Done():
			return
		default:
			currentBlock, blockErr = l.client.BlockNumber(ctx)
			if blockErr != nil {
				l.errorChannel <- blockErr
				return
			}

			if endBlock > currentBlock {
				endBlock = currentBlock
				for endBlock == currentBlock {
					currentBlock, blockErr = l.client.BlockNumber(ctx)
					if blockErr != nil {
						l.errorChannel <- blockErr
						return
					}
				}
			}

			if startBlock > endBlock {
				startBlock = endBlock
			}

			log.Printf("Blocks: %d - %d\n", blockNumber, endBlock)

			filter := ethereum.FilterQuery{
				Addresses: l.addresses,
				FromBlock: big.NewInt(int64(blockNumber)),
				ToBlock:   big.NewInt(int64(endBlock)),
			}
			logs, logErr := l.client.FilterLogs(context.Background(), filter)
			if logErr != nil {
				l.errorChannel <- logErr
				return
			}

			if len(logs) != 0 {
				for id := range logs {
					select {
					case <-ctx.Done():
						return
					default:
						l.logsChannel <- logs[id]
					}
				}
			}

			blockNumber = endBlock
			endBlock += MAX_BLOCK_DELTA
		}
	}
}

func NewLogsCollector(client *ethclient.Client, chainId uint64) LogsCollector {
	lc := &logsCollector{
		client:       client,
		chainId:      chainId,
		logsChannel:  make(chan types.Log),
		errorChannel: make(chan error),
		addresses:    make([]common.Address, 0),
	}

	return lc
}

var _ LogsCollector = (*logsCollector)(nil)
