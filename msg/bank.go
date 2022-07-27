package msg

import (
	"errors"

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func MakeBankMsg(bankMsgData cns.BankMsgStruct) (*bank.MsgSend, error) {
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
		ToAddress: toAddress,
		Amount: sdk.NewCoins(sdk.NewInt64Coin(denom, util.FromStringToInt64(amount))),
	}

	return msg, nil
}
