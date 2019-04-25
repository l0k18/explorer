package rts

import (
	"fmt"
	"net/http"
	"strconv"

	"git.parallelcoin.io/marcetin/explorer/exp"
	"git.parallelcoin.io/marcetin/explorer/jdb"
	"git.parallelcoin.io/marcetin/explorer/utl"
)

func NodeAdd(w http.ResponseWriter, r *http.Request) {
	nodeid := r.FormValue("nodeid")
	rpcuser := r.FormValue("rpcuser")
	rpcpassword := r.FormValue("rpcpassword")
	ip := r.FormValue("ip")
	port, _ := strconv.ParseInt(r.FormValue("port"), 10, 64)
	slug := utl.MakeSlug(nodeid)
	var node = exp.Node{
		NodeID:      nodeid,
		Slug:        slug,
		RPCUser:     rpcuser,
		RPCPassword: rpcpassword,
		IP:          ip,
		Port:        port,
	}
	jdb.JDB.Write("nodes", slug, node)
	fmt.Println("asssssss", node)
	w.Header().Set("AMP-Redirect-To", "http://localhost:4000/")
	w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin, AMP-Redirect-To")
	SendJsonResponse(w, node)
}
