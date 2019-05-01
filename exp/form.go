package exp

import (
	"fmt"
	"net/http"
	"strconv"

	"git.parallelcoin.io/marcetin/explorer/jdb"
	"git.parallelcoin.io/marcetin/explorer/utl"
)

func AddCoin(w http.ResponseWriter, r *http.Request) {
	c := coinAdd(r)
	w.Header().Set("AMP-Redirect-To", "http://localhost:4000/")
	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin, AMP-Redirect-To")
	utl.SendJsonResponse(w, c)
}
func coinAdd(r *http.Request) (c Coin) {
	c.Name = r.FormValue("name")
	c.Slug = utl.MakeSlug(c.Name)
	c.Abbr = r.FormValue("abbr")
	c.X, _ = strconv.ParseInt(r.FormValue("x"), 10, 64)
	jdb.JDB.Write("coins/", c.Slug, c)
	return c
}
func rpcSourceAdd(r *http.Request) (rpc RPCSource) {
	rpc.Coin = r.FormValue("coin")
	fmt.Println("asassaaaablockCountblockCountaaaa", rpc.Coin)
	rpc.RPCSourceID = r.FormValue("rpcsrcid")
	rpc.Slug = utl.MakeSlug(rpc.RPCSourceID)
	rpc.RPCUser = r.FormValue("rpcuser")
	rpc.RPCPassword = r.FormValue("rpcpassword")
	rpc.IP = r.FormValue("ip")
	rpc.Port, _ = strconv.ParseInt(r.FormValue("port"), 10, 64)
	jdb.JDB.Write("data/"+rpc.Coin+"/rpcsrc", rpc.Slug, rpc)
	return rpc
}

func AddRPCSource(w http.ResponseWriter, r *http.Request) {
	n := rpcSourceAdd(r)
	w.Header().Set("AMP-Redirect-To", "http://localhost:4000/")
	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin, AMP-Redirect-To")
	utl.SendJsonResponse(w, n)
}
