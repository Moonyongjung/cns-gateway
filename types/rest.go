//-- Get account info parameters
package types

type ErrStruct struct {
	Code    float64
	Message string
	Details string
}

type ResponseStruct struct {
	Account AccountStruct
}

type AccountStruct struct {
	Type          string
	Address       string
	Pubkey        PubKeyStruct
	AccountNumber string
	Sequence      string
}

type PubKeyStruct struct {
	Type string
	Key  string
}

//-- Ethermint module chain response
type EthResponseStruct struct {
	Account EthAccountStruct
}

type EthAccountStruct struct {
	Type        string               `mapstructure:"@type"`
	BaseAccount EthBaseAccountStruct `mapstructure:"base_account"`
	CodeHash    string               `mapstructure:"code_hash"`
}

type EthBaseAccountStruct struct {
	Address       string `mapstructure:"address"`
	PubKey        string `mapstructure:"pub_key"`
	AccountNumber string `mapstructure:"account_number"`
	Sequence      string `mapstructure:"sequence"`
}
