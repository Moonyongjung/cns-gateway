package types

//-- Saved contract code ID format
type ContractCodeJsonStruct struct {
	TypeName string
	CodeId   string
}

//-- Save contract address format
type ContractAddressJsonStruct struct {
	TypeName        string
	ContractAddress string
	CodeType        string
	CodeId          string
}
