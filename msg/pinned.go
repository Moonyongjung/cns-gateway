package msg

import (
	"github.com/Moonyongjung/cns-gw/msg/parse"	

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakePinnedMsg() (wasm.QueryPinnedCodesRequest, error) {		
	msg := parse.ParsePinnedArgs()
	return msg, nil
}