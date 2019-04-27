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

	r.HandleFunc("/addnode", EnableCors(NodeAdd)).Methods("POST")

	////////////////

	// r.HandleFunc("/a/last", apiLast)
	// r.HandleFunc("/a/info", apiInfo)
	// r.HandleFunc("/a/mining", apiMiningInfo)
	// r.HandleFunc("/a/rawpool", apiRawPool)
	// r.HandleFunc("/search", doSearch)

	//////////////

	r.HandleFunc("/a/last", exp.LastBlock).Methods("GET")
	r.HandleFunc("/a/b", exp.Block).Methods("GET")
	r.HandleFunc("/a/block/{blockheight}", exp.BlockHeight).Methods("GET")
	r.HandleFunc("/a/b/{blockheight}", exp.BHeight).Methods("GET")
	r.HandleFunc("/a/hash/{blockhash}", exp.Hash).Methods("GET")

	r.HandleFunc("/a/tx/{txid}", exp.Tx).Methods("GET")
	r.HandleFunc("/a/rawpool", exp.RawMemPool).Methods("GET")
	r.HandleFunc("/a/mining", exp.MiningInfo).Methods("GET")
	r.HandleFunc("/a/info", exp.Info).Methods("GET")
	r.HandleFunc("/a/peer", exp.Peers).Methods("GET")
	r.HandleFunc("/a/addr/{addr}", exp.Address).Methods("GET")

	// r.HandleFunc("/a/{type}/{id}", apiData)
	return r
}
