package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	"github.com/Moonyongjung/cns-gw/key"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeExecuteMsg(executeMsgData cns.ExecuteMsgStruct) (wasm.MsgExecuteContract, error) {	
	gwAdd := util.GetAddrByPrivKeyArmor(key.GwKey().GetPriKey())
	
	if (cns.ExecuteMsgStruct{}) == executeMsgData {
		return wasm.MsgExecuteContract{}, errors.New("Empty request or type of parameter is not correct")
	}
	
	msg, err := parse.ParseExecuteArgs(executeMsgData, gwAdd)
	if err != nil {
		util.LogGw(err)
		return wasm.MsgExecuteContract{}, err
	}

	if err = msg.ValidateBasic(); err != nil {
		util.LogGw(err)
		return wasm.MsgExecuteContract{}, err
	}

	return msg, nil
}