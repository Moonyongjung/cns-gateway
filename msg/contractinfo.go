package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	cns "github.com/Moonyongjung/cns-gw/types"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeContractInfoMsg(contractInfoMsgData cns.ContractInfoMsgStruct) (wasm.QueryContractInfoRequest, error) {		
	
	if (cns.ContractInfoMsgStruct{}) == contractInfoMsgData {
		return wasm.QueryContractInfoRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseContractInfoArgs(contractInfoMsgData)
	return msg, nil
}