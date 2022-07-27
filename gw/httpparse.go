package gw

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Moonyongjung/cns-gw/msg"
	"github.com/Moonyongjung/cns-gw/types"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
	"github.com/Moonyongjung/cns-gw/contract"

	"github.com/mitchellh/mapstructure"
)

func doTransactionbyType(request *http.Request, channel cns.ChannelStruct) {
	var body []byte
	var response cns.HttpResponseStruct

	checkRequest(request)
	transactionType := strings.Split(request.URL.Path, "/")[2:]
	
	if request.Method == "POST" {
		bodyByte, err := ioutil.ReadAll(request.Body)
		if err != nil {
			httpParseErrReturn(err, response, channel)
		}
		body = bodyByte
	}

	if transactionType[0] == "bank" {
		if transactionType[1] == "send" {			
			var bankMsgStruct cns.BankMsgStruct
			bankMsgData := util.JsonUnmarshalData(bankMsgStruct, body)
			mapstructure.Decode(bankMsgData, &bankMsgStruct)			

			msg, err := msg.MakeBankMsg(bankMsgStruct)			
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.BankMsgSendChan <- msg
			}
		}
	} else if transactionType[0] == "wasm" {
		if transactionType[1] == "store"{	

			msg, err := msg.MakeStoreMsg()
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.StoreMsgSendChan <- msg
			}
		}
		if transactionType[1] == "instantiate"{	
			var instantiateMsgStruct cns.InstantiateMsgSturct
			instantiateMsgData := util.JsonUnmarshalData(instantiateMsgStruct, body)
			mapstructure.Decode(instantiateMsgData, &instantiateMsgStruct)

			msg, err := msg.MakeInstantiateMsg(instantiateMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.InstantiateMsgSendChan <- msg
			}
		}
		if transactionType[1] == "default-cns-instantiate"{	
			instantiateMsgStruct := contract.DefaultCnsContractInstantiateMsg()

			msg, err := msg.MakeInstantiateMsg(instantiateMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.InstantiateMsgSendChan <- msg
			}
		}
		if transactionType[1] == "execute" {
			var executeMsgStruct cns.ExecuteMsgStruct
			executeMsgData := util.JsonUnmarshalData(executeMsgStruct, body)
			mapstructure.Decode(executeMsgData, &executeMsgStruct)

			msg, err := msg.MakeExecuteMsg(executeMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.ExecuteMsgSendChan <- msg
			}
		}
		if transactionType[1] == "cns-execute" {	

			executeMsgStruct := contract.CnsContractExecuteMsg(body)

			msg, err := msg.MakeExecuteMsg(executeMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.ExecuteMsgSendChan <- msg
			}
		}
		if transactionType[1] == "query" {
			var queryMsgStruct cns.QueryMsgStruct
			queryMsgData := util.JsonUnmarshalData(queryMsgStruct, body)
			mapstructure.Decode(queryMsgData, &queryMsgStruct)

			msg, err := msg.MakeQueryMsg(queryMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.QueryMsgSendChan <- msg
			}
		}
		if transactionType[1] == "cns-query-by-domain" {	

			queryMsgStruct := contract.CnsContractQueryDomainMsg(body)

			msg, err := msg.MakeQueryMsg(queryMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.QueryMsgSendChan <- msg
			}
		}
		if transactionType[1] == "cns-query-by-account" {	

			queryMsgStruct := contract.CnsContractQueryAccountMsg(body)

			msg, err := msg.MakeQueryMsg(queryMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.QueryMsgSendChan <- msg
			}
		}
		if transactionType[1] == "list-code" {	
			msg, err := msg.MakeListcodeMsg()
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.ListcodeMsgSendChan <- msg
			}
		}
		if transactionType[1] == "list-contract-by-code" {	
			var listContractByCodeMsgStruct cns.ListContractByCodeMsgStruct
			listContractByCodeMsgData := util.JsonUnmarshalData(listContractByCodeMsgStruct, body)
			mapstructure.Decode(listContractByCodeMsgData, &listContractByCodeMsgStruct)

			msg, err := msg.MakeListContractByCodeMsg(listContractByCodeMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.ListContractByCodeMsgSendChan <- msg
			}
		}
		if transactionType[1] == "download" {	
			var downloadMsgStruct cns.DownloadMsgStruct
			downloadMsgData := util.JsonUnmarshalData(downloadMsgStruct, body)
			mapstructure.Decode(downloadMsgData, &downloadMsgStruct)

			msg, err := msg.MakeDownloadMsg(downloadMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.DownloadMsgSendChan <- msg
			}
		}
		if transactionType[1] == "code-info" {	
			var codeInfoMsgStruct cns.CodeInfoMsgStruct
			codeInfoMsgData := util.JsonUnmarshalData(codeInfoMsgStruct, body)
			mapstructure.Decode(codeInfoMsgData, &codeInfoMsgStruct)

			msg, err := msg.MakeCodeInfoMsg(codeInfoMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.CodeInfoMsgSendChan <- msg
			}
		}
		if transactionType[1] == "contract-info" {	
			var contractInfoMsgStruct cns.ContractInfoMsgStruct
			contractInfoMsgData := util.JsonUnmarshalData(contractInfoMsgStruct, body)
			mapstructure.Decode(contractInfoMsgData, &contractInfoMsgStruct)

			msg, err := msg.MakeContractInfoMsg(contractInfoMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.ContractInfoMsgSendChan <- msg
			}
		}
		if transactionType[1] == "contract-state-all" {	
			var contractStateAllMsgStruct cns.ContractStateAllMsgStruct
			contractStateAllMsgData := util.JsonUnmarshalData(contractStateAllMsgStruct, body)
			mapstructure.Decode(contractStateAllMsgData, &contractStateAllMsgStruct)

			msg, err := msg.MakeContractStateAllMsg(contractStateAllMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.ContractStateAllMsgSendChan <- msg
			}
		}
		if transactionType[1] == "contract-history" {	
			var contractHistoryMsgStruct cns.ContractHistoryMsgStruct
			contractHistoryMsgData := util.JsonUnmarshalData(contractHistoryMsgStruct, body)
			mapstructure.Decode(contractHistoryMsgData, &contractHistoryMsgStruct)

			msg, err := msg.MakeContractHistoryMsg(contractHistoryMsgStruct)
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.ContractHistoryMsgSendChan <- msg
			}
		}
		if transactionType[1] == "pinned" {			

			msg, err := msg.MakePinnedMsg()
			if err != nil {
				httpParseErrReturn(err, response, channel)
			} else {
				channel.PinnedMsgSendChan <- msg
			}
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

	util.LogGw("ERROR, ", err)
	response.ResCode = 106
	response.ResData = err.Error()
	channel.ErrorChan <- response
}