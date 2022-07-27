package gw

import (
	"strings"

	"github.com/Moonyongjung/cns-gw/bc"
	"github.com/Moonyongjung/cns-gw/bc/rest"
	"github.com/Moonyongjung/cns-gw/client"
	"github.com/Moonyongjung/cns-gw/contract"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/msg"
	"github.com/Moonyongjung/cns-gw/util"

	"github.com/mitchellh/mapstructure"
)

func DomainMapping(userAccount string) []byte {	

	if userAccount != "" {
		isExist := alreadyAccountMappingCheck([]byte(userAccount))
		if isExist {
			responseByte := httpResponseByte(108, "")		
			return responseByte

		} else {
			responseByte := httpResponseByte(0, userAccount)
			return responseByte
		}
		
	} else {
		responseData := httpResponseByte(109, "")		
		return responseData
	}
}

func DomainConfirm(requestData string) []byte {
	data := strings.Split(requestData, ",")
	userAccount := data[0]
	domainName := data[1]

	isExist := domainExistCheck([]byte(domainName))
	if isExist {
		responseByte := httpResponseByte(107, "")
		return responseByte

	} else {

		gatewayServerPort := util.GetConfig().Get("gatewayServerPort")		
		
		util.LogGw(userAccount)
		util.LogGw(domainName)
	
		url := "http://localhost:" + gatewayServerPort + "/api/wasm/cns-execute"
		domainMappingRequest := cns.DomainMappingRequest{
			DomainName: domainName,
			AccountAddress: userAccount,
		}
		jsonData, err := util.JsonMarshalData(domainMappingRequest)
		if err != nil {
			util.LogGw(err)
		}
	
		responseBody := rest.HttpClient("POST", url, jsonData)
		util.LogGw(string(responseBody))	
	
		var httpResponseStruct cns.HttpResponseStruct
		httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
		mapstructure.Decode(httpResponseStructData, &httpResponseStruct)

		responseByte, err := util.JsonMarshalData(httpResponseStruct)
		if err != nil {
			util.LogGw(err)
		}			
	
		return responseByte
	}
}

func SendInquiry(requestData string, pkArmor string) []byte {
	data := strings.Split(requestData, ",")
	domainName := data[0]
	amount := data[1]

	isExist := domainExistCheck([]byte(domainName))
	if !isExist {
		responseByte := httpResponseByte(111, "")
		return responseByte

	} else {
		gatewayServerPort := util.GetConfig().Get("gatewayServerPort")		
		
		
		util.LogGw(domainName)
		util.LogGw(amount)
	
		url := "http://localhost:" + gatewayServerPort + "/api/wasm/cns-query-by-domain"
		domainMappingRequest := cns.DomainMappingRequest{
			DomainName: domainName,
			// AccountAddress: userAccount,
		}
		jsonData, err := util.JsonMarshalData(domainMappingRequest)
		if err != nil {
			util.LogGw(err)
		}
	
		responseBody := rest.HttpClient("POST", url, jsonData)
		util.LogGw("Send inquiry, domain check :", string(responseBody))	
	
		var httpResponseStruct cns.HttpResponseStruct
		httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
		mapstructure.Decode(httpResponseStructData, &httpResponseStruct)

		if httpResponseStruct.ResCode != 0 {
			responseByte := httpResponseByte(111, "")		
			return responseByte
		} else {
			userClient := client.SetClientUser()
			pk := util.GetPriKeyByArmor(pkArmor)
			fromAddr := util.GetAddrByPrivKey(pk)
			accNum, _, _ := rest.GetAccountInfoHttpClient(fromAddr.String())			

			splitData := strings.Split(httpResponseStruct.ResData, ":")
			convertStr := strings.Replace(splitData[2], "\"", "", -1)
			toAddr := strings.Replace(convertStr, "}", "", -1)

			bankMsgStruct := cns.BankMsgStruct{
				FromAddress: fromAddr.String(),
				ToAddress: toAddr,
				Amount: amount,
			}						

			msg, err := msg.MakeBankMsg(bankMsgStruct)		
			if err != nil {
				util.LogGw(err)
			}

			_, httpResponse := bc.TxCreate(pk, accNum, userClient, msg, "bank")
			
			responseByte, err := util.JsonMarshalData(httpResponse)
			if err != nil {
				util.LogGw(err)
			}

			return responseByte			
		}
	}
}

func domainExistCheck(requestData []byte) bool {
	gatewayServerPort := util.GetConfig().Get("gatewayServerPort")
	url := "http://localhost:" + gatewayServerPort + "/api/wasm/query"

	domainMappingRequest := cns.DomainMappingRequest{
		DomainName: string(requestData),
	}

	jsonData, err := util.JsonMarshalData(domainMappingRequest)
	if err != nil {
		util.LogGw(err)
	}

	queryMsgData := contract.CnsContractQueryDomainMsg(jsonData)
	queryMsgDataJson, err := util.JsonMarshalData(queryMsgData)
	if err != nil {
		util.LogGw(err)
	}

	util.LogGw(queryMsgData.ContractAddress)
	util.LogGw(queryMsgData.QueryMsg)

	responseBody := rest.HttpClient("POST", url, queryMsgDataJson)
	util.LogGw("Domain exist check : ", string(responseBody))

	var httpResponseStruct cns.HttpResponseStruct
	httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
	mapstructure.Decode(httpResponseStructData, &httpResponseStruct)	

	if httpResponseStruct.ResCode == 0 {
		return true
	} else {
		return false
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
		util.LogGw(err)
	}

	queryMsgData := contract.CnsContractQueryAccountMsg(jsonData)
	queryMsgDataJson, err := util.JsonMarshalData(queryMsgData)
	if err != nil {
		util.LogGw(err)
	}

	responseBody := rest.HttpClient("POST", url, queryMsgDataJson)
	util.LogGw("Already account mapped check : ", string(responseBody))

	var httpResponseStruct cns.HttpResponseStruct
	httpResponseStructData := util.JsonUnmarshalData(httpResponseStruct, responseBody)
	mapstructure.Decode(httpResponseStructData, &httpResponseStruct)

	if httpResponseStruct.ResCode == 0 {
		return true
	} else {
		return false
	}
}