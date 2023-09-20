# how to run

ENVS

- DB, at least one should be provided, PG_URL will be preferred:
    - ```PG_URL```: connection string to postgres db
    - ```SQLITE_URL```: connection string to sqlite db
- config:
    - ```CONFIG_FILE```: path to config file, default: ```$(pwd)/config.yaml```

Config

Template config:
```yaml
providers: # list of available providers
  my-provider: https://my-provider.com/ # provider name and provider url 
  my-other-provider: https://my-other-provider.com/

contracts: # list of contracts to track
  - type: ERC20 # contract type, available contract types: ERC20/ERC721/Bridge
    address: "0x0000000000000000000000000000000000000001" # contract address to track
    network: my-provider # provider for the contract to use
    start_block: 836187 # block to start tracking from (usually - block when contract was deployed)


  - type: ERC721
    address: "0x0000000000000000000000000000000000000002"
    network: my-other-provider 
    start_block: 12378 
```
