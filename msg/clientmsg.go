package msg

import (
	"net/http"

	cnsdb "github.com/Moonyongjung/cns-gw/db"
	"github.com/Moonyongjung/cns-gw/key"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
	"github.com/mitchellh/mapstructure"
)

type CnsClientParse struct {
	CreateMnemonicRes         cns.CreateMnemonicResponse
	AddressRes                cns.AddressResponse
	DomainMappingRequest      cns.DomainMappingRequest
	ContractCodeJsonStruct    cns.ContractCodeJsonStruct
	ContractAddressJsonStruct cns.ContractAddressJsonStruct
}

func (c CnsClientParse) ResList(
	reqType string,
	w http.ResponseWriter,
	body []byte) (http.ResponseWriter, []byte) {

	if reqType == cns.WalletCreate {
		c.CreateMnemonicRes.Mnemonic = key.UserNewMnemonic()
		responseData, err := util.JsonMarshalData(c.CreateMnemonicRes)
		if err != nil {
			util.LogErr(err)
		}

		return w, responseData
	} else if reqType == cns.WalletAddress {
		db := cns.Db
		util.LogGw("From client request body : ", string(body))

		resMsg, redirectPage, pk := key.UserNewKey(string(body))
		c.AddressRes.ResMsg = resMsg
		c.AddressRes.DirectPage = redirectPage
		w = cnsdb.UserKeySessionDb(w, db, pk)

		responseData, err := util.JsonMarshalData(c.AddressRes)
		if err != nil {
			util.LogErr(err)
		}

		return w, responseData
	}

	return w, nil
}

func (c CnsClientParse) CnsContractConv(reqType string) CnsClientParse {
	if reqType == "code" {
		contractCodeData := util.JsonUnmarshal(c.ContractCodeJsonStruct, "./contract/info/contractCode.json")
		mapstructure.Decode(contractCodeData, &c.ContractCodeJsonStruct)
	} else {
		contractAddressData := util.JsonUnmarshal(c.ContractAddressJsonStruct, "./contract/info/contractAddress.json")
		mapstructure.Decode(contractAddressData, &c.ContractAddressJsonStruct)
	}

	return c
}

func (c CnsClientParse) DomainMapConv(body []byte) CnsClientParse {

	domainMappingRequestData := util.JsonUnmarshalData(c.DomainMappingRequest, body)
	mapstructure.Decode(domainMappingRequestData, &c.DomainMappingRequest)

	return c
}
