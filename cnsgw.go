package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Moonyongjung/cns-gw/bc"
	"github.com/Moonyongjung/cns-gw/client"
	cnsdb "github.com/Moonyongjung/cns-gw/db"
	"github.com/Moonyongjung/cns-gw/gw"
	"github.com/Moonyongjung/cns-gw/key"
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
	cnsdb.DbInit()
}

func main() {
	channel := cns.ChannelInit()

	go bc.TxInit(channel, clientCtx)
	go gw.HttpServerInit(channel)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	util.LogGw("Shutting down the server...")
	util.LogGw("Server gracefully stopped")
}
