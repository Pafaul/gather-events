package main

import (
	"gather-events/config"
	"gather-events/db"
	logsource "gather-events/log-source"
	"log"
	"time"
)

func main() {
	db.OpenDBConnection()
	config.ParseConfig()

	providers := config.GetProivders()

	collectors := config.GetEventCollectors()

	for _, collector := range collectors {
		err := collector.LoadContractInfo()
		if err != nil {
			log.Fatalf(
				"Could not load contract state\nContract address: %v\nChain id: %d\nErr: %v\n",
				collector.Address().Hex(),
				collector.ChainId(),
				err,
			)
		}
	}

	hubs := make([]Hub, 0)

	for _, provider := range providers {
		logCollector := logsource.NewLogsCollector(provider.Client, provider.ChainId)
		hub, err := NewHub(provider, collectors, logCollector)
		if hub == nil {
			log.Println(err)
			continue
		}
		hubs = append(hubs, hub)
	}

	for _, h := range hubs {
		h.StartHub()
	}

	// TODO: remove time limitations and handle exit gracefully
	time.Sleep(time.Hour)

	for _, hub := range hubs {
		hub.StopHub()
	}
}
