package gw

import (
	"io/ioutil"
	"net/http"
	"strings"

	cnsdb "github.com/Moonyongjung/cns-gw/db"
	"github.com/Moonyongjung/cns-gw/msg"
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
	var clientRes msg.CnsClientParse

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
		returnWriter, returnData = clientRes.ResList(reqType, w, nil)

	} else if reqType == cns.WalletAddress {
		returnWriter, returnData = clientRes.ResList(reqType, w, body)

	} else if reqType == cns.DomainMapping {

		if string(body) != "" {
			util.LogGw("From client request body: ", string(body))
			returnWriter, returnData = w, DomainMapping(string(body))

		} else {
			returnWriter, returnData = returnValuewithCookie(reqType, request, w, body)
		}

	} else if reqType == cns.DomainConfirm {
		util.LogGw("From client request body: ", string(body))
		returnWriter, returnData = w, DomainConfirm(string(body))

	} else if reqType == cns.SendIndex {
		returnWriter, returnData = returnValuewithCookie(reqType, request, w, body)

	} else if reqType == cns.SendInquiry {
		returnWriter, returnData = returnValuewithCookie(reqType, request, w, body)
	}

	return returnWriter, returnData
}

func clientParseErrReturn(
	err error,
	w http.ResponseWriter) (http.ResponseWriter, []byte) {

	util.LogErr("ERROR, ", err)
	response := util.HttpResponseByte(110, "", err.Error())

	return w, response
}

func returnValuewithCookie(
	reqType string,
	request *http.Request,
	w http.ResponseWriter,
	body []byte) (http.ResponseWriter, []byte) {

	var returnNormalByte []byte
	db := cns.Db
	cookie, err := request.Cookie("session")
	if err != nil {
		util.LogErr("cookie read err : ", err)
		return clientParseErrReturn(err, w)
	} else {
		pk := cnsdb.CheckSession(db, cookie.Value)

		if pk == "" {
			return clientParseErrReturn(err, w)
		} else {
			if reqType == cns.DomainMapping {
				returnNormalByte = DomainMapping(util.GetAddrByPrivKeyArmor(pk).String())
			} else if reqType == cns.SendIndex {
				returnNormalByte = util.HttpResponseByte(0, "", util.GetAddrByPrivKeyArmor(pk).String())
			} else if reqType == cns.SendInquiry {
				returnNormalByte = SendInquiry(string(body), pk)
			}

			return w, returnNormalByte
		}
	}
}
