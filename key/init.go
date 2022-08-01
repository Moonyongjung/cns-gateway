package key

import (
	"io/ioutil"
	"os"

	"github.com/Moonyongjung/cns-gw/chain"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	cmclient "github.com/cosmos/cosmos-sdk/client"
)

var priKeyFileName = "pri.info"
var pubKeyFileName = "pub.info"
var mnemonicFileName = "mnemonic_backup.info"
var algoFileName = "keyalgo.info"

func KeyInit(clientCtx cmclient.Context) {
	keyStorePath := util.GetConfig().Get("keyStorePath")
	keyOwnerName := util.GetConfig().Get("keyOwnerName")
	keyOwnerPw := util.GetConfig().Get("keyOwnerPw")
	keyStoreFilePath := util.GetConfig().Get("keyStoreFilePath")

	if _, err := os.Stat(keyStorePath + keyStoreFilePath); os.IsNotExist(err) {
		err := os.Mkdir(keyStorePath+keyStoreFilePath, 0755)
		if err != nil {
			util.LogErr(err)
		}
		algo := chain.SelectChainType()
		chain.SaveChainType(keyStorePath+keyStoreFilePath, algo, algoFileName)
		clientCtx = storeKeyringBackend(keyOwnerName, keyStorePath+keyStoreFilePath, keyOwnerPw, algo, clientCtx)
		exportPrivKeyArmor(keyOwnerName, keyStorePath+keyStoreFilePath, keyOwnerPw, algo, clientCtx)
		cns.SelectKeyAlgorithm = algo

		util.LogGw("Key algorithm : ", cns.SelectKeyAlgorithm)

	} else {
		util.LogGw("Key file directory :", keyStorePath+"keystore/")
		recoverPrivkeyArmor(keyStorePath + keyStoreFilePath)

		algoBytes, err := ioutil.ReadFile(keyStorePath + keyStoreFilePath + algoFileName)
		if err != nil {
			util.LogErr(err)
		}
		cns.SelectKeyAlgorithm = string(algoBytes)

		util.LogGw("Key algorithm : ", cns.SelectKeyAlgorithm)
	}
}
