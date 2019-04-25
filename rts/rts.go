package rts

import (
	"git.parallelcoin.io/marcetin/explorer/exp"
	"github.com/gorilla/mux"
)

func RTS() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/block/{id}", viewBlock)
	r.HandleFunc("/blockhash/{blockhash}", viewBlockHash)
	r.HandleFunc("/tx/{id}", viewTx)
	r.HandleFunc("/addr/{id}", viewAddr)

	// crs := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	// 	AllowCredentials: true,
	// })

	r.HandleFunc("/addnode", EnableCors(NodeAdd)).Methods("POST")

	////////////////

	r.HandleFunc("/a/last", apiLast)
	r.HandleFunc("/a/info", apiInfo)
	r.HandleFunc("/a/mining", apiMiningInfo)
	r.HandleFunc("/a/rawpool", apiRawPool)
	r.HandleFunc("/search", doSearch)

	r.HandleFunc("/e/last", exp.LastBlock).Methods("GET")
	r.HandleFunc("/e/b", exp.Block).Methods("GET")
	r.HandleFunc("/e/block/{blockheight}", exp.BlockHeight).Methods("GET")
	r.HandleFunc("/e/b/{blockheight}", exp.BHeight).Methods("GET")
	r.HandleFunc("/e/hash/{blockhash}", exp.Hash).Methods("GET")

	r.HandleFunc("/e/tx/{txid}", exp.Tx).Methods("GET")
	r.HandleFunc("/e/rmp", exp.RawMemPool).Methods("GET")
	r.HandleFunc("/e/gmi", exp.MiningInfo).Methods("GET")
	r.HandleFunc("/e/info", exp.Info).Methods("GET")
	r.HandleFunc("/e/peer", exp.Peers).Methods("GET")
	r.HandleFunc("/e/addr/{addr}", exp.Address).Methods("GET")

	r.HandleFunc("/a/{type}/{id}", apiData)
	return r
}
