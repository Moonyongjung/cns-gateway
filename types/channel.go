package types

import (
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

type ChannelStruct struct {
	BankMsgSendChan chan *bank.MsgSend
	StoreMsgSendChan chan wasm.MsgStoreCode
	InstantiateMsgSendChan chan wasm.MsgInstantiateContract
	ExecuteMsgSendChan chan wasm.MsgExecuteContract
	QueryMsgSendChan chan wasm.QuerySmartContractStateRequest
	ListcodeMsgSendChan chan wasm.QueryCodesRequest
	ListContractByCodeMsgSendChan chan wasm.QueryContractsByCodeRequest
	DownloadMsgSendChan chan []interface{}
	CodeInfoMsgSendChan chan wasm.QueryCodeRequest
	ContractInfoMsgSendChan chan wasm.QueryContractInfoRequest
	ContractStateAllMsgSendChan chan wasm.QueryAllContractStateRequest
	ContractHistoryMsgSendChan chan wasm.QueryContractHistoryRequest
	PinnedMsgSendChan chan wasm.QueryPinnedCodesRequest

	TxResponseChan chan *sdk.TxResponse
	QueryResponseChan chan string
	
	HttpServerChan chan []byte
	ErrorChan chan HttpResponseStruct
}

func ChannelInit() ChannelStruct{
	var channel ChannelStruct
	channel.BankMsgSendChan = make(chan *bank.MsgSend)
	channel.StoreMsgSendChan = make(chan wasm.MsgStoreCode)
	channel.InstantiateMsgSendChan = make(chan wasm.MsgInstantiateContract)
	channel.ExecuteMsgSendChan = make(chan wasm.MsgExecuteContract)
	channel.QueryMsgSendChan = make(chan wasm.QuerySmartContractStateRequest)
	channel.ListcodeMsgSendChan = make(chan wasm.QueryCodesRequest)
	channel.ListContractByCodeMsgSendChan = make(chan wasm.QueryContractsByCodeRequest)
	channel.DownloadMsgSendChan = make(chan []interface{})
	channel.CodeInfoMsgSendChan = make(chan wasm.QueryCodeRequest)
	channel.ContractInfoMsgSendChan = make(chan wasm.QueryContractInfoRequest)
	channel.ContractStateAllMsgSendChan = make(chan wasm.QueryAllContractStateRequest)
	channel.ContractHistoryMsgSendChan = make(chan wasm.QueryContractHistoryRequest)
	channel.PinnedMsgSendChan = make(chan wasm.QueryPinnedCodesRequest)

	channel.TxResponseChan = make(chan *sdk.TxResponse)
	channel.QueryResponseChan = make(chan string)
	
	channel.HttpServerChan = make(chan []byte)
	channel.ErrorChan = make(chan HttpResponseStruct)

	return channel
}