package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	cns "github.com/Moonyongjung/cns-gw/types"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
)

func MakeCodeInfoMsg(codeInfoMsgData cns.CodeInfoMsgStruct) (wasm.QueryCodeRequest, error) {		
	
	if (cns.CodeInfoMsgStruct{}) == codeInfoMsgData {
		return wasm.QueryCodeRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseCodeInfoArgs(codeInfoMsgData)
	return msg, nil
}