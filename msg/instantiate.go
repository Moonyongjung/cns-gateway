package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	"github.com/Moonyongjung/cns-gw/key"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeInstantiateMsg(instantiateMsgData cns.InstantiateMsgSturct) (wasm.MsgInstantiateContract, error) {	
	gwAdd := util.GetAddrByPrivKeyArmor(key.GwKey().GetPriKey())
	
	if (cns.InstantiateMsgSturct{}) == instantiateMsgData {
		return wasm.MsgInstantiateContract{}, errors.New("Empty request or type of parameter is not correct")
	}
	
	msg, err := parse.ParseInstantiateArgs(instantiateMsgData, gwAdd)
	if err != nil {
		util.LogGw(err)
		return wasm.MsgInstantiateContract{}, err
	}

	if err = msg.ValidateBasic(); err != nil {
		util.LogGw(err)
		return wasm.MsgInstantiateContract{}, err
	}

	return msg, nil
}