package gw

import (
	"database/sql"

	cns "github.com/Moonyongjung/cns-gw/types"
)

func HttpServerInit(channel cns.ChannelStruct, db *sql.DB) {
	httpResponseInit()
	go RunHttpServer(channel, db)
	for {
		select {
			case txRes := <- channel.TxResponseChan:
				res := httpResponseByte(int(txRes.Code), txRes.RawLog)
				channel.HttpServerChan <- res

			case queryRes := <- channel.QueryResponseChan:				
				res := httpResponseByte(0, queryRes)
				channel.HttpServerChan <- res
			
			case err := <- channel.ErrorChan:
				res := httpResponseByte(err.ResCode, err.ResData)
				channel.HttpServerChan <- res
		}
	}
}