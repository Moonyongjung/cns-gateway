package gw

import (
	"database/sql"	
	"net/http"
)

//-- CNS client response
func ClientResponseMux(mux *http.ServeMux, db *sql.DB) *http.ServeMux{	

	mux.HandleFunc("/client/wallet/create/", func(w http.ResponseWriter, r *http.Request) {			
		clientHttpCons(w , r , db)		
	})
	mux.HandleFunc("/client/wallet/address/", func(w http.ResponseWriter, r *http.Request) {			
		clientHttpCons(w , r , db)
	})
	mux.HandleFunc("/client/domain/mapping/", func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w , r , db)		
	})
	mux.HandleFunc("/client/domain/confirm/", func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w , r , db)		
	})
	mux.HandleFunc("/client/send/index/", func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w , r , db)		
	})
	mux.HandleFunc("/client/send/inquiry/", func(w http.ResponseWriter, r *http.Request) {
		clientHttpCons(w , r , db)		
	})

	return mux
}

func clientHttpCons(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	res, responsedata := doResponsebyClientRequest(w, r, db)
	res.WriteHeader(http.StatusOK)
	res.Write(responsedata)		
}