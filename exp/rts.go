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
	var lastblocks []interface{}
	// node := Node{}
	// fmt.Println("nodenodenodenode", aNode)

	blockcount := SrcNode().JNGetBlockCount()
	minusblockcount := int(blockcount - 20)
	for ibh := minusblockcount; ibh <= blockcount; {
		//ib := strconv.Itoa(ibh)
		blk := SrcNode().JNGetBlockByHeight(ibh)
		lastblocks = append(lastblocks, blk)
		ibh++
	}
	lb := map[string]interface{}{
		"d": lastblocks,
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
	lastblock := SrcNode().JNGetBlockCount()

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
	block := SrcNode().JNGetBlockByHeight(bhi)
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
	block := SrcNode().JNGetBlockTxAddr(bhi)
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

	block := SrcNode().JNGetBlock(bh)
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

func Tx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txid := vars["txid"]
	// node := Node{}

	tX := SrcNode().JNGetTx(txid)

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
	rawMemPool := SrcNode().JNGetRawMemPool()
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
	miningInfo := SrcNode().JNGetMiningInfo()

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
	info := SrcNode().JNGetInfo()

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
	info := SrcNode().JNGetPeerInfo()
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

	// aD := node.JNGetAddr(addr)

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
