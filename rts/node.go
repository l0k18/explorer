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

	// r.ParseForm()

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
	SendJsonResponse(w, exp.Node{})
	http.Redirect(w, r, "/", 302)
}

func isUserTryingTheInputTextErrorDemo(name string) bool {
	return name == "error"
}
