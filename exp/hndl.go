package exp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ViewCoins(w http.ResponseWriter, r *http.Request) {
	c := map[string]interface{}{
		"d": map[string]interface{}{
			"coins": Coins(),
		},
	}
	out, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewBlocks(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	per, _ := strconv.Atoi(v["per"])
	page, _ := strconv.Atoi(v["page"])
	lb := map[string]interface{}{
		"d": map[string]interface{}{
			"currentPage": page,
			"pageCount":   RPCSRC(v["coin"]).GetBlockCount() / per,
			"blocks":      RPCSRC(v["coin"]).GetBlocks(per, page),
		},
	}

	out, err := json.Marshal(lb)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ViewBlock(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	lastblock := RPCSRC(v["coin"]).GetBlockCount()
	bl := map[string]interface{}{
		"d": lastblock,
	}
	out, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ViewBlockHeight(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	bh := v["blockheight"]
	// node := Node{}
	bhi, _ := strconv.Atoi(bh)
	block := RPCSRC(v["coin"]).GetBlockByHeight(bhi)
	bl := map[string]interface{}{
		"d": block,
	}
	out, err := json.Marshal(bl)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func ViewHash(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	bh := v["blockhash"]
	block := RPCSRC(v["coin"]).GetBlock(bh)
	h := block["height"].(string)
	http.Redirect(w, r, "/a/block/"+h, 301)
}

func ViewTx(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	txid := v["txid"]
	// node := Node{}

	tX := RPCSRC(v["coin"]).GetTx(txid)

	tx := map[string]interface{}{
		"d": tX,
	}
	out, err := json.Marshal(tx)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewRawMemPool(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rawMemPool := RPCSRC(v["coin"]).GetRawMemPool()
	rmp := map[string]interface{}{
		"d": rawMemPool,
	}
	out, err := json.Marshal(rmp)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewMiningInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	miningInfo := RPCSRC(v["coin"]).GetMiningInfo()

	mi := map[string]interface{}{
		"d": miningInfo,
	}
	out, err := json.Marshal(mi)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewInfo(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := RPCSRC(v["coin"]).GetInfo()

	in := map[string]interface{}{
		"d": info,
	}
	out, err := json.Marshal(in)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
func ViewPeers(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	info := RPCSRC(v["coin"]).GetPeerInfo()
	pi := map[string]interface{}{
		"d": info,
	}
	out, err := json.Marshal(pi)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
