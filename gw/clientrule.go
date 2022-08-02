package gw

import (
	"strings"

	"github.com/Moonyongjung/cns-gw/bc"
	"github.com/Moonyongjung/cns-gw/bc/rest"
	"github.com/Moonyongjung/cns-gw/client"
	"github.com/Moonyongjung/cns-gw/contract"
	"github.com/Moonyongjung/cns-gw/msg"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
)

func DomainMapping(userAccount string) []byte {

	if userAccount != "" {
		isExist := alreadyAccountMappingCheck([]byte(userAccount))
		if isExist {
			responseByte := util.HttpResponseByte(108, "", "")
			return responseByte

		} else {
			responseByte := util.HttpResponseByte(0, "", userAccount)
			return responseByte
		}

	} else {
		responseByte := util.HttpResponseByte(109, "", "")
		return responseByte
	}
}

func DomainConfirm(requestData string) []byte {
	data := strings.Split(requestData, ",")
	userAccount := data[0]
	domainName := data[1]

	isExist, _ := domainExistCheck([]byte(domainName))
	if isExist {
		responseByte := util.HttpResponseByte(107, "", "")
		return responseByte

	} else {
		gatewayServerPort := util.GetConfig().Get("gatewayServerPort")
		url := "http://localhost:" + gatewayServerPort + "/api/wasm/cns-execute"
		domainMappingRequest := cns.DomainMappingRequest{
			DomainName:     domainName,
			AccountAddress: userAccount,
		}
		jsonData, err := util.JsonMarshalData(domainMappingRequest)
		if err != nil {
			util.LogErr(err)
		}

		responseBody := rest.HttpClient("POST", url, jsonData)
		util.LogGw(string(responseBody))

		return responseBody
	}
}

func SendInquiry(requestData string, pkArmor string) []byte {
	data := strings.Split(requestData, ",")
	domainName := data[0]
	amount := data[1]

	isExist, domainQueryHttpResponse := domainExistCheck([]byte(domainName))
	if !isExist {
		responseByte := util.HttpResponseByte(111, "", "")
		return responseByte

	} else {
		if domainQueryHttpResponse.ResCode != 0 {
			responseByte := util.HttpResponseByte(111, "", "")
			return responseByte
		} else {
			userClient := client.SetClientUser()
			pk := util.GetPriKeyByArmor(pkArmor)
			fromAddr := util.GetAddrByPrivKey(pk)
			accNum, _, _ := rest.GetAccountInfoHttpClient(fromAddr.String())
			if accNum == "" {
				return util.HttpResponseByte(112, "", "")
			}
			var contractMsg msg.ContractMsg

			splitData := strings.Split(domainQueryHttpResponse.ResData, ":")
			convertStr := strings.Replace(splitData[2], "\"", "", -1)
			toAddr := strings.Replace(convertStr, "}", "", -1)

			bankMsgStruct := cns.BankMsgStruct{
				FromAddress: fromAddr.String(),
				ToAddress:   toAddr,
				Amount:      amount,
			}

			contractMsg.BankMsg = bankMsgStruct
			msg, err := contractMsg.MakeBankMsg()
			if err != nil {
				util.LogErr(err)
			}
			_, httpResponse := bc.TxCreate(pk, accNum, userClient, msg, "bank")
			responseByte := util.HttpResponseByte(httpResponse.ResCode, httpResponse.ResMsg, httpResponse.ResData)

			return responseByte
		}
	}
}

func domainExistCheck(requestData []byte) (bool, cns.HttpResponseStruct) {
	gatewayServerPort := util.GetConfig().Get("gatewayServerPort")
	url := "http://localhost:" + gatewayServerPort + "/api/wasm/query"

	domainMappingRequest := cns.DomainMappingRequest{
		DomainName: string(requestData),
	}

	jsonData, err := util.JsonMarshalData(domainMappingRequest)
	if err != nil {
		util.LogErr(err)
	}

	queryMsgData := contract.CnsContractQueryDomainMsg(jsonData)
	queryMsgDataJson, err := util.JsonMarshalData(queryMsgData)
	if err != nil {
		util.LogErr(err)
	}

	responseBody := rest.HttpClient("POST", url, queryMsgDataJson)
	util.LogGw("Domain exist check : ", string(responseBody))

	httpResponseStruct := util.HttpResonseByteToStruct(responseBody)

	if httpResponseStruct.ResCode == 0 {
		return true, httpResponseStruct
	} else {
		return false, cns.HttpResponseStruct{}
	}
}

func alreadyAccountMappingCheck(requestData []byte) bool {
	gatewayServerPort := util.GetConfig().Get("gatewayServerPort")
	url := "http://localhost:" + gatewayServerPort + "/api/wasm/query"

	domainMappingRequest := cns.DomainMappingRequest{
		AccountAddress: string(requestData),
	}

	jsonData, err := util.JsonMarshalData(domainMappingRequest)
	if err != nil {
		util.LogErr(err)
	}

	queryMsgData := contract.CnsContractQueryAccountMsg(jsonData)
	queryMsgDataJson, err := util.JsonMarshalData(queryMsgData)
	if err != nil {
		util.LogErr(err)
	}

	responseBody := rest.HttpClient("POST", url, queryMsgDataJson)
	util.LogGw("Already account mapped check : ", string(responseBody))

	httpResponseStruct := util.HttpResonseByteToStruct(responseBody)

	if httpResponseStruct.ResCode == 0 {
		return true
	} else {
		return false
	}
}
