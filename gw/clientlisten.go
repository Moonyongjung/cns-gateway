package gw

import (
	"net/http"

	cns "github.com/Moonyongjung/cns-gw/types"
)

//-- CNS client response
func ClientResponseMux(mux *http.ServeMux) *http.ServeMux {

	mux.HandleFunc(cns.ClientApiVersion+cns.WalletCreate, func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w, r)
	})
	mux.HandleFunc(cns.ClientApiVersion+cns.WalletAddress, func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w, r)
	})
	mux.HandleFunc(cns.ClientApiVersion+cns.DomainMapping, func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w, r)
	})
	mux.HandleFunc(cns.ClientApiVersion+cns.DomainConfirm, func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w, r)
	})
	mux.HandleFunc(cns.ClientApiVersion+cns.SendIndex, func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w, r)
	})
	mux.HandleFunc(cns.ClientApiVersion+cns.SendInquiry, func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w, r)
	})

	return mux
}

func clientHttpCons(w http.ResponseWriter, r *http.Request) {
	res, responsedata := doResponsebyClientRequest(w, r)
	res.WriteHeader(http.StatusOK)
	res.Write(responsedata)
}
