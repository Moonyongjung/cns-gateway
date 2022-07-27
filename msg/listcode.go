package msg

import (
	"github.com/Moonyongjung/cns-gw/msg/parse"	

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeListcodeMsg() (wasm.QueryCodesRequest, error) {		
	msg := parse.ParseListcodeArgs()
	return msg, nil
}