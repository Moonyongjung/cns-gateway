package types

type CreateMnemonicResponse struct {
	Mnemonic string
}

type AddressResponse struct {
	ResMsg     string
	DirectPage string
}

type DomainMappingRequest struct {
	DomainName     string `json:"domainName`
	AccountAddress string `json:"accountAddress"`
}
