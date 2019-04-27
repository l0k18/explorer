package rts

import (
	"fmt"
	"net/http"
	"sort"

	"git.parallelcoin.io/marcetin/explorer/amp"
	"git.parallelcoin.io/marcetin/explorer/exp"
	"git.parallelcoin.io/marcetin/explorer/jdb"
	"git.parallelcoin.io/marcetin/explorer/mod"
	"git.parallelcoin.io/marcetin/explorer/tpl"
)

// var templates = make(map[string]*template.Template)

var last string

// var tpls = template.Must(template.ParseFiles("templates/index.gohtml", "templates/addnode.gohtml", "templates/block.gohtml", "templates/blockhash.gohtml", "templates/tx.gohtml", "templates/addr.gohtml", "templates/addr.gohtml", "templates/spectre.gohtml", "templates/style.gohtml", "templates/base.gohtml", "templates/sidebar.gohtml"))

// func renderTemplate(w http.ResponseWriter, viewModel interface{}) {
// 	err := tpls.ExecuteTemplate(w, "base", viewModel)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func respondWithJSON(w http.ResponseWriter, code int, block interface{}) {
// 	response, _ := json.Marshal(block)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	node := exp.SrcNode()
	var rich = make(map[string]float64)
	var richFull = make(map[string]float64)
	jdb.JDB.Read("addrs", "addrs", &richFull)

	for r, m := range richFull {
		if m < 10 {
			rich["10"] += m
		}
		if m < 100 {
			rich["100"] += m
		}
		if m < 200 {
			rich["200"] += m
		}
		if m < 300 {
			rich["300"] += m
		}
		if m < 500 {
			rich["500"] += m
		}
		if m < 1000 {
			rich["1000"] += m
		}
		if m < 3000 {
			rich["3000"] += m
		}
		if m > 3000 {
			// rich["3000"] += m
			rich[r] = m
		}

	}

	keys := make([]string, 0, len(rich))
	for k, r := range rich {
		fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-k", k)
		fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-r.-.", r)
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for r, k := range keys {
		// fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.", rich)
		fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-k", k)
		fmt.Println("-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-r.-.", r)
		fmt.Println(k, rich[k])
	}

	data := mod.Index{
		LastBlocks: *exp.SrcNode().GetLastBlocks(),
		Node:       node,
		Blocks:     []mod.Block{},
		AMP:        amp.AMPS(),
		Rich:       rich,
	}
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	// w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://localhost:4000")

	tpl.TPLHandler().ExecuteTemplate(w, "base_gohtml", data)
	// t.ExecuteTemplate(w, "base_html", nil)
	// renderTemplate(w, data)
}

// func viewBlock(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	data := mod.BlVw{
// 		ID:    id,
// 		Block: mod.Block{},
// 		AMP:   amp.AMPS(),
// 	}
// 	tpl.TPLHandler().ExecuteTemplate(w, "block_gohtml", data)
// }
// func viewBlockHash(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	blockhash := vars["blockhash"]
// 	data := mod.BlVw{
// 		ID:    blockhash,
// 		Block: mod.Block{},
// 		AMP:   amp.AMPS(),
// 	}
// 	tpl.TPLHandler().ExecuteTemplate(w, "hash_gohtml", data)
// }
// func viewTx(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	data := mod.TxVw{
// 		ID:  id,
// 		Tx:  mod.Tx{},
// 		AMP: amp.AMPS(),
// 	}
// 	tpl.TPLHandler().ExecuteTemplate(w, "tx_gohtml", data)

// }
// func viewAddr(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]
// 	data := mod.AdVw{
// 		ID:   id,
// 		Addr: mod.Addr{},
// 		AMP:  amp.AMPS(),
// 	}
// 	tpl.TPLHandler().ExecuteTemplate(w, "addr_gohtml", data)
// }

// // func GetData(url string) ([]byte, error) {
// // 	resp, err := http.Get(url)
// // 	if err != nil {
// // 	}
// // 	defer resp.Body.Close()
// // 	mapBody, err := ioutil.ReadAll(resp.Body)
// // 	return mapBody, err
// // }
