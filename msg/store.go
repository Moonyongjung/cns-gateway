package msg

import (
	"github.com/Moonyongjung/cns-gw/msg/parse"
	"github.com/Moonyongjung/cns-gw/key"
	"github.com/Moonyongjung/cns-gw/util"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeStoreMsg() (wasm.MsgStoreCode, error) {	
	contractPath := util.GetConfig().Get("contractPath")
	contractName := util.GetConfig().Get("contractName")
	gwAdd := util.GetAddrByPrivKeyArmor(key.GwKey().GetPriKey())
	
	msg, err := parse.ParseStoreCodeArgs(contractPath + contractName, gwAdd)
	if err != nil {
		util.LogGw(err)
		return wasm.MsgStoreCode{}, err
	}

	if err = msg.ValidateBasic(); err != nil {
		util.LogGw(err)
		return wasm.MsgStoreCode{}, err
	}

	return msg, nil
}