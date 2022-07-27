package chain

import (
	"fmt"
	"io/ioutil"

	"github.com/Moonyongjung/cns-gw/util"
)

//-- Select chain's key algorithm
func SelectChainType() string {
	for {
		var s string
		util.LogGw("=======================================================================")
		util.LogGw("                    Initialize Gateway Key Type                        ")
		util.LogGw("                                                                       ")
		util.LogGw("SELECT the chain type                                                  ")
		util.LogGw("                                                                       ")
		util.LogGw("A chain includes ethermint module uses eth_secp256k1 key algorithm     ")
		util.LogGw("but general chains based on cosmos sdk uses secp256k1 algo.            ")
		util.LogGw("If you want to use the chain includes ethermint module,                ")
		util.LogGw("input [y/N]                                                            ")
		util.LogGw("                                                                       ")
		util.LogGw("y. Use includes ethermint module, N. general Cosmos SDK based chain    ")
		util.LogGw("=======================================================================")
	
		fmt.Scan(&s)
	
		if s == "y" {
			return "eth_secp256k1"
		} else if s == "N" {
			return "secp256k1"
		} else {
			util.LogGw("Input correct string [y/N]")
		}
	}
}

//-- Save chain's key algorithm
func SaveChainType(keyStorePath string, algo string, algoFileName string) {	
	
	err := ioutil.WriteFile(keyStorePath+algoFileName, []byte(algo), 0660)
	if err != nil {
		util.LogGw(err)			
	}
}