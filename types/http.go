package types

var ResponseType map[int]string

type CreateMnemonicResponse struct {
	Mnemonic string
}

type AddressResponse struct {
	ResMsg string
	DirectPage string
}

type DomainMappingResponse struct {
	ResMsg string
	DirectPage string
}

type DomainMappingRequest struct {
	DomainName string `json:"domainName`
	AccountAddress string `json:"accountAddress"`
}

