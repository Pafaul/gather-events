package db

var (
	pgContractInfoTable = `
        CREATE TABLE IF NOT EXISTS contracts (
            contract_id serial PRIMARY KEY,
            contract_address VARCHAR (42) NOT NULL,
            chain_id INT NOT NULL,
            last_known_block BIGINT NOT NULL,
            UNIQUE (contract_address, chain_id)
        )
    `

	pgBaseEventTable = `
        CREATE TABLE IF NOT EXISTS base_event (
            unique_hash VARCHAR (66) PRIMARY KEY,
            contract_id INT NOT NULL,
            block_number BIGINT NOT NULL,
            CONSTRAINT base_event_contract
                FOREIGN KEY (contract_id)
                REFERENCES contracts (contract_id)
        )
    `

	pgERC20TransfersTable = `
        CREATE TABLE IF NOT EXISTS erc20_transfers (
            transfer_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            CONSTRAINT base_event_erc20_transfer
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	pgBridgeTeleportsTable = `
        CREATE TABLE IF NOT EXISTS bridge_teleports (
            teleport_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            nonce NUMERIC (80) NOT NULL,
            refund_status INT NOT NULL,
            CONSTRAINT base_event_bridge_telepots
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	pgBridgeRefundStatusTable = `
        CREATE TABLE IF NOT EXISTS bridge_teleport_refund_status (
            refund_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            teleport_id INT NOT NULL,
            status smallint NOT NULL,
            timestamp bigint NOT NULL,
            CONSTRAINT base_event_refund_status
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            CONSTRAINT refund_status_teleport
                FOREIGN KEY (teleport_id)
                REFERENCES bridge_teleports (teleport_id),
            UNIQUE (base_event_id)
        )
    `

	pgBridgeClaimsTable = `
        CREATE TABLE IF NOT EXISTS bridge_claims (
            claim_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            nonce NUMERIC (80) NOT NULL,
            CONSTRAINT base_event_bridge_claims
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	pgERC20BalanceChangesTable = `
        CREATE TABLE IF NOT EXISTS user_erc20_balance (
            base_event_id VARCHAR (66) NOT NULL, 
            user_address VARCHAR (42) NOT NULL,
            balance_change NUMERIC (80) NOT NULL,
            PRIMARY KEY (base_event_id, user_address),
            UNIQUE (base_event_id, user_address)
        )
    `

	pgERC721Transfers = `
        CREATE TABLE IF NOT EXISTS erc721_transfers (
            transfer_id serial PRIMARY KEY,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            token_id NUMERIC (80) NOT NULL,
            CONSTRAINT base_event_erc721_transfer
                FOREIGN KEY (base_event_id)
                REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	pgERC721Balances = `
        CREATE TABLE IF NOT EXISTS erc721_balances (
            contract_id INT NOT NULL, 
            token_id NUMERIC (80) NOT NULL,
            user_address VARCHAR (42) NOT NULL,
            CONSTRAINT erc721_balances_contract_id
                FOREIGN KEY (contract_id)
                REFERENCES contracts (contract_id),
            PRIMARY KEY (contract_id, token_id)
        )
    `
)

var (
	sqliteContractInfoTable = `
        CREATE TABLE IF NOT EXISTS contracts (
            contract_id INTEGER PRIMARY KEY   AUTOINCREMENT,
            contract_address VARCHAR (42) NOT NULL,
            chain_id INT NOT NULL,
            last_known_block BIGINT NOT NULL,
            UNIQUE (contract_address, chain_id)
        )
    `

	sqliteBaseEventTable = `
        CREATE TABLE IF NOT EXISTS base_event (
            unique_hash VARCHAR (66) PRIMARY KEY,
            contract_id INT NOT NULL,
            block_number BIGINT NOT NULL,
            FOREIGN KEY (contract_id) REFERENCES contracts (contract_id)
        )
    `

	sqliteERC20TransfersTable = `
        CREATE TABLE IF NOT EXISTS erc20_transfers (
            transfer_id INTEGER PRIMARY KEY   AUTOINCREMENT,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            FOREIGN KEY (base_event_id) REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	sqliteBridgeTeleportsTable = `
        CREATE TABLE IF NOT EXISTS bridge_teleports (
            teleport_id INTEGER PRIMARY KEY   AUTOINCREMENT,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            nonce NUMERIC (80) NOT NULL,
            refund_status INT NOT NULL,
            FOREIGN KEY (base_event_id) REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	sqliteBridgeRefundStatusTable = `
        CREATE TABLE IF NOT EXISTS bridge_teleport_refund_status (
            refund_id INTEGER PRIMARY KEY   AUTOINCREMENT,
            base_event_id VARCHAR (66) NOT NULL,
            teleport_id INT NOT NULL,
            status smallint NOT NULL,
            timestamp bigint NOT NULL,
            FOREIGN KEY (base_event_id) REFERENCES base_event (unique_hash),
            FOREIGN KEY (teleport_id) REFERENCES bridge_teleports (teleport_id),
            UNIQUE (base_event_id)
        )
    `

	sqliteBridgeClaimsTable = `
        CREATE TABLE IF NOT EXISTS bridge_claims (
            claim_id INTEGER PRIMARY KEY   AUTOINCREMENT,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            amount NUMERIC (80) NOT NULL,
            nonce NUMERIC (80) NOT NULL,
            FOREIGN KEY (base_event_id) REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	sqliteERC20BalanceChangesTable = `
        CREATE TABLE IF NOT EXISTS user_erc20_balance (
            base_event_id VARCHAR (66) NOT NULL, 
            user_address VARCHAR (42) NOT NULL,
            balance_change NUMERIC (80) NOT NULL,
            PRIMARY KEY (base_event_id, user_address),
            UNIQUE (base_event_id, user_address)
        )
    `

	sqliteERC721Transfers = `
        CREATE TABLE IF NOT EXISTS erc721_transfers (
            transfer_id INTEGER PRIMARY KEY   AUTOINCREMENT,
            base_event_id VARCHAR (66) NOT NULL,
            from_address VARCHAR (42) NOT NULL,
            to_address VARCHAR (42) NOT NULL,
            token_id NUMERIC (80) NOT NULL,
            FOREIGN KEY (base_event_id) REFERENCES base_event (unique_hash),
            UNIQUE (base_event_id)
        )
    `

	sqliteERC721Balances = `
        CREATE TABLE IF NOT EXISTS erc721_balances (
            contract_id INT NOT NULL, 
            token_id NUMERIC (80) NOT NULL,
            user_address VARCHAR (42) NOT NULL,
            FOREIGN KEY (contract_id) REFERENCES contracts (contract_id),
            PRIMARY KEY (contract_id, token_id)
        )
    `
)
