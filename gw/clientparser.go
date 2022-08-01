package gw

import (
	"io/ioutil"
	"net/http"
	"strings"

	cnsdb "github.com/Moonyongjung/cns-gw/db"
	"github.com/Moonyongjung/cns-gw/key"
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"
)

func doResponsebyClientRequest(
	w http.ResponseWriter,
	request *http.Request,
) (http.ResponseWriter, []byte) {

	var body []byte
	var returnWriter http.ResponseWriter
	var returnData []byte

	db := cns.Db

	checkRequest(request)
	reqType := strings.TrimLeft(request.URL.Path, cns.ClientApiVersion)
	util.LogGw(reqType)

	if request.Method == "POST" {
		bodyByte, err := ioutil.ReadAll(request.Body)
		if err != nil {
			util.LogErr(err)
		}
		body = bodyByte
	}

	if reqType == cns.WalletCreate {
		var createMnemonicResponse cns.CreateMnemonicResponse
		createMnemonicResponse.Mnemonic = key.UserNewMnemonic()
		responseData, err := util.JsonMarshalData(createMnemonicResponse)
		if err != nil {
			util.LogErr(err)
		}

		returnWriter = w
		returnData = responseData

	} else if reqType == cns.WalletAddress {

		util.LogGw("From client request body : ", string(body))

		var addressResponse cns.AddressResponse
		resMsg, redirectPage, pk := key.UserNewKey(string(body))
		addressResponse.ResMsg = resMsg
		addressResponse.DirectPage = redirectPage

		w = cnsdb.UserKeySessionDb(w, db, pk)

		responseData, err := util.JsonMarshalData(addressResponse)
		if err != nil {
			util.LogErr(err)
		}

		returnWriter = w
		returnData = responseData
	} else if reqType == cns.DomainMapping {

		if string(body) != "" {
			util.LogGw("From client request body: ", string(body))
			responseData := DomainMapping(string(body))

			returnWriter = w
			returnData = responseData
		} else {
			cookie, err := request.Cookie("session")
			if err != nil {
				util.LogErr("cookie read err : ", err)

				returnWriter = w
				returnData = util.HttpResponseByte(110, "", "")

			} else {
				pk := cnsdb.CheckSession(db, cookie.Value)

				if pk == "" {
					returnWriter = w
					returnData = util.HttpResponseByte(110, "", "")
				} else {
					returnWriter = w
					returnData = DomainMapping(util.GetAddrByPrivKeyArmor(pk).String())
				}
			}
		}

	} else if reqType == cns.DomainConfirm {
		util.LogGw("From client request body: ", string(body))
		responseData := DomainConfirm(string(body))

		returnWriter = w
		returnData = responseData
	} else if reqType == cns.SendIndex {
		cookie, err := request.Cookie("session")
		if err != nil {
			util.LogErr("cookie read err : ", err)

			returnWriter = w
			returnData = util.HttpResponseByte(110, "", "")

		} else {
			pk := cnsdb.CheckSession(db, cookie.Value)

			if pk == "" {
				returnWriter = w
				returnData = util.HttpResponseByte(110, "", "")
			} else {
				returnWriter = w
				returnData = util.HttpResponseByte(0, "", util.GetAddrByPrivKeyArmor(pk).String())
			}
		}
	} else if reqType == cns.SendInquiry {
		cookie, err := request.Cookie("session")
		if err != nil {
			util.LogErr("cookie read err : ", err)

			returnWriter = w
			returnData = util.HttpResponseByte(110, "", "")

		} else {
			pk := cnsdb.CheckSession(db, cookie.Value)

			if pk == "" {
				returnWriter = w
				returnData = util.HttpResponseByte(110, "", "")
			} else {
				responseByte := SendInquiry(string(body), pk)

				returnWriter = w
				returnData = responseByte
			}
		}
	}

	return returnWriter, returnData
}
