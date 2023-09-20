package config

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	exampleConfig := `
        providers:
            mainnet: http://mainnet.com
            bsc_testnet: http://bsc-testnet.com
        contracts:
            - type: ERC20
              address: 0x01
              network: mainnet
            - type: Bridge
              address: 0x02
              network: bsc_testnet
    `

	expectedConfig := fullConfig{
		Providers: map[string]string{
			"mainnet":     "http://mainnet.com",
			"bsc_testnet": "http://bsc-testnet.com",
		},
		Contracts: []contractInfo{
			{Type: "ERC20", Address: "0x01", Network: "mainnet"},
			{Type: "Bridge", Address: "0x02", Network: "bsc_testnet"},
		},
	}

	parsedConfig := parseConfig("", []byte(exampleConfig))
	providers := []string{"mainnet", "bsc_testnet"}

	for _, providerName := range providers {
		if parsedConfig.Providers[providerName] != expectedConfig.Providers[providerName] {
			t.Logf(
				"parsed: %v\nexpected: %v\n",
				parsedConfig.Providers[providerName],
				expectedConfig.Providers[providerName],
			)
			t.Fatal("Invalid provider at:", providerName)
		}
	}

	for id := 0; id < len(expectedConfig.Contracts); id++ {
		if parsedConfig.Contracts[id].Type != parsedConfig.Contracts[id].Type &&
			parsedConfig.Contracts[id].Address != parsedConfig.Contracts[id].Address &&
			parsedConfig.Contracts[id].Network != parsedConfig.Contracts[id].Network {
			t.Fatal("Invalid contract at position:", id)
		}
	}
}
