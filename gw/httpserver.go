package gw

import (
	"database/sql"
	"net/http"

	cns "github.com/Moonyongjung/cns-gw/types"
	"github.com/Moonyongjung/cns-gw/util"

	"github.com/rs/cors"
)

//-- HTTPServer operates for sending or invoking transaction when user call
func RunHttpServer(channel cns.ChannelStruct, db *sql.DB) {
	gatewayServerPort := util.GetConfig().Get("gatewayServerPort")
	mux := http.NewServeMux()	

	mux = ClientResponseMux(mux, db)	
	
	//-- APIs for interacting CosmWasm contract
	mux.HandleFunc("/api/bank/send", func(w http.ResponseWriter, r *http.Request) {			
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/store", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/instantiate", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/default-cns-instantiate", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/execute", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)	
	})
	mux.HandleFunc("/api/wasm/cns-execute", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)	
	})
	mux.HandleFunc("/api/wasm/query", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/cns-query-by-domain", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/cns-query-by-account", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/list-code", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/list-contract-by-code", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/download", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/code-info", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/contract-info", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/contract-state-all", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/contract-history", func(w http.ResponseWriter, r *http.Request) {		
		httpCons(w, r, channel)
	})
	mux.HandleFunc("/api/wasm/pinned", func(w http.ResponseWriter, r *http.Request) {		
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
	case result := <- channel.HttpServerChan:
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	}
}