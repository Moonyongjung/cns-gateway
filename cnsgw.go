package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Moonyongjung/cns-gw/bc"
	"github.com/Moonyongjung/cns-gw/client"
	"github.com/Moonyongjung/cns-gw/gw"
	"github.com/Moonyongjung/cns-gw/key"
	cnsdb "github.com/Moonyongjung/cns-gw/db"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	cmclient "github.com/cosmos/cosmos-sdk/client"

)

var configPath = "./config/config.json"
var configKeyPath = "./config/configKey.json"
var configDbPath = "./config/configDb.json"

var Channel cns.ChannelStruct
var clientCtx cmclient.Context

func init() {
	util.GetConfig().Read(configPath)
	util.GetConfig().Read(configKeyPath)
	util.GetConfig().Read(configDbPath)
	util.SetChainPrefixConfig()	

	clientCtx = client.SetClient()
	key.KeyInit(clientCtx)
}

func main() {
	channel := cns.ChannelInit()
	db := cnsdb.DbInit()
	
	go bc.TxInit(channel, clientCtx)
	go gw.HttpServerInit(channel, db)
	
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	util.LogGw("Shutting down the server...")
	util.LogGw("Server gracefully stopped")	
}