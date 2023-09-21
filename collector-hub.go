package main

import (
	"context"
	"fmt"
	eventcollectors "gather-events/event-collectors"
	logsource "gather-events/log-source"
	"gather-events/providers"
	"log"
	"math"

	"github.com/ethereum/go-ethereum/common"
)

type (
	Hub interface {
		StopHub()
		StartHub()
	}

	hub struct {
		provider             providers.Provider
		collectorCtxCancel   context.CancelFunc
		collector            logsource.LogsCollector
		broadcasterCtxCancel context.CancelFunc
		broadcaster          logsource.LogBroadcaster
		eventCollectorErr    chan error
		eventCollectors      []eventcollectors.EventCollector
		stopChannels         []chan bool
	}
)

func getCollectorForChainId(chainId uint64, collectors []eventcollectors.EventCollector) []eventcollectors.EventCollector {
	collectorsForProvider := make([]eventcollectors.EventCollector, 0)
	for _, collector := range collectors {
		if collector.ChainId() == chainId {
			collectorsForProvider = append(collectorsForProvider, collector)
		}
	}
	return collectorsForProvider
}

func getAddessesForProvider(collectors []eventcollectors.EventCollector) []common.Address {
	addressesForProvider := make([]common.Address, len(collectors))
	for id, collector := range collectors {
		addressesForProvider[id] = collector.Address()
	}
	return addressesForProvider
}

func NewHub(
	provider providers.Provider,
	collectors []eventcollectors.EventCollector,
	logCollector logsource.LogsCollector,
) (Hub, error) {
	h := &hub{}
	h.provider = provider
	h.collector = logCollector

	collectorsForProvider := getCollectorForChainId(provider.ChainId, collectors)
	if len(collectorsForProvider) == 0 {
		return nil, fmt.Errorf("No event collectors for provider: %d", provider.ChainId)
	}
	h.eventCollectors = collectorsForProvider

	addressesForProvider := getAddessesForProvider(collectorsForProvider)
	logCollector.SetFilterAddresses(addressesForProvider)

	logBroadcaster := logsource.NewLogBroadcaster(
		logCollector.LogsChannel(),
	)
	h.broadcaster = logBroadcaster

	for _, collector := range collectorsForProvider {
		collector.SubscribeToLogs(logBroadcaster)
	}

	return h, nil
}

// Starting in order:
// 1. Event collectors
// 2. Event broadcaster
// 3. Logs collector
// In this order every instance will be ready and waiting on the data source
func (h *hub) StartHub() {
	h.stopChannels = make([]chan bool, len(h.eventCollectors))
	h.eventCollectorErr = make(chan error)
	for id, collector := range h.eventCollectors {
		h.stopChannels[id] = make(chan bool)
		go collector.StartEventCollector(h.stopChannels[id], h.eventCollectorErr)
	}

	broadcasterCtx, broacasterCancel := context.WithCancel(context.Background())
	h.broadcasterCtxCancel = broacasterCancel

	collectorCtx, collectorCancel := context.WithCancel(context.Background())
	h.collectorCtxCancel = collectorCancel

	h.broadcaster.StartBroadcasting(broadcasterCtx)
	_, logErr := h.collector.CollectLogs(
		collectorCtx,
		getMinLastBlock(h.eventCollectors),
	)

	go func() {
		err := <-logErr
		h.StopHub()
		log.Fatalln(err)
	}()
}

func (h *hub) StopHub() {
	h.collectorCtxCancel()
	h.broadcasterCtxCancel()

	for _, collector := range h.eventCollectors {
		err := collector.SaveLastKnownBlock()
		if err != nil {
			log.Println(err)
		}
	}

	for _, stopChannel := range h.stopChannels {
		stopChannel <- true
	}
}

func getMinLastBlock(ecs []eventcollectors.EventCollector) uint64 {
	minBlock := uint64(math.MaxUint64)

	for _, ec := range ecs {
		if ec.LastKnownBlock() < minBlock {
			minBlock = ec.LastKnownBlock()
		}
	}

	return minBlock
}

var _ Hub = (*hub)(nil)
