package contract

import (
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
	"github.com/mitchellh/mapstructure"
)

//-- Saved contract code ID format
type contractCodeJsonStruct struct {
	TypeName string
	CodeId string
}

//-- Save contract address format
type contractAddressJsonStruct struct {
	TypeName string
	ContractAddress string
	CodeType string
	CodeId string
}

func StoreContractCode(typeName string, codeId string) {
	jsonData := contractCodeJsonStruct{
		TypeName: typeName,
		CodeId: codeId,
	}

	f := "./contract/info/contractCode.json"

	util.JsonMarshal(jsonData, f)
}

func StoreContractAddress(typeName string, contractAddr string, codeType string, codeId string) {

	jsonData := contractAddressJsonStruct{
		TypeName: typeName,
		ContractAddress: contractAddr,
		CodeType: codeType,
		CodeId: codeId,
	}
	
	f := "./contract/info/contractAddress.json"
	
	util.JsonMarshal(jsonData, f)
}

//-- Only for CNS, create default instantiate message
func DefaultCnsContractInstantiateMsg() cns.InstantiateMsgSturct{
	denom := util.GetConfig().Get("denom")

	var contractCodeJsonStruct contractCodeJsonStruct
	contractCodeData := util.JsonUnmarshal(contractCodeJsonStruct, "./contract/info/contractCode.json")
	mapstructure.Decode(contractCodeData, &contractCodeJsonStruct)

	instantiateMsgData := cns.InstantiateMsgSturct{
		CodeId: contractCodeJsonStruct.CodeId,
		Amount: "0" + denom,
		Label: "contract inst",
		InitMsg: "{\"owner\":\"" + cns.GatewayAccount + "\"}",
	}

	return instantiateMsgData
}

//-- Only for CNS, create execute message
func CnsContractExecuteMsg(body []byte) cns.ExecuteMsgStruct {
	denom := util.GetConfig().Get("denom")

	var contractAddressJsonStruct contractAddressJsonStruct
	contractAddressData := util.JsonUnmarshal(contractAddressJsonStruct, "./contract/info/contractAddress.json")
	mapstructure.Decode(contractAddressData, &contractAddressJsonStruct)

	var domainMappingRequest cns.DomainMappingRequest
	domainMappingRequestData := util.JsonUnmarshalData(domainMappingRequest, body)
	mapstructure.Decode(domainMappingRequestData, &domainMappingRequest)

	executeMsgData := cns.ExecuteMsgStruct{
		ContractAddress: contractAddressJsonStruct.ContractAddress,
		Amount: "0" + denom,
		ExecMsg: "{\"save_domain_address_mapping\":{\"domain\":\"" + domainMappingRequest.DomainName +
		 "\",\"account_address\":\""+ domainMappingRequest.AccountAddress +"\"}}",
	}

	return executeMsgData
}

//-- Only for CNS, query message
func CnsContractQueryDomainMsg(body []byte) cns.QueryMsgStruct {	

	var contractAddressJsonStruct contractAddressJsonStruct
	contractAddressData := util.JsonUnmarshal(contractAddressJsonStruct, "./contract/info/contractAddress.json")
	mapstructure.Decode(contractAddressData, &contractAddressJsonStruct)

	var domainMappingRequest cns.DomainMappingRequest
	domainMappingRequestData := util.JsonUnmarshalData(domainMappingRequest, body)
	mapstructure.Decode(domainMappingRequestData, &domainMappingRequest)	

	queryMsgData := cns.QueryMsgStruct{
		ContractAddress: contractAddressJsonStruct.ContractAddress,		
		QueryMsg: "{\"domain_mapping\":{\"domain\":\"" + domainMappingRequest.DomainName +"\"}}",
	}

	return queryMsgData
}

//-- Only for CNS, query message
func CnsContractQueryAccountMsg(body []byte) cns.QueryMsgStruct {	

	var contractAddressJsonStruct contractAddressJsonStruct
	contractAddressData := util.JsonUnmarshal(contractAddressJsonStruct, "./contract/info/contractAddress.json")
	mapstructure.Decode(contractAddressData, &contractAddressJsonStruct)

	var domainMappingRequest cns.DomainMappingRequest
	domainMappingRequestData := util.JsonUnmarshalData(domainMappingRequest, body)
	mapstructure.Decode(domainMappingRequestData, &domainMappingRequest)

	queryMsgData := cns.QueryMsgStruct{
		ContractAddress: contractAddressJsonStruct.ContractAddress,		
		QueryMsg: "{\"account_mapping\":{\"account_address\":\"" + domainMappingRequest.AccountAddress +"\"}}",
	}

	return queryMsgData
}