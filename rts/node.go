package rts

import (
	"fmt"
	"net/http"
	"strconv"

	"git.parallelcoin.io/marcetin/explorer/exp"
	"git.parallelcoin.io/marcetin/explorer/jdb"
	"git.parallelcoin.io/marcetin/explorer/utl"
)

// func Start(w http.ResponseWriter, r *http.Request) {
// 	cf := coinAdd(r)
// 	c := coinAdd(r)
// 	n := nodeAdd(r)
// 	jr := map[string]interface{}{
// 		"conf": cf,
// 		"coin": c,
// 		"node": n,
// 	}
// 	w.Header().Set("AMP-Redirect-To", "http://localhost:4000/")
// 	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin, AMP-Redirect-To")
// 	SendJsonResponse(w, jr)
// }

func AddCoin(w http.ResponseWriter, r *http.Request) {
	c := coinAdd(r)
	w.Header().Set("AMP-Redirect-To", "http://localhost:4000/")
	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin, AMP-Redirect-To")
	SendJsonResponse(w, c)
}

func coinAdd(r *http.Request) (c exp.Coin) {
	c.Name = r.FormValue("name")
	c.Slug = utl.MakeSlug(c.Name)
	c.Abbr = r.FormValue("abbr")
	c.X, _ = strconv.ParseInt(r.FormValue("x"), 10, 64)
	jdb.JDB.Write("coins/", c.Slug, c)
	return c
}
func nodeAdd(r *http.Request) (n exp.Node) {
	n.Coin = r.FormValue("coin")
	fmt.Println("asassaaaablockCountblockCountaaaa", n.Coin)

	n.NodeID = r.FormValue("nodeid")
	n.Slug = utl.MakeSlug(n.NodeID)
	n.RPCUser = r.FormValue("rpcuser")
	n.RPCPassword = r.FormValue("rpcpassword")
	n.IP = r.FormValue("ip")
	n.Port, _ = strconv.ParseInt(r.FormValue("port"), 10, 64)
	jdb.JDB.Write("data/"+n.Coin+"/nodes", n.Slug, n)
	return n
}

func AddNode(w http.ResponseWriter, r *http.Request) {
	n := nodeAdd(r)
	w.Header().Set("AMP-Redirect-To", "http://localhost:4000/")
	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin, AMP-Redirect-To")
	SendJsonResponse(w, n)
}
