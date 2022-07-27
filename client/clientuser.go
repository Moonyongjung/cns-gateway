package client

import (

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	"github.com/evmos/ethermint/crypto/hd"	
	cmclient "github.com/cosmos/cosmos-sdk/client"
)

func SetClientUser() cmclient.Context {
	rpcEndpoint := util.GetConfig().Get("rpcEndpoint")
	chainId := util.GetConfig().Get("chainId")
	encodingConfig := cns.MakeEncodingConfigEth()
		
	clientCtx := cmclient.Context{}

	//-- for using resolve wasm api, need to wasm's txconfig
	//-- mod : using ethermint's txconfig
	clientCtx = clientCtx.WithTxConfig(encodingConfig.TxConfig)
	clientCtx = clientCtx.WithChainID(chainId)
	clientCtx = clientCtx.WithCodec(encodingConfig.Marshaler)
	clientCtx = clientCtx.WithLegacyAmino(encodingConfig.Amino)
	clientCtx = clientCtx.WithInterfaceRegistry(encodingConfig.InterfaceRegistry)
	clientCtx = clientCtx.WithNodeURI(rpcEndpoint)
	clientCtx = clientCtx.WithKeyringOptions(hd.EthSecp256k1Option())

	client, _ := cmclient.NewClientFromNode(rpcEndpoint)
	clientCtx = clientCtx.WithClient(client)

	//-- To check code ID of contract, broadcast mode = block
	clientCtx = clientCtx.WithBroadcastMode("block")	
	
	return clientCtx
}
