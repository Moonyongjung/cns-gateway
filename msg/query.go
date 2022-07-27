package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	"github.com/Moonyongjung/cns-gw/key"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeQueryMsg(queryMsgData cns.QueryMsgStruct) (wasm.QuerySmartContractStateRequest, error) {	
	gwAdd := util.GetAddrByPrivKeyArmor(key.GwKey().GetPriKey())
	
	if (cns.QueryMsgStruct{}) == queryMsgData {
		return wasm.QuerySmartContractStateRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	
	msg, err := parse.ParseQueryArgs(queryMsgData, gwAdd)
	if err != nil {
		util.LogGw(err)
		return wasm.QuerySmartContractStateRequest{}, err
	}

	return msg, nil
}