package gw

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"strings"
	
	"github.com/Moonyongjung/cns-gw/key"
	cnsdb "github.com/Moonyongjung/cns-gw/db"	
	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"	
)

func doResponsebyClientRequest(
	w http.ResponseWriter, 
	request *http.Request, 
	db *sql.DB,
	) (http.ResponseWriter, []byte) {

	var body []byte	
	var returnWriter http.ResponseWriter
	var returnData []byte

	checkRequest(request)
	requestType := strings.Split(request.URL.Path, "/")[2:]
	
	if request.Method == "POST" {		
		bodyByte, err := ioutil.ReadAll(request.Body)
		if err != nil {
			util.LogGw(err)			
		}
		body = bodyByte
	}

	if requestType[0] == "wallet" {
		if requestType[1] == "create" {
			var createMnemonicResponse cns.CreateMnemonicResponse
			createMnemonicResponse.Mnemonic = key.UserNewMnemonic()
			responseData, err := util.JsonMarshalData(createMnemonicResponse)
			if err != nil {
				util.LogGw(err)
			}		

			returnWriter = w
			returnData = responseData		

		} else if requestType[1] == "address" {
			
			util.LogGw("From client : ", string(body))
	
			var addressResponse cns.AddressResponse
			resMsg, redirectPage, pk := key.UserNewKey(string(body))
			addressResponse.ResMsg = resMsg
			addressResponse.DirectPage = redirectPage
	
			util.LogGw(pk)
			w = cnsdb.UserKeySessionDb(w, db, pk)
	
			responseData, err := util.JsonMarshalData(addressResponse)
			if err != nil {
				util.LogGw(err)
			}		

			returnWriter = w
			returnData = responseData		
		}
	} else if requestType[0] == "domain" {
		if requestType[1] == "mapping" {

			if string(body) != "" {
				util.LogGw("From client : ", string(body))			
				responseData := DomainMapping(string(body))
	
				returnWriter = w
				returnData = responseData
			} else {
				cookie, err := request.Cookie("session")
				if err != nil {
					util.LogGw("cookie read err : ", err)
	
					returnWriter = w
					returnData = httpResponseByte(110, "")				
	
				} else {
					pk := cnsdb.CheckSession(db, cookie.Value)
	
					if pk == "" {				
						returnWriter = w
						returnData = httpResponseByte(110, "")									
					} else {						
						returnWriter = w
						returnData = DomainMapping(util.GetAddrByPrivKeyArmor(pk).String())						
					}				
				}
			}

		} else if requestType[1] == "confirm" {
			util.LogGw("From client : ", string(body))	
			responseData := DomainConfirm(string(body))		

			returnWriter = w
			returnData = responseData
		}
	} else if requestType[0] == "send" {
		if requestType[1] == "index" {
			cookie, err := request.Cookie("session")
			if err != nil {
				util.LogGw("cookie read err : ", err)

				returnWriter = w
				returnData = httpResponseByte(110, "")				

			} else {
				pk := cnsdb.CheckSession(db, cookie.Value)

				if pk == "" {				
					returnWriter = w
					returnData = httpResponseByte(110, "")									
				} else {
					returnWriter = w
					returnData = httpResponseByte(0, util.GetAddrByPrivKeyArmor(pk).String())
				}				
			}
		} else if requestType[1] == "inquiry" {
			cookie, err := request.Cookie("session")
			if err != nil {
				util.LogGw("cookie read err : ", err)

				returnWriter = w
				returnData = httpResponseByte(110, "")				

			} else {
				pk := cnsdb.CheckSession(db, cookie.Value)

				if pk == "" {				
					returnWriter = w
					returnData = httpResponseByte(110, "")									
				} else {
					responseByte := SendInquiry(string(body), pk)
					
					returnWriter = w
					returnData = responseByte
				}				
			}
		}
	}	

	return returnWriter, returnData
}