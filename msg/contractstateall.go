package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	cns "github.com/Moonyongjung/cns-gw/types"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeContractStateAllMsg(contractStateAllMsgData cns.ContractStateAllMsgStruct) (wasm.QueryAllContractStateRequest, error) {		
	
	if (cns.ContractStateAllMsgStruct{}) == contractStateAllMsgData {
		return wasm.QueryAllContractStateRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseContractStateAllArgs(contractStateAllMsgData)
	return msg, nil
}