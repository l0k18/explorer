package rts

import (
	"net/http"

	"git.parallelcoin.io/marcetin/explorer/amp"
	"git.parallelcoin.io/marcetin/explorer/exp"
	"git.parallelcoin.io/marcetin/explorer/tpl"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := exp.Index{
		Coins: exp.Coins(),
		// Blocks:   []exp.Block{},
		AMP: amp.AMPS(),
		// RichList: richList,
	}
	tpl.TPLHandler().ExecuteTemplate(w, "base_gohtml", data)
}
