package util

import (
	cns "github.com/Moonyongjung/cns-gw/types"
)

//-- Response code
//   If response code is not included in below response code, it is Cosmos SDK or Wasm code

func HttpResponseInit() {
	cns.ResponseType = make(map[int]string)

	cns.ResponseType[0] = "Success"
	cns.ResponseType[102] = "Transaction signing error"
	cns.ResponseType[103] = "Transaction encoding error"
	cns.ResponseType[104] = "Transaction broadcast error"
	cns.ResponseType[105] = "Query error"
	cns.ResponseType[106] = "Message creating error"
	cns.ResponseType[107] = "Exist Domain, select another domain name"
	cns.ResponseType[108] = "This account already has domain"
	cns.ResponseType[109] = "No user Account"
	cns.ResponseType[110] = "No session, input mnemonic"
	cns.ResponseType[111] = "No domain name"
}

func HttpResponseByte(resCode int, resMsg string, resData string) []byte {
	httpResponse := HttpResponseTypeStruct(resCode, resMsg, resData)
	responseByte, err := JsonMarshalData(httpResponse)
	if err != nil {
		LogErr(err)
	}

	return responseByte
}

func HttpResponseTypeStruct(resCode int, resMsg string, resData string) cns.HttpResponseStruct {
	var httpResponse cns.HttpResponseStruct

	httpResponse.ResCode = resCode
	if resMsg == "" {
		httpResponse.ResMsg = cns.ResponseType[resCode]
	} else {
		httpResponse.ResMsg = resMsg
	}

	httpResponse.ResData = resData

	return httpResponse
}
