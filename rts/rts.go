package rts

import (
	"git.parallelcoin.io/marcetin/explorer/exp"
	"github.com/gorilla/mux"
)

func RTS() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/block/{id}", viewBlock)
	// r.HandleFunc("/blockhash/{blockhash}", viewBlockHash)
	// r.HandleFunc("/tx/{id}", viewTx)
	// r.HandleFunc("/addr/{id}", viewAddr)

	// crs := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	// 	AllowCredentials: true,
	// })

	////////////////

	// r.HandleFunc("/{coin}/last", apiLast)
	// r.HandleFunc("/{coin}/info", apiInfo)
	// r.HandleFunc("/{coin}/mining", apiMiningInfo)
	// r.HandleFunc("/{coin}/rawpool", apiRawPool)
	// r.HandleFunc("/search", doSearch)

	//////////////
	a := r.PathPrefix("/a").Subrouter()
	r.HandleFunc("/addcoin", EnableCors(AddCoin)).Methods("POST")
	a.HandleFunc("/coins", exp.ViewCoins).Methods("GET")

	r.HandleFunc("/{coin}/addnode", EnableCors(AddNode)).Methods("POST")
	a.HandleFunc("/{coin}/blocks/{per}/{page}", exp.ViewBlocks).Methods("GET")
	a.HandleFunc("/{coin}/b", exp.ViewBlock).Methods("GET")
	a.HandleFunc("/{coin}/block/{blockheight}", exp.ViewBlockHeight).Methods("GET")
	a.HandleFunc("/{coin}/b/{blockheight}", exp.ViewHeight).Methods("GET")
	a.HandleFunc("/{coin}/hash/{blockhash}", exp.ViewHash).Methods("GET")

	a.HandleFunc("/{coin}/tx/{txid}", exp.ViewTx).Methods("GET")
	a.HandleFunc("/{coin}/rawpool", exp.ViewRawMemPool).Methods("GET")
	a.HandleFunc("/{coin}/mining", exp.ViewMiningInfo).Methods("GET")
	a.HandleFunc("/{coin}/info", exp.ViewInfo).Methods("GET")
	a.HandleFunc("/{coin}/peer", exp.ViewPeers).Methods("GET")
	a.HandleFunc("/{coin}/addr/{addr}", exp.ViewAddress).Methods("GET")

	// r.HandleFunc("/{coin}/{type}/{id}", apiData)
	return r
}
