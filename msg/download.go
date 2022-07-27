package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/msg/parse"
	cns "github.com/Moonyongjung/cns-gw/types"
)

func MakeDownloadMsg(downloadMsgData cns.DownloadMsgStruct) ([]interface{}, error) {		
	var msgInterfaceSlice []interface{}
	if (cns.DownloadMsgStruct{}) == downloadMsgData {
		return nil, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseDownloadArgs(downloadMsgData)
	msgInterfaceSlice = append(msgInterfaceSlice, msg)
	msgInterfaceSlice = append(msgInterfaceSlice, downloadMsgData.DownloadFileName)
	return msgInterfaceSlice, nil
}