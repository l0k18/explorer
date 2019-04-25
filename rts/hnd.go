package rts

import (
	"net/http"

	"git.parallelcoin.io/marcetin/explorer/amp"
	"git.parallelcoin.io/marcetin/explorer/exp"
	"git.parallelcoin.io/marcetin/explorer/mod"
	"git.parallelcoin.io/marcetin/explorer/tpl"
	"github.com/gorilla/mux"
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
	data := mod.Index{
		ID:     "Explorer",
		Node:   node,
		Blocks: []mod.Block{},
		AMP:    amp.AMPS(),
	}
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	// w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://localhost:4000")

	tpl.TPLHandler().ExecuteTemplate(w, "base_gohtml", data)
	// t.ExecuteTemplate(w, "base_html", nil)
	// renderTemplate(w, data)
}

func viewBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.BlVw{
		ID:    id,
		Block: mod.Block{},
		AMP:   amp.AMPS(),
	}
	tpl.TPLHandler().ExecuteTemplate(w, "base_gohtml", data)
}
func viewBlockHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockhash := vars["blockhash"]
	data := mod.BlVw{
		ID:    blockhash,
		Block: mod.Block{},
		AMP:   amp.AMPS(),
	}
	tpl.TPLHandler().ExecuteTemplate(w, "base_gohtml", data)
}
func viewTx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.TxVw{
		ID:  id,
		Tx:  mod.Tx{},
		AMP: amp.AMPS(),
	}
	tpl.TPLHandler().ExecuteTemplate(w, "base_gohtml", data)

}
func viewAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.AdVw{
		ID:   id,
		Addr: mod.Addr{},
		AMP:  amp.AMPS(),
	}
	tpl.TPLHandler().ExecuteTemplate(w, "base_gohtml", data)
}

// func GetData(url string) ([]byte, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 	}
// 	defer resp.Body.Close()
// 	mapBody, err := ioutil.ReadAll(resp.Body)
// 	return mapBody, err
// }
