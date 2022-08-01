package gw

import (
	"net/http"

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	"github.com/rs/cors"
)

//-- HTTPServer operates for sending or invoking transaction when user call
func RunHttpServer(channel cns.ChannelStruct) {
	gatewayServerPort := util.GetConfig().Get("gatewayServerPort")
	mux := http.NewServeMux()

	mux = ClientResponseMux(mux)

	mux.HandleFunc(cns.ApiVersion+cns.BankSend, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmStore, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmInstantiate, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.CnsWasmInstantiate, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmExecute, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.CnsWasmExecute, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmQuery, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.CnsWasmQueryByDomain, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.CnsWasmQueryByAccount, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmListCode, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmListContractByCode, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmDownload, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmCodeInfo, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmContractInfo, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmContractStateAll, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmContractHistory, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})
	mux.HandleFunc(cns.ApiVersion+cns.WasmPinned, func(w http.ResponseWriter, r *http.Request) {
		httpCons(w, r, channel)
	})

	handler := cors.Default().Handler(mux)
	util.LogHttpServer("Server Listen...")

	err := http.ListenAndServe(":"+gatewayServerPort, handler)
	if err != nil {
		util.LogHttpServer(err)
	}
}

func httpCons(w http.ResponseWriter, r *http.Request, channel cns.ChannelStruct) {
	doTransactionbyType(r, channel)
	select {
	case result := <-channel.HttpServerChan:
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}
