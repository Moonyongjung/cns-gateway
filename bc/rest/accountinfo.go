package rest

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	"github.com/mitchellh/mapstructure"
)

const userInfoUrl = "/cosmos/auth/v1beta1/accounts/"

//-- Get account number and sequence
func HttpClient(method string, url string, body []byte) []byte {
	var request *http.Request
	var err error

	if method == "GET" {
		request, err = http.NewRequest("GET", url, nil)
		if err != nil {
			util.LogHttpClient(err)
		}
	} else {
		buf := bytes.NewBuffer(body)
		request, err = http.NewRequest("POST", url, buf)
		if err != nil {
			util.LogHttpClient(err)
		}
	}

	hClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 60 * time.Second,
	}
	hClient.Timeout = time.Second * 30
	defer func() {
		if err := recover(); err != nil {
			util.LogHttpClient(err)
		}
	}()
	response, err := hClient.Do(request)
	if err != nil {
		util.LogHttpClient(err)
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		util.LogHttpClient(err)
	}

	response.Body.Close()

	return responseBody
}

func GetAccountInfoHttpClient(gwKeyAddress string) (string, string, error) {
	restEndpoint := util.GetConfig().Get("restEndpoint")

	url := restEndpoint + userInfoUrl + gwKeyAddress

	responseBody := HttpClient("GET", url, nil)

	util.LogHttpClient("Account Response : " + "\n" + string(responseBody))

	//-- The account does not have any coins or tokens, not included the chain
	//-- EthAccount -> eth_secp256k1
	if strings.Contains(string(responseBody), "EthAccount") {
		var ethResponseStruct cns.EthResponseStruct
		ethResponseData := util.JsonUnmarshalData(ethResponseStruct, responseBody)
		err := mapstructure.Decode(ethResponseData, &ethResponseStruct)
		if err != nil {
			util.LogErr(err)
		}

		accountNumber := ethResponseStruct.Account.BaseAccount.AccountNumber
		accountSequence := ethResponseStruct.Account.BaseAccount.Sequence

		cns.GatewayAccount = ethResponseStruct.Account.BaseAccount.Address
		util.LogHttpClient("Gateway account address : ", cns.GatewayAccount)

		return accountNumber, accountSequence, nil

	} else {
		//-- secp256k1
		if strings.Contains(string(responseBody), "code") {
			var errStruct cns.ErrStruct
			errCheckData := util.JsonUnmarshalData(errStruct, responseBody)

			code := errCheckData.(map[string]interface{})["code"].(float64)

			return "", "", errors.New("Code : " + util.ToString(code, ""))

		} else {
			var responseStruct cns.ResponseStruct
			responseData := util.JsonUnmarshalData(responseStruct, responseBody)

			accountNumber := responseData.(map[string]interface{})["account"].(map[string]interface{})["account_number"].(string)

			sequence := responseData.(map[string]interface{})["account"].(map[string]interface{})["sequence"].(string)

			cns.GatewayAccount = responseData.(map[string]interface{})["account"].(map[string]interface{})["address"].(string)

			return accountNumber, sequence, nil
		}
	}
}
