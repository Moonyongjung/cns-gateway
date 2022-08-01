package gw

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Moonyongjung/cns-gw/contract"
	"github.com/Moonyongjung/cns-gw/msg"
	"github.com/Moonyongjung/cns-gw/types"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
)

func doTransactionbyType(request *http.Request, channel cns.ChannelStruct) {
	var body []byte
	var response cns.HttpResponseStruct
	var c msg.ContractMsg

	checkRequest(request)
	reqType := strings.TrimLeft(request.URL.Path, cns.ApiVersion)

	if request.Method == "POST" {
		bodyByte, err := ioutil.ReadAll(request.Body)
		if err != nil {
			httpParseErrReturn(err, response, channel)
		}
		body = bodyByte
	}

	if reqType == cns.BankSend {
		msg, err := c.MsgConvert(body, reqType).MakeBankMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.BankMsgSendChan <- msg
		}
	} else if reqType == cns.WasmStore {
		msg, err := c.MakeStoreMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.StoreMsgSendChan <- msg
		}
	} else if reqType == cns.WasmInstantiate {
		msg, err := c.MsgConvert(body, reqType).MakeInstantiateMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.InstantiateMsgSendChan <- msg
		}
	} else if reqType == cns.CnsWasmInstantiate {
		c.InstantiateMsg = contract.DefaultCnsContractInstantiateMsg()

		msg, err := c.MakeInstantiateMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.InstantiateMsgSendChan <- msg
		}
	} else if reqType == cns.WasmExecute {
		msg, err := c.MsgConvert(body, reqType).MakeExecuteMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.ExecuteMsgSendChan <- msg
		}
	} else if reqType == cns.CnsWasmExecute {
		c.ExecuteMsg = contract.CnsContractExecuteMsg(body)

		msg, err := c.MakeExecuteMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.ExecuteMsgSendChan <- msg
		}
	} else if reqType == cns.WasmQuery {
		msg, err := c.MsgConvert(body, reqType).MakeQueryMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.QueryMsgSendChan <- msg
		}
	} else if reqType == cns.CnsWasmQueryByDomain {
		c.QueryMsg = contract.CnsContractQueryDomainMsg(body)

		msg, err := c.MakeQueryMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.QueryMsgSendChan <- msg
		}
	} else if reqType == cns.CnsWasmQueryByAccount {
		c.QueryMsg = contract.CnsContractQueryAccountMsg(body)

		msg, err := c.MakeQueryMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.QueryMsgSendChan <- msg
		}
	} else if reqType == cns.WasmListCode {
		msg, err := c.MakeListcodeMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.ListcodeMsgSendChan <- msg
		}
	} else if reqType == cns.WasmListContractByCode {
		msg, err := c.MsgConvert(body, reqType).MakeListContractByCodeMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.ListContractByCodeMsgSendChan <- msg
		}
	} else if reqType == cns.WasmDownload {
		msg, err := c.MsgConvert(body, reqType).MakeDownloadMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.DownloadMsgSendChan <- msg
		}
	} else if reqType == cns.WasmCodeInfo {
		msg, err := c.MsgConvert(body, reqType).MakeCodeInfoMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.CodeInfoMsgSendChan <- msg
		}
	} else if reqType == cns.WasmContractInfo {
		msg, err := c.MsgConvert(body, reqType).MakeContractInfoMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.ContractInfoMsgSendChan <- msg
		}
	} else if reqType == cns.WasmContractStateAll {
		msg, err := c.MsgConvert(body, reqType).MakeContractStateAllMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.ContractStateAllMsgSendChan <- msg
		}
	} else if reqType == cns.WasmContractHistory {
		msg, err := c.MsgConvert(body, reqType).MakeContractHistoryMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.ContractHistoryMsgSendChan <- msg
		}
	} else if reqType == cns.WasmPinned {
		msg, err := c.MakePinnedMsg()
		if err != nil {
			httpParseErrReturn(err, response, channel)
		} else {
			channel.PinnedMsgSendChan <- msg
		}
	}
}

func checkRequest(request *http.Request) {
	util.LogHttpServer("Client IP addr : " + request.RemoteAddr)
	util.LogHttpServer("Request API : " + request.URL.Path)
}

func httpParseErrReturn(err error,
	response cns.HttpResponseStruct,
	channel types.ChannelStruct) {

	util.LogErr("ERROR, ", err)
	response = util.HttpResponseTypeStruct(106, "", err.Error())
	channel.ErrorChan <- response
}
