package types

var (
	//-- HTTP server API
	ApiVersion = "/api/"

	BankSend               = "bank/send"
	WasmStore              = "wasm/store"
	WasmInstantiate        = "wasm/instantiate"
	WasmExecute            = "wasm/execute"
	WasmQuery              = "wasm/query"
	WasmListCode           = "wasm/list-code"
	WasmListContractByCode = "wasm/list-contract-by-code"
	WasmDownload           = "wasm/download"
	WasmCodeInfo           = "wasm/code-info"
	WasmContractInfo       = "wasm/contract-info"
	WasmContractStateAll   = "wasm/contract-state-all"
	WasmContractHistory    = "wasm/contract-history"
	WasmPinned             = "wasm/pinned"

	CnsWasmInstantiate    = "wasm/default-cns-instantiate"
	CnsWasmExecute        = "wasm/cns-execute"
	CnsWasmQueryByDomain  = "wasm/cns-query-by-domain"
	CnsWasmQueryByAccount = "wasm/cns-query-by-account"

	//-- CNS clinet API
	ClientApiVersion = "/client/"

	WalletCreate  = "wallet/create/"
	WalletAddress = "wallet/address/"
	DomainMapping = "domain/mapping/"
	DomainConfirm = "domain/confirm/"
	SendIndex     = "send/index/"
	SendInquiry   = "send/inquiry/"
)
