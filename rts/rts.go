package rts

import (
	"net/http"

	"git.parallelcoin.io/marcetin/explorer/exp"
	"git.parallelcoin.io/marcetin/explorer/utl"
	"github.com/gorilla/mux"
)

func RTS() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", HomeHandler)
	a := r.PathPrefix("/a").Subrouter()

	a.HandleFunc("/coins", exp.ViewCoins).Methods("GET")
	a.HandleFunc("/addcoin", utl.EnableCors(exp.AddCoin)).Methods("POST")

	a.HandleFunc("/{coin}/addrpcsrc", utl.EnableCors(exp.AddRPCSource)).Methods("POST")

	a.HandleFunc("/{coin}/blocks/{per}/{page}", exp.ViewBlocks).Methods("GET")
	a.HandleFunc("/{coin}/b", exp.ViewBlock).Methods("GET")
	a.HandleFunc("/{coin}/block/{blockheight}", exp.ViewBlockHeight).Methods("GET")
	a.HandleFunc("/{coin}/b/{blockheight}", exp.ViewBlockHeight).Methods("GET")
	a.HandleFunc("/{coin}/hash/{blockhash}", exp.ViewHash).Methods("GET")

	a.HandleFunc("/{coin}/tx/{txid}", exp.ViewTx).Methods("GET")
	a.HandleFunc("/{coin}/rawpool", exp.ViewRawMemPool).Methods("GET")
	a.HandleFunc("/{coin}/mining", exp.ViewMiningInfo).Methods("GET")
	a.HandleFunc("/{coin}/info", exp.ViewInfo).Methods("GET")
	a.HandleFunc("/{coin}/peer", exp.ViewPeers).Methods("GET")
	a.HandleFunc("/{coin}/addr/{addr}", exp.ViewAddress).Methods("GET")

	a.HandleFunc("/{coin}/getrichlist", exp.GetAddrsList).Methods("GET")
	a.HandleFunc("/{coin}/getnodes", exp.GetNodes).Methods("GET")

	a.HandleFunc("/{coin}/nodes", exp.NodesMap).Methods("GET")

	c := r.PathPrefix("/c").Subrouter()
	c.HandleFunc("/{coin}/{chart}/{cat}/{data}", exp.GetChart).Methods("GET")

	// c.HandleFunc("/{coin}/{chart}/{cat}/{data}", exp.GetChart).Methods("GET")
	j := r.PathPrefix("/j").Subrouter()
	j.HandleFunc("/{coin}/{cat}/{data}", exp.GetChartData).Methods("GET")

	d := r.PathPrefix("/d").Subrouter()
	d.HandleFunc("/{coin}/{cat}/{data}", exp.GetData).Methods("GET")

	s := r.PathPrefix("/s").Subrouter()
	s.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./tpl/static/"))))

	// r.HandleFunc("/{coin}/{type}/{id}", apiData)
	return r
}
