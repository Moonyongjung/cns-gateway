package gw

import (
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
)

func HttpServerInit(channel cns.ChannelStruct) {
	util.HttpResponseInit()
	go RunHttpServer(channel)
	for {
		select {
		case txRes := <-channel.TxResponseChan:
			res := util.HttpResponseByte(int(txRes.Code), txRes.RawLog, txRes.Data)
			channel.HttpServerChan <- res

		case queryRes := <-channel.QueryResponseChan:
			res := util.HttpResponseByte(0, "", queryRes)
			channel.HttpServerChan <- res

		case err := <-channel.ErrorChan:
			res := util.HttpResponseByte(err.ResCode, err.ResMsg, err.ResData)
			channel.HttpServerChan <- res
		}
	}
}
