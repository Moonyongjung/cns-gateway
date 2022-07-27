package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	cns "github.com/Moonyongjung/cns-gw/types"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeContractHistoryMsg(contractHistoryMsgData cns.ContractHistoryMsgStruct) (wasm.QueryContractHistoryRequest, error) {		
	
	if (cns.ContractHistoryMsgStruct{}) == contractHistoryMsgData {
		return wasm.QueryContractHistoryRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseContractHistoryArgs(contractHistoryMsgData)
	return msg, nil
}