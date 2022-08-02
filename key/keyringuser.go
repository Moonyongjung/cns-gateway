package key

import (
	"os"

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

func UserNewMnemonic() string {
	keyOwnerName := util.GetConfig().Get("keyOwnerName")
	keyOwnerPw := util.GetConfig().Get("keyOwnerPw")
	keyStorePath := util.GetConfig().Get("keyStorePath")
	keyStoreFilePath := util.GetConfig().Get("keyStoreFilePath")

	input(keyOwnerPw)

	kr, algorithm := keyringSetup(keyStorePath+keyStoreFilePath, cns.SelectKeyAlgorithm)

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

	return mnemonic
}

func UserNewKey(mnemonic string) (string, string, string) {
	if mnemonic == "" {
		return "No mnemonic words. Go to create page", "redirect", ""
	}

	keyOwnerName := util.GetConfig().Get("keyOwnerName")
	keyOwnerPw := util.GetConfig().Get("keyOwnerPw")
	keyStorePath := util.GetConfig().Get("keyStorePath")
	keyStoreFilePath := util.GetConfig().Get("keyStoreFilePath")

	input(keyOwnerPw)

	kr, algorithm := keyringSetup(keyStorePath+keyStoreFilePath, cns.SelectKeyAlgorithm)

	info, err := kr.NewAccount(
		keyOwnerName,
		mnemonic,
		keyring.DefaultBIP39Passphrase,
		cns.FullFundraiserPath,
		algorithm,
	)
	if err != nil {
		util.LogErr("New Account err : ", err)
		return err.Error(), "redirect", ""
	}

	priv, err := kr.ExportPrivKeyArmor(keyOwnerName, keyOwnerPw)
	if err != nil {
		util.LogErr(err)
		return err.Error(), "redirect", ""
	}

	util.LogGw(info.GetAddress().String())
	util.LogGw(util.GetAddrByPrivKeyArmor(priv).String())
	return info.GetAddress().String(), "normal", priv
}

func input(keyOwnerPw string) {
	input := []byte(keyOwnerPw + "\n" + keyOwnerPw + "\n")
	r, w, err := os.Pipe()
	if err != nil {
		util.LogErr(err)
	}

	_, err = w.Write(input)
	if err != nil {
		util.LogErr(err)
	}
	w.Close()

	stdin := os.Stdin

	defer func() { os.Stdin = stdin }()

	os.Stdin = r
}
