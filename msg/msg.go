package msg

import (
	"errors"

	"github.com/Moonyongjung/cns-gw/key"
	"github.com/Moonyongjung/cns-gw/msg/parse"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	wasm "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/mitchellh/mapstructure"
)

type ContractMsg struct {
	BankMsg               cns.BankMsgStruct
	InstantiateMsg        cns.InstantiateMsgStruct
	ExecuteMsg            cns.ExecuteMsgStruct
	QueryMsg              cns.QueryMsgStruct
	ListContractByCodeMsg cns.ListContractByCodeMsgStruct
	DownloadMsg           cns.DownloadMsgStruct
	CodeInfoMsg           cns.CodeInfoMsgStruct
	ContractInfoMsg       cns.ContractInfoMsgStruct
	ContractStateAllMsg   cns.ContractStateAllMsgStruct
	ContractHistoryMsg    cns.ContractHistoryMsgStruct
}

func (c ContractMsg) MsgConvert(body []byte, typeId string) ContractMsg {
	//-- No msg convert Get method
	if typeId == cns.BankSend {
		bankMsgData := util.JsonUnmarshalData(c.BankMsg, body)
		mapstructure.Decode(bankMsgData, &c.BankMsg)

		return c
	} else if typeId == cns.WasmInstantiate {
		instantiateMsgData := util.JsonUnmarshalData(c.InstantiateMsg, body)
		mapstructure.Decode(instantiateMsgData, &c.InstantiateMsg)

		return c
	} else if typeId == cns.WasmExecute {
		executeMsgData := util.JsonUnmarshalData(c.ExecuteMsg, body)
		mapstructure.Decode(executeMsgData, &c.ExecuteMsg)

		return c
	} else if typeId == cns.WasmQuery {
		queryMsgData := util.JsonUnmarshalData(c.QueryMsg, body)
		mapstructure.Decode(queryMsgData, &c.QueryMsg)

		return c
	} else if typeId == cns.WasmListContractByCode {
		listContractByCodeMsgData := util.JsonUnmarshalData(c.ListContractByCodeMsg, body)
		mapstructure.Decode(listContractByCodeMsgData, &c.ListContractByCodeMsg)

		return c
	} else if typeId == cns.WasmDownload {
		downloadMsgData := util.JsonUnmarshalData(c.DownloadMsg, body)
		mapstructure.Decode(downloadMsgData, &c.DownloadMsg)

		return c
	} else if typeId == cns.WasmCodeInfo {
		codeInfoMsgData := util.JsonUnmarshalData(c.CodeInfoMsg, body)
		mapstructure.Decode(codeInfoMsgData, &c.CodeInfoMsg)

		return c
	} else if typeId == cns.WasmContractInfo {
		contractInfoMsgData := util.JsonUnmarshalData(c.ContractInfoMsg, body)
		mapstructure.Decode(contractInfoMsgData, &c.ContractInfoMsg)

		return c
	} else if typeId == cns.WasmContractStateAll {
		contractStateAllMsgData := util.JsonUnmarshalData(c.ContractStateAllMsg, body)
		mapstructure.Decode(contractStateAllMsgData, &c.ContractStateAllMsg)

		return c
	} else if typeId == cns.WasmContractHistory {
		contractHistoryMsgData := util.JsonUnmarshalData(c.ContractHistoryMsg, body)
		mapstructure.Decode(contractHistoryMsgData, &c.ContractHistoryMsg)

		return c
	} else {
		return ContractMsg{}
	}
}

func (c ContractMsg) MakeBankMsg() (*bank.MsgSend, error) {

	bankMsgData := c.BankMsg
	denom := util.GetConfig().Get("denom")

	if (cns.BankMsgStruct{}) == bankMsgData {
		return nil, errors.New("Empty request or type of parameter is not correct")
	}

	fromAddress := bankMsgData.FromAddress
	if fromAddress == "" {
		return nil, errors.New("No fromAddress")
	}

	toAddress := bankMsgData.ToAddress
	if toAddress == "" {
		return nil, errors.New("No toAddress")
	}

	amount := bankMsgData.Amount
	if amount == "" {
		return nil, errors.New("No amount")
	}

	msg := &bank.MsgSend{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Amount:      sdk.NewCoins(sdk.NewInt64Coin(denom, util.FromStringToInt64(amount))),
	}

	return msg, nil
}

func (c ContractMsg) MakeStoreMsg() (wasm.MsgStoreCode, error) {
	contractPath := util.GetConfig().Get("contractPath")
	contractName := util.GetConfig().Get("contractName")
	gwAdd := util.GetAddrByPrivKeyArmor(key.GwKey().GetPriKey())

	msg, err := parse.ParseStoreCodeArgs(contractPath+contractName, gwAdd)
	if err != nil {
		util.LogErr(err)
		return wasm.MsgStoreCode{}, err
	}

	if err = msg.ValidateBasic(); err != nil {
		util.LogErr(err)
		return wasm.MsgStoreCode{}, err
	}

	return msg, nil
}

func (c ContractMsg) MakeInstantiateMsg() (wasm.MsgInstantiateContract, error) {
	gwAdd := util.GetAddrByPrivKeyArmor(key.GwKey().GetPriKey())

	if (cns.InstantiateMsgStruct{}) == c.InstantiateMsg {
		return wasm.MsgInstantiateContract{}, errors.New("Empty request or type of parameter is not correct")
	}

	msg, err := parse.ParseInstantiateArgs(c.InstantiateMsg, gwAdd)
	if err != nil {
		util.LogErr(err)
		return wasm.MsgInstantiateContract{}, err
	}

	if err = msg.ValidateBasic(); err != nil {
		util.LogErr(err)
		return wasm.MsgInstantiateContract{}, err
	}

	return msg, nil
}

func (c ContractMsg) MakeExecuteMsg() (wasm.MsgExecuteContract, error) {
	gwAdd := util.GetAddrByPrivKeyArmor(key.GwKey().GetPriKey())

	if (cns.ExecuteMsgStruct{}) == c.ExecuteMsg {
		return wasm.MsgExecuteContract{}, errors.New("Empty request or type of parameter is not correct")
	}

	msg, err := parse.ParseExecuteArgs(c.ExecuteMsg, gwAdd)
	if err != nil {
		util.LogErr(err)
		return wasm.MsgExecuteContract{}, err
	}

	if err = msg.ValidateBasic(); err != nil {
		util.LogErr(err)
		return wasm.MsgExecuteContract{}, err
	}

	return msg, nil
}

func (c ContractMsg) MakeQueryMsg() (wasm.QuerySmartContractStateRequest, error) {
	gwAdd := util.GetAddrByPrivKeyArmor(key.GwKey().GetPriKey())

	if (cns.QueryMsgStruct{}) == c.QueryMsg {
		return wasm.QuerySmartContractStateRequest{}, errors.New("Empty request or type of parameter is not correct")
	}

	msg, err := parse.ParseQueryArgs(c.QueryMsg, gwAdd)
	if err != nil {
		util.LogErr(err)
		return wasm.QuerySmartContractStateRequest{}, err
	}

	return msg, nil
}

func (c ContractMsg) MakeListcodeMsg() (wasm.QueryCodesRequest, error) {
	msg := parse.ParseListcodeArgs()
	return msg, nil
}

func (c ContractMsg) MakeListContractByCodeMsg() (wasm.QueryContractsByCodeRequest, error) {

	if (cns.ListContractByCodeMsgStruct{}) == c.ListContractByCodeMsg {
		return wasm.QueryContractsByCodeRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseListContractByCodeArgs(c.ListContractByCodeMsg)
	return msg, nil
}

func (c ContractMsg) MakeDownloadMsg() ([]interface{}, error) {
	var msgInterfaceSlice []interface{}
	if (cns.DownloadMsgStruct{}) == c.DownloadMsg {
		return nil, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseDownloadArgs(c.DownloadMsg)
	msgInterfaceSlice = append(msgInterfaceSlice, msg)
	msgInterfaceSlice = append(msgInterfaceSlice, c.DownloadMsg.DownloadFileName)
	return msgInterfaceSlice, nil
}

func (c ContractMsg) MakeCodeInfoMsg() (wasm.QueryCodeRequest, error) {

	if (cns.CodeInfoMsgStruct{}) == c.CodeInfoMsg {
		return wasm.QueryCodeRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseCodeInfoArgs(c.CodeInfoMsg)
	return msg, nil
}

func (c ContractMsg) MakeContractInfoMsg() (wasm.QueryContractInfoRequest, error) {

	if (cns.ContractInfoMsgStruct{}) == c.ContractInfoMsg {
		return wasm.QueryContractInfoRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseContractInfoArgs(c.ContractInfoMsg)
	return msg, nil
}

func (c ContractMsg) MakeContractStateAllMsg() (wasm.QueryAllContractStateRequest, error) {

	if (cns.ContractStateAllMsgStruct{}) == c.ContractStateAllMsg {
		return wasm.QueryAllContractStateRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseContractStateAllArgs(c.ContractStateAllMsg)
	return msg, nil
}

func (c ContractMsg) MakeContractHistoryMsg() (wasm.QueryContractHistoryRequest, error) {

	if (cns.ContractHistoryMsgStruct{}) == c.ContractHistoryMsg {
		return wasm.QueryContractHistoryRequest{}, errors.New("Empty request or type of parameter is not correct")
	}
	msg := parse.ParseContractHistoryArgs(c.ContractHistoryMsg)
	return msg, nil
}

func (c ContractMsg) MakePinnedMsg() (wasm.QueryPinnedCodesRequest, error) {
	msg := parse.ParsePinnedArgs()
	return msg, nil
}
