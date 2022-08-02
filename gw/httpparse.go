package gw

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Moonyongjung/cns-gw/contract"
	"github.com/Moonyongjung/cns-gw/msg"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func doTransactionbyType(request *http.Request, channel cns.ChannelStruct) {
	var body []byte
	var c msg.ContractMsg

	checkRequest(request)
	reqType := strings.TrimLeft(request.URL.Path, cns.ApiVersion)

	if request.Method == "POST" {
		bodyByte, err := ioutil.ReadAll(request.Body)
		if err != nil {
			httpParseErrReturn(err, channel)
		}
		body = bodyByte
	}

	if reqType == cns.BankSend {
		msg, err := c.MsgConvert(body, reqType).MakeBankMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmStore {
		msg, err := c.MakeStoreMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmInstantiate {
		msg, err := c.MsgConvert(body, reqType).MakeInstantiateMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.CnsWasmInstantiate {
		c.InstantiateMsg = contract.DefaultCnsContractInstantiateMsg()
		msg, err := c.MakeInstantiateMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmExecute {
		msg, err := c.MsgConvert(body, reqType).MakeExecuteMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.CnsWasmExecute {
		c.ExecuteMsg = contract.CnsContractExecuteMsg(body)
		msg, err := c.MakeExecuteMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmQuery {
		msg, err := c.MsgConvert(body, reqType).MakeQueryMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.CnsWasmQueryByDomain {
		c.QueryMsg = contract.CnsContractQueryDomainMsg(body)
		msg, err := c.MakeQueryMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.CnsWasmQueryByAccount {
		c.QueryMsg = contract.CnsContractQueryAccountMsg(body)
		msg, err := c.MakeQueryMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmListCode {
		msg, err := c.MakeListcodeMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmListContractByCode {
		msg, err := c.MsgConvert(body, reqType).MakeListContractByCodeMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmDownload {
		msg, err := c.MsgConvert(body, reqType).MakeDownloadMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmCodeInfo {
		msg, err := c.MsgConvert(body, reqType).MakeCodeInfoMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmContractInfo {
		msg, err := c.MsgConvert(body, reqType).MakeContractInfoMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmContractStateAll {
		msg, err := c.MsgConvert(body, reqType).MakeContractStateAllMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmContractHistory {
		msg, err := c.MsgConvert(body, reqType).MakeContractHistoryMsg()
		returnDataUsingChannel(err, reqType, channel, msg)

	} else if reqType == cns.WasmPinned {
		msg, err := c.MakePinnedMsg()
		returnDataUsingChannel(err, reqType, channel, msg)
	}
}

func checkRequest(request *http.Request) {
	util.LogHttpServer("Client IP addr : " + request.RemoteAddr)
	util.LogHttpServer("Request API : " + request.URL.Path)
}

func httpParseErrReturn(err error,
	channel cns.ChannelStruct) {

	util.LogErr("ERROR, ", err)
	response := util.HttpResponseTypeStruct(106, "", err.Error())
	channel.ErrorChan <- response
}

func returnDataUsingChannel(err error,
	reqType string,
	channel cns.ChannelStruct,
	msg interface{}) {

	if err != nil {
		httpParseErrReturn(err, channel)
	} else {
		if reqType == cns.BankSend {
			msgConv := msg.(*bank.MsgSend)
			channel.BankMsgSendChan <- msgConv

		} else if reqType == cns.WasmStore {
			msgConv := msg.(wasm.MsgStoreCode)
			channel.StoreMsgSendChan <- msgConv

		} else if reqType == cns.WasmInstantiate || reqType == cns.CnsWasmInstantiate {
			msgConv := msg.(wasm.MsgInstantiateContract)
			channel.InstantiateMsgSendChan <- msgConv

		} else if reqType == cns.WasmExecute || reqType == cns.CnsWasmExecute {
			msgConv := msg.(wasm.MsgExecuteContract)
			channel.ExecuteMsgSendChan <- msgConv

		} else if reqType == cns.WasmQuery || reqType == cns.CnsWasmQueryByDomain || reqType == cns.CnsWasmQueryByAccount {
			msgConv := msg.(wasm.QuerySmartContractStateRequest)
			channel.QueryMsgSendChan <- msgConv

		} else if reqType == cns.WasmListCode {
			msgConv := msg.(wasm.QueryCodesRequest)
			channel.ListcodeMsgSendChan <- msgConv

		} else if reqType == cns.WasmListContractByCode {
			msgConv := msg.(wasm.QueryContractsByCodeRequest)
			channel.ListContractByCodeMsgSendChan <- msgConv

		} else if reqType == cns.WasmDownload {
			msgConv := msg.([]interface{})
			channel.DownloadMsgSendChan <- msgConv

		} else if reqType == cns.WasmCodeInfo {
			msgConv := msg.(wasm.QueryCodeRequest)
			channel.CodeInfoMsgSendChan <- msgConv

		} else if reqType == cns.WasmContractInfo {
			msgConv := msg.(wasm.QueryContractInfoRequest)
			channel.ContractInfoMsgSendChan <- msgConv

		} else if reqType == cns.WasmContractStateAll {
			msgConv := msg.(wasm.QueryAllContractStateRequest)
			channel.ContractStateAllMsgSendChan <- msgConv

		} else if reqType == cns.WasmContractHistory {
			msgConv := msg.(wasm.QueryContractHistoryRequest)
			channel.ContractHistoryMsgSendChan <- msgConv

		} else if reqType == cns.WasmPinned {
			msgConv := msg.(wasm.QueryPinnedCodesRequest)
			channel.PinnedMsgSendChan <- msgConv
		}
	}
}
