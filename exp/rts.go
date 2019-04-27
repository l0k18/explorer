package exp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// var node = Node{
// 	NodeID:      "xxxx",
// 	Slug:        "xxxx",
// 	RPCUser:     "user",
// 	RPCPassword: "pass",
// 	IP:          "127.0.0.1",
// 	Port:        11122,
// 	Published:   true,
// }
// var aNode = aNode()

func LastBlock(w http.ResponseWriter, r *http.Request) {
	lb := map[string]interface{}{
		"d": SrcNode().GetLastBlocks(),
	}
	out, err := json.Marshal(lb)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func Block(w http.ResponseWriter, r *http.Request) {
	// node := Node{}
	lastblock := SrcNode().GetBlockCount()

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

func BlockHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bh := vars["blockheight"]
	// node := Node{}
	bhi, _ := strconv.Atoi(bh)
	block := SrcNode().GetBlockByHeight(bhi)
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

func BHeight(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bh := vars["blockheight"]
	// node := Node{}

	bhi, _ := strconv.Atoi(bh)
	block := SrcNode().GetBlockTxAddr(bhi)
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

func Hash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bh := vars["blockhash"]
	// node := Node{}

	block := SrcNode().GetBlock(bh)
	b := block.(map[string]interface{})
	h := b["height"].(string)
	http.Redirect(w, r, "/a/block/"+h, 301)
}

func Tx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txid := vars["txid"]
	// node := Node{}

	tX := SrcNode().GetTx(txid)

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
func RawMemPool(w http.ResponseWriter, r *http.Request) {
	// node := Node{}
	rawMemPool := SrcNode().GetRawMemPool()
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
func MiningInfo(w http.ResponseWriter, r *http.Request) {
	// node := Node{}
	miningInfo := SrcNode().GetMiningInfo()

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
func Info(w http.ResponseWriter, r *http.Request) {
	// node := Node{}
	info := SrcNode().GetInfo()

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
func Peers(w http.ResponseWriter, r *http.Request) {
	// node := Node{}
	info := SrcNode().GetPeerInfo()
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
func Address(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)

	// addr := vars["addr"]
	// node := GetBitNodes(coin)

	// aD := node.GetAddr(addr)

	// fmt.Println("343434343444", aD)
	// sty, err := json.Marshal(aD)
	// if err != nil {
	// 	fmt.Println(string(sty), err.Error())
	// }

	address := map[string]interface{}{
		// "d": aD,
	}
	out, err := json.Marshal(address)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}
