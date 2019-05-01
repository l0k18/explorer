package exp

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.parallelcoin.io/marcetin/explorer/utl"
	"github.com/gorilla/mux"

	"git.parallelcoin.io/marcetin/explorer/jdb"
	"git.parallelcoin.io/marcetin/explorer/tpl"
)

func RPCSRC(c string) (rpc *RPCSource) {
	rpcSources, _ := jdb.JDB.ReadAll("data/" + c + "/rpcsrc")
	for _, rpcSrc := range rpcSources {
		if err := json.Unmarshal([]byte(rpcSrc), &rpc); err != nil {
			fmt.Println("Error", err)
		}
	}
	return
}

func Coins() (cs []Coin) {
	coins, _ := jdb.JDB.ReadAll("coins")
	for _, coinItem := range coins {
		var c Coin
		if err := json.Unmarshal([]byte(coinItem), &c); err != nil {
			fmt.Println("Error", err)
		}
		cs = append(cs, c)
	}
	return
}
func ViewAddress(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var addrs map[uint32]Addr
	jdb.JDB.Read("data/"+v["coin"]+"/addrs", "full", &addrs)
	addrHash := hash(v["addr"])
	addr := addrs[addrHash]
	// var a Addr
	txS := make(map[string]interface{})

	// for _, txID := range addr.Txs {
	// 	// tx := RPCSRC(v["coin"]).GetTx(txID)
	// 	// if tx["vout"] != nil {
	// 	// 	txS["vout"] = tx["vout"].([]interface{})
	// 	// }
	// 	// if tx["vin"] != nil {
	// 	// 	txS["vin"] = tx["vin"].([]interface{})
	// 	// }
	// 	// if tx["time"] != nil {
	// 	// 	txS["time"] = tx["time"]
	// 	// }
	// }
	a := map[string]interface{}{
		"d": []interface{}{addr, txS},
	}
	out, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var rawData map[uint32]interface{}
	jdb.JDB.Read("data/"+v["coin"]+"/"+v["cat"], v["data"], &rawData)
	data := map[string]interface{}{
		"coin":     v["coin"],
		"category": v["cat"],
		"d":        rawData,
	}
	fmt.Println("Error encdsfsdfsdfsdfsdfON", rawData)

	out, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(out))
}

func GetChart(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var rich map[uint32]Addr
	jdb.JDB.Read("data/"+v["coin"]+"/"+v["cat"], v["data"], &rich)
	data := Chart{
		Coin:     v["coin"],
		Chart:    v["chart"],
		Category: v["cat"],
		Data:     v["data"],
		CData:    rich,
		// Template: tpl.TPLHandler().Templates,
	}
	// fmt.Println("ssssssssssssssssssssssssasasas", data)
	tpl.TPLHandler().ExecuteTemplate(w, v["chart"]+"_gohtml", data)
}

func GetAddrsList(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	utl.SendJsonResponse(w, CreateAddrsList(v["coin"]))
}
