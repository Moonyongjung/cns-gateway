# CNS Gateway
CNS(Chain Naming Service) gateway links between client and blockchain based on tendermint. Not only the gateway supports CNS system but also HTTP API gateway for blockchain based on Tendermint core. It can support basic functionalities of Smart contract

## Introduction
Most addresses of blockchain are difficult to recognize as address is generated randomly by using key algorithm like secp256k1 and is too long. For easily reminding address, Naming service is that short name(as 'domain') is mapped to blockchain account address similarly DNS. This gateway supports to operate methods are creating/retrieving domain and send coin using domain (not account address) in the CNS system as middleware. It is inspired by [ENS](https://ens.domains/).
This CNS gateway, but also CNS client and CNS Cosmwasm contract composed of CNS system, is able to operate on Tendermint based blockchains. Also, it supports chains included ethermint module for using EVM by selecting secp256k1 is used to general cosmos/tendermint and eth_secp256k1 is used to ethermint module.
The aim of a chain provides EVM by applying to ethermint module is that chain executes solidity contract and cosmwasm contract at the same time. So, `coinType` of HD derivation path in order to create private key set 60(ETH) not 118(Cosmos) in the CNS gateway, because Keplr representing tendermint based chain wallet provides to select HD path but Metamask cannot provide it. (Don't concern about it if this gateway connects the chain without ethermint module)

## Prerequisites
- Mysql
  - run .sql file (./db/sql/table_create.sql)
- Set config file
  - `./config/config.json`
  ```yaml
  {
    "gatewayServerPort": "12000",
    "coinType":"60",
    "chainId": "localnet",
    "restEndpoint": "http://localhost:1317",
    "rpcEndpoint": "http://localhost:26657",
    "denom": "unom",
    "contractPath": "./wasm/",
    "contractName": "contract.wasm",
    "bech32Prefix": "noname",
    "gasLimit": "5000000",
    "feeAmount": "5000"
  }
  ```
  - `gatewayServerPort` : Gateway server TCP port number
  - `coinType` : Blockchain coin type
  - `chainId`, `restEndpoint`, `rpcEndpoint` : Target chain information
  - `denom` : Currency denomination (e.g. uatom)
  - `contractPath`, `contractName` : Contract local path and name to deploy
  - `bech32Prefix` : Prefix of account address and smart contract address
  - `gasLimit`, `feeAmount` : Fees to send transactions
- Set key config file
  - `./config/configKey.json`
  - `keyStorePath`, `keyStoreFilePath` : Private key path
  - `keyOwnerName`, `keyOwnerPw` : Gateway key owner. `keyOwnerPw` must be the same as the password entered in the cli when generating the initial key. (i.e. Enter keyring passpharse of Cosmos SDK)
- Set DB config file
  ```yaml
  {
    "dbUserName":"root",
    "dbPassword":"spdytest",
    "dbHost":"localhost",
    "dbPort":"3306",
    "databaseName":"cnsgateway"
  }
  ```
  - DB environment set
  
## Start
```shell
// wasmd v0.27.0 is not supported by go package (latest version v0.7.2)
go get github.com/CosmWasm/wasmd@0.27.0
go mod tidy
go build cnsgw.go
./cnsgw
```

## Usage
1. Run gateway
2. Initialize gateway key type
  - Use EVM or not
3. Send coin (to use gas fee) to created gw address
  - After send coin, re-run gateway for cheking account response and getting account sequence
4. Store, instantiate CNS contract(use below API_for CNS contract)
  - When store msg & instantiate msg are sended to blockchain using below APIs, contract code ID, contract address and saving file directory are created automatically in the contract info directory(./contract/info)
5. Interact CNS client

## API(Only for CNS contract)
### Instantiate CNS contract
  - (GET) `/api/wasm/default-cns-instantiate`
### Execute CNS contract
  - (POST) `/api/wasm/cns-execute`
  - Request body
  ``` yaml
  {
        "domainName": "domaintest",
        "accountAddress": "noname1qnzpr4znu049ydj7qfcu6x7hyjjsaa4fm4le8w"        
  }
  ```
### Query domain
  - (POST) `/api/wasm/cns-query-by-domain`
  - Request body
  ``` yaml
  {
        "domainName": "domaintest"        
  }
  ```
  - Response data is `accountAddress`
### Query account address
  - (POST) `/api/wasm/cns-query-by-account`
  - Request body
  ``` yaml
  {
        "accountAddress": "noname1qnzpr4znu049ydj7qfcu6x7hyjjsaa4fm4le8w"        
  }
  ```
  - Response data is `domain`

## API(General)
All Parameters of request json value are string type
### Coin send (only gateway address can send)
  - (POST) `/api/bank/send`
  - Request body
  ``` yaml
  {
        "fromAddress": "a",
        "toAddress": "b",
        "amount": "100"
  }
  ```
  - `fromAddress` : GW address
### Store contract
  - (GET) `/api/wasm/store`
  - Request after copy .wasm file to the `contractPath`
### Instantiate 
  - (POST) `/api/wasm/instantiate`
  - Request body
  ```yaml
  {
        "codeId": "1",
        "amount": "1000unoname",
        "label": "contract inst",
        "initMsg": "{\"purchase_price\":{\"amount\":\"100\",\"denom\":\"unoname\"}"
  }
  ```
  - `codeId` : Contract code ID
### Execute
  - (POST) `/api/wasm/execute`
  - Request body
  ```yaml
  {
        "contractAddress": "noname19h0d6k4mtxw5qjr0aretjy9kwyem0hxclf88ka2uwjn47e90mqrqk4tkjt",
        "amount": "0unoname",
        "execMsg": "{\"register\":{\"name\":\"fred\"}}"
  }
  ```
### Query - Contract state
  - (POST) `/api/wasm/query`
  - Request body
  ```yaml
  {
        "contractAddress": "noname19h0d6k4mtxw5qjr0aretjy9kwyem0hxclf88ka2uwjn47e90mqrqk4tkjt",
        "queryMsg": "{\"resolve_record\": {\"name\": \"fred\"}}"
  }
  ```
### Query - Contract list
  - (GET) `/api/wasm/list-code`
### Query - Contract information by code ID
  - (POST) `/api/wasm/list-contract-by-code`
  - Request body
  ```yaml
  {
        "codeId": "1"
  }
  ```
### Query - Download contract wasm file
  - (POST) `/api/wasm/download`
  - Request body
  ```yaml
  {
        "codeId": "1",
        "downloadFileName":"download"
  }
  ```
### Query - Code information for a given code ID
  - (POST) `/api/wasm/code-info`
  - Request body
  ```yaml
  {
        "codeId": "1"
  }
  ```
### Query - Contract information for a given contract address
  - (POST) `/api/wasm/contract-info`
  - Request body
  ```yaml
  {
        "contractAddress": "noname19h0d6k4mtxw5qjr0aretjy9kwyem0hxclf88ka2uwjn47e90mqrqk4tkjt"
  }
  ```
### Query - All of contract internal state
  - (POST) `/api/wasm/contract-state-all`
  - Request body
  ```yaml
  {
        "contractAddress": "noname19h0d6k4mtxw5qjr0aretjy9kwyem0hxclf88ka2uwjn47e90mqrqk4tkjt"
  }
  ```
### Query - Contract history
  - (POST) `/api/wasm/contract-history`
  - Request body
  ```yaml
  {
        "contractAddress": "noname19h0d6k4mtxw5qjr0aretjy9kwyem0hxclf88ka2uwjn47e90mqrqk4tkjt"
  }
  ```
### Query - Pinned code
  - (Get) `/api/wasm/pinned`  

### Gatway HTTP Response 
  - `resCode` is int type, `resMsg` and `resData` are string type
  - Response type of Most over APIs
  - e.g.
  ```yaml
  {
        "resCode": 0,
        "resMsg": "",
        "resData": ""
  }
  ```

## Future work
- Supporting EVM (solidity smart contract)