package contract

import (
	"github.com/Moonyongjung/cns-gw/msg"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
)

var cnsClientParse msg.CnsClientParse

func StoreContractCode(typeName string, codeId string) {
	jsonData := cns.ContractCodeJsonStruct{
		TypeName: typeName,
		CodeId:   codeId,
	}

	f := "./contract/info/contractCode.json"

	util.JsonMarshal(jsonData, f)
}

func StoreContractAddress(typeName string, contractAddr string, codeType string, codeId string) {

	jsonData := cns.ContractAddressJsonStruct{
		TypeName:        typeName,
		ContractAddress: contractAddr,
		CodeType:        codeType,
		CodeId:          codeId,
	}

	f := "./contract/info/contractAddress.json"

	util.JsonMarshal(jsonData, f)
}

//-- Only for CNS, create default instantiate message
func DefaultCnsContractInstantiateMsg() cns.InstantiateMsgStruct {
	denom := util.GetConfig().Get("denom")

	instantiateMsgData := cns.InstantiateMsgStruct{
		CodeId:  cnsClientParse.CnsContractConv("code").ContractCodeJsonStruct.CodeId,
		Amount:  "0" + denom,
		Label:   "contract inst",
		InitMsg: "{\"owner\":\"" + cns.GatewayAccount + "\"}",
	}

	return instantiateMsgData
}

//-- Only for CNS, create execute message
func CnsContractExecuteMsg(body []byte) cns.ExecuteMsgStruct {
	denom := util.GetConfig().Get("denom")

	c := cnsClientParse.CnsContractConv("address").DomainMapConv(body)

	executeMsgData := cns.ExecuteMsgStruct{
		ContractAddress: c.ContractAddressJsonStruct.ContractAddress,
		Amount:          "0" + denom,
		ExecMsg: "{\"save_domain_address_mapping\":{\"domain\":\"" + c.DomainMappingRequest.DomainName +
			"\",\"account_address\":\"" + c.DomainMappingRequest.AccountAddress + "\"}}",
	}

	return executeMsgData
}

//-- Only for CNS, query message
func CnsContractQueryDomainMsg(body []byte) cns.QueryMsgStruct {

	c := cnsClientParse.CnsContractConv("address").DomainMapConv(body)

	queryMsgData := cns.QueryMsgStruct{
		ContractAddress: c.ContractAddressJsonStruct.ContractAddress,
		QueryMsg:        "{\"domain_mapping\":{\"domain\":\"" + c.DomainMappingRequest.DomainName + "\"}}",
	}

	return queryMsgData
}

//-- Only for CNS, query message
func CnsContractQueryAccountMsg(body []byte) cns.QueryMsgStruct {

	c := cnsClientParse.CnsContractConv("address").DomainMapConv(body)

	queryMsgData := cns.QueryMsgStruct{
		ContractAddress: c.ContractAddressJsonStruct.ContractAddress,
		QueryMsg:        "{\"account_mapping\":{\"account_address\":\"" + c.DomainMappingRequest.AccountAddress + "\"}}"}

	return queryMsgData
}
