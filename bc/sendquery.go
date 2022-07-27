package bc

import (
	"context"
	"errors"
	"io/ioutil"
	"strings"

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cmclient "github.com/cosmos/cosmos-sdk/client"
	
	"github.com/gogo/protobuf/proto"
)

//-- Query contract state using by NewQueryClient of wasm
func querySend(clientCtx cmclient.Context, 
	msg interface{},
	option string) (string, cns.HttpResponseStruct){
	
	var response cns.HttpResponseStruct
	var out []byte

	queryClient := wasmtypes.NewQueryClient(clientCtx)

	if option == "query" {
		convertMsg, _ := msg.(wasmtypes.QuerySmartContractStateRequest)		
		res, err := queryClient.SmartContractState(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}

		out, err = PrintProto(clientCtx, res)
		if err != nil {
			return queryErrReturn(err, response)
		}
	}

	if option == "listcode" {
		convertMsg, _ := msg.(wasmtypes.QueryCodesRequest)		
		res, err := queryClient.Codes(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}	

		out, err = PrintProto(clientCtx, res)
		if err != nil {
			return queryErrReturn(err, response)
		}
	}

	if option == "listcontractbycode" {
		convertMsg, _ := msg.(wasmtypes.QueryContractsByCodeRequest)		
		res, err := queryClient.ContractsByCode(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}	

		out, err = PrintProto(clientCtx, res)
		if err != nil {
			return queryErrReturn(err, response)
		}
	}

	if option == "download" {
		convertMsg, _ := msg.([]interface{})[0].(wasmtypes.QueryCodeRequest)
		downloadFileName, _ := msg.([]interface{})[1].(string)
		
		if !strings.Contains(downloadFileName, ".wasm") {
			downloadFileName = downloadFileName + ".wasm"
		}
		res, err := queryClient.Code(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}	
		
		if len(res.Data) == 0 {
			return queryErrReturn(errors.New("contract not found"), response)
		}
		out = []byte("Downloading wasm code to "+downloadFileName)
		

		ioutil.WriteFile(downloadFileName, res.Data, 0o600)
	}

	if option == "codeinfo" {
		convertMsg, _ := msg.(wasmtypes.QueryCodeRequest)		
		res, err := queryClient.Code(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}	

		out, err = PrintProto(clientCtx, res.CodeInfoResponse)
		if err != nil {
			return queryErrReturn(err, response)
		}
	}

	if option == "contractinfo" {
		convertMsg, _ := msg.(wasmtypes.QueryContractInfoRequest)		
		res, err := queryClient.ContractInfo(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}	

		out, err = PrintProto(clientCtx, res)
		if err != nil {
			return queryErrReturn(err, response)
		}
	}

	if option == "contractstateall" {
		convertMsg, _ := msg.(wasmtypes.QueryAllContractStateRequest)		
		res, err := queryClient.AllContractState(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}	

		out, err = PrintProto(clientCtx, res)
		if err != nil {
			return queryErrReturn(err, response)
		}
	}

	if option == "contracthistory" {
		convertMsg, _ := msg.(wasmtypes.QueryContractHistoryRequest)		
		res, err := queryClient.ContractHistory(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}	

		out, err = PrintProto(clientCtx, res)
		if err != nil {
			return queryErrReturn(err, response)
		}
	}

	if option == "pinned" {
		convertMsg, _ := msg.(wasmtypes.QueryPinnedCodesRequest)		
		res, err := queryClient.PinnedCodes(
			context.Background(),
			&convertMsg,
		)
		if err != nil {
			return queryErrReturn(err, response)
		}	

		out, err = PrintProto(clientCtx, res)
		if err != nil {
			return queryErrReturn(err, response)
		}
	}

	util.LogGw("Transaction response", string(out))	
	response.ResCode = 0
	return string(out), response
}

func PrintProto(clientCtx cmclient.Context, toPrint proto.Message) ([]byte, error) {
	out, err := clientCtx.Codec.MarshalJSON(toPrint)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func queryErrReturn(err error, 
	response cns.HttpResponseStruct) (string, cns.HttpResponseStruct){

	response.ResCode = 105
	response.ResData = err.Error()

	if strings.Contains(err.Error(), "InvalidArgument") {
		response.ResData = "Requested domain name or account address does not exist"
	}
	return "", response
}
