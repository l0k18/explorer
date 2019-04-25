package rts

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"git.parallelcoin.io/marcetin/explorer/amp"
	"git.parallelcoin.io/marcetin/explorer/exp"
	"git.parallelcoin.io/marcetin/explorer/mod"
	"github.com/gorilla/mux"
)

var templates = make(map[string]*template.Template)

var last string

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("templates/index.gohtml", "templates/addnode.gohtml", "templates/spectre.gohtml", "templates/style.gohtml", "templates/base.gohtml", "templates/sidebar.gohtml"))
	templates["block"] = template.Must(template.ParseFiles("templates/block.gohtml", "templates/spectre.gohtml", "templates/style.gohtml", "templates/base.gohtml", "templates/sidebar.gohtml"))
	templates["blockhash"] = template.Must(template.ParseFiles("templates/blockhash.gohtml", "templates/spectre.gohtml", "templates/style.gohtml", "templates/base.gohtml", "templates/sidebar.gohtml"))
	templates["tx"] = template.Must(template.ParseFiles("templates/tx.gohtml", "templates/spectre.gohtml", "templates/style.gohtml", "templates/base.gohtml", "templates/sidebar.gohtml"))
	templates["addr"] = template.Must(template.ParseFiles("templates/addr.gohtml", "templates/spectre.gohtml", "templates/style.gohtml", "templates/base.gohtml", "templates/sidebar.gohtml"))
}

func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func respondWithJSON(w http.ResponseWriter, code int, block interface{}) {
// 	response, _ := json.Marshal(block)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	w.Write(response)
// }

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	node := exp.SrcNode()
	data := mod.Index{
		Node:   node,
		Blocks: []mod.Block{},
		AMP:    amp.AMPS(),
	}
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	w.Header().Set("AMP-Access-Control-Allow-Source-Origin", "http://localhost:4000")
	renderTemplate(w, "index", "base", data)
}

func viewBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.BlVw{
		ID:    id,
		Block: mod.Block{},
		AMP:   amp.AMPS(),
	}
	renderTemplate(w, "block", "base", data)
}
func viewBlockHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockhash := vars["blockhash"]
	data := mod.BlVw{
		ID:    blockhash,
		Block: mod.Block{},
		AMP:   amp.AMPS(),
	}
	renderTemplate(w, "blockhash", "base", data)
}
func viewTx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.TxVw{
		ID:  id,
		Tx:  mod.Tx{},
		AMP: amp.AMPS(),
	}
	renderTemplate(w, "tx", "base", data)
}
func viewAddr(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := mod.AdVw{
		ID:   id,
		Addr: mod.Addr{},
		AMP:  amp.AMPS(),
	}
	renderTemplate(w, "addr", "base", data)
}
func apiData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	tp := vars["type"]
	url := "http://194.135.89.49:8999/e/parallelcoin/" + tp + "/" + id
	data, _ := GetData(url)
	w.Write([]byte(data))
}
func apiLast(w http.ResponseWriter, r *http.Request) {
	url := "http://194.135.89.49:8999/e/parallelcoin/last"
	data, _ := GetData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func apiInfo(w http.ResponseWriter, r *http.Request) {
	url := "http://194.135.89.49:8999/e/parallelcoin/info"
	data, _ := GetData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func apiMiningInfo(w http.ResponseWriter, r *http.Request) {
	url := "http://194.135.89.49:8999/e/parallelcoin/gmi"
	data, _ := GetData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}
func apiRawPool(w http.ResponseWriter, r *http.Request) {
	url := "http://194.135.89.49:8999/e/parallelcoin/rmp"
	data, _ := GetData(url)
	//fmt.Println("blkblkblkblkblkblk", data)
	w.Write([]byte(data))
}

func doSearch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	search := r.FormValue("src")
	fmt.Println("searchsearchsearchsearchsearchsearchsearchsearch", search)

	tps := []string{"block", "blockhash", "tx", "addr"}
	var tpt string
	for _, tp := range tps {
		url := "http://194.135.89.49:8999/e/parallelcoin/" + tp + "/" + search
		fmt.Println("urlurlurlurlurlurlurlurlurlurlurl", url)
		gg, err := GetData(url)
		if err != nil {
			tpt = tp
		}
		fmt.Println("urlurlurlurlurlurlurlurlurlurlurl", gg)

	}

	http.Redirect(w, r, fmt.Sprintf("/"+tpt+"/"+search), http.StatusPermanentRedirect)
}

func GetData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	mapBody, err := ioutil.ReadAll(resp.Body)
	return mapBody, err
}
