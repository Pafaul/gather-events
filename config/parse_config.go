package config

import (
	"context"
	eventcollectors "gather-events/event-collectors"
	"gather-events/providers"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v3"
)

/*
providers:

	name: string

contracts:

	type:
	address:
	network_name:
*/
type (
	contractInfo struct {
		Type       string `yaml:"type"`
		Address    string `yaml:"address"`
		Network    string `yaml:"network"`
		StartBlock uint64 `yaml:"start_block"`
	}
	fullConfig struct {
		Providers map[string]string `yaml:"providers"`
		Contracts []contractInfo    `yaml:"contracts"`
	}
)

var (
	configfile                                       = "config.yaml"
	parsedConfig       fullConfig                    = fullConfig{}
	availableProviders []providers.Provider          = make([]providers.Provider, 0)
	nameToProvider     map[string]providers.Provider = make(map[string]providers.Provider)
	chainIdToName      map[uint64]string             = make(map[uint64]string)
	nameToChainId      map[string]uint64             = make(map[string]uint64)
)

func ParseConfig() {
	configFilename, configFilenameProvided := os.LookupEnv("CONFIG_FILE")
	if configFilenameProvided {
		configfile = configFilename
	}

	parsedConfig := parseConfig(configfile, nil)
	for name, url := range parsedConfig.Providers {
		getProviderFromURL(name, url)
	}

	for _, config := range parsedConfig.Contracts {
		getContractFromConfig(config)
	}
}

func GetProivders() []providers.Provider {
	return availableProviders
}

func getProviderFromURL(name string, providerURL string) providers.Provider {
	provider, providerExists := nameToProvider[name]
	if providerExists {
		return provider
	}

	client, clientErr := ethclient.Dial(providerURL)
	if clientErr != nil {
		log.Fatalf(
			"Could not initialize provider %v with url:\n%v",
			name,
			providerURL,
		)
	}

	chainId, chainIdErr := client.ChainID(context.Background())
	if chainIdErr != nil {
		log.Fatalf(
			"Could not run test query on provider %v with url:\n%v",
			name,
			providerURL,
		)
	}

	provider = providers.Provider{
		Name:    name,
		ChainId: chainId.Uint64(),
		Client:  client,
	}

	availableProviders = append(availableProviders, provider)
	nameToProvider[name] = provider
	chainIdToName[chainId.Uint64()] = name
	nameToChainId[name] = chainId.Uint64()

	return provider
}

func GetProviderByName(name string) providers.Provider {
	return nameToProvider[name]
}

func GetProviderByChainId(chainId uint64) providers.Provider {
	return nameToProvider[chainIdToName[chainId]]
}

func GetEventCollectors() []eventcollectors.EventCollector {
	conf := parseConfig(configfile, nil)

	if len(conf.Contracts) == 0 {
		log.Fatalln("No contracts provided")
	}

	contractsConfig := make([]eventcollectors.EventCollector, len(conf.Contracts))
	for id, contractInfo := range conf.Contracts {
		contractsConfig[id] = getContractFromConfig(contractInfo)
	}

	return contractsConfig
}

func getContractFromConfig(contractConfig contractInfo) eventcollectors.EventCollector {
	var eventCollector eventcollectors.EventCollector

	contractAddressValid := common.IsHexAddress(contractConfig.Address)
	if !contractAddressValid {
		log.Fatalf(
			"Invalid address provided for %v with address %v",
			contractConfig.Type,
			contractConfig.Address,
		)
	}

	_ = GetProviderByName(contractConfig.Network)

	chainId, ok := nameToChainId[contractConfig.Network]
	if !ok {
		log.Fatalf(
			"Could not get provider for contrac\nAddress: %v\nProvider name:%v\n",
			contractConfig.Address,
			contractConfig.Network,
		)
	}
	eventCollectorConfig := eventcollectors.EventCollectorConfig{
		Name:           eventcollectors.ContractType(contractConfig.Type),
		Address:        common.HexToAddress(contractConfig.Address),
		LastKnownBlock: contractConfig.StartBlock,
		ChainId:        chainId,
	}

	eventCollector = eventcollectors.CreateEventCollector(eventCollectorConfig)

	return eventCollector
}

func parseConfig(filename string, content []byte) fullConfig {
	if len(content) == 0 {
		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			log.Fatalf("File %v does not exist", filename)
		}

		var readErr error
		content, readErr = os.ReadFile(filename)
		if readErr != nil {
			log.Fatalln(readErr)
		}
	}

	config := fullConfig{}

	err := yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalln(err)
	}

	return config
}
