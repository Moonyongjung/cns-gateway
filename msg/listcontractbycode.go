package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	cns "github.com/Moonyongjung/cns-gw/types"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeListContractByCodeMsg(listContractByCodeMsgData cns.ListContractByCodeMsgStruct) (wasm.QueryContractsByCodeRequest, error) {		
	
	if (cns.ListContractByCodeMsgStruct{}) == listContractByCodeMsgData {
		return wasm.QueryContractsByCodeRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseListContractByCodeArgs(listContractByCodeMsgData)
	return msg, nil
}