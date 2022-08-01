package key

import (
	"io/ioutil"
	"os"

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	cmclient "github.com/cosmos/cosmos-sdk/client"
	cosmoshd "github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/evmos/ethermint/crypto/hd"
)

type EncryptedJson struct {
	PriKey       string
	PubKey       string
	GwKeyAddress sdk.AccAddress
}

func recoverPrivkeyArmor(keyStorePath string) {
	priKeyBytes, err := ioutil.ReadFile(keyStorePath + priKeyFileName)
	if err != nil {
		util.LogErr(err)
	}

	pubKeyBytes, err := ioutil.ReadFile(keyStorePath + pubKeyFileName)
	if err != nil {
		util.LogErr(err)
	}

	GwKey().Set(string(priKeyBytes), string(pubKeyBytes))
}

func exportPrivKeyArmor(
	keyOwnerName string,
	keyStorePath string,
	keyOwnerPw string,
	algo string,
	clientCtx cmclient.Context) {

	input(keyOwnerPw)

	pri, err := clientCtx.Keyring.ExportPrivKeyArmor(keyOwnerName, keyOwnerPw)
	if err != nil {
		util.LogErr(err)
	}

	pub, err := clientCtx.Keyring.ExportPubKeyArmor(keyOwnerName)
	if err != nil {
		util.LogErr(err)
	}

	err = ioutil.WriteFile(keyStorePath+priKeyFileName, []byte(pri), 0660)
	if err != nil {
		util.LogErr(err)
	}

	err = ioutil.WriteFile(keyStorePath+pubKeyFileName, []byte(pub), 0660)
	if err != nil {
		util.LogErr(err)
	}

	GwKey().Set(pri, pub)
}

//-- Use keyring
func storeKeyringBackend(
	keyOwnerName string,
	keyStorePath string,
	keyOwnerPw string,
	algo string,
	clientCtx cmclient.Context) cmclient.Context {

	kr, algorithm := keyringSetup(keyStorePath, algo)

	_, mnemonic, err := kr.NewMnemonic(
		keyOwnerName,
		keyring.English,
		cns.FullFundraiserPath,
		keyring.DefaultBIP39Passphrase,
		algorithm,
	)
	if err != nil {
		util.LogErr(err)
	}

	util.LogGw("Gateway mnemonic : ", mnemonic)

	err = ioutil.WriteFile(keyStorePath+mnemonicFileName, []byte(mnemonic), 0660)
	if err != nil {
		util.LogErr(err)
	}

	clientCtx = clientCtx.WithKeyring(kr)
	return clientCtx
}

func keyringSetup(keyStorePath string, algo string) (keyring.Keyring, keyring.SignatureAlgo) {
	var kr keyring.Keyring
	var err error
	var algorithm keyring.SignatureAlgo

	if algo == "eth_secp256k1" {
		kr, err = keyring.New(
			"gw",
			keyring.BackendMemory,
			keyStorePath,
			os.Stdin,
			hd.EthSecp256k1Option(),
		)
		if err != nil {
			util.LogErr(err)
		}

		algorithm = hd.EthSecp256k1
	} else {
		kr, err = keyring.New(
			"gw",
			keyring.BackendMemory,
			keyStorePath,
			os.Stdin,
		)
		if err != nil {
			util.LogErr(err)
		}

		algorithm = cosmoshd.Secp256k1
	}

	return kr, algorithm
}
