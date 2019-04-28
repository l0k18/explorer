package rts

import (
	"net/http"
	"sort"

	"git.parallelcoin.io/marcetin/explorer/amp"
	"git.parallelcoin.io/marcetin/explorer/exp"
	"git.parallelcoin.io/marcetin/explorer/jdb"
	"git.parallelcoin.io/marcetin/explorer/tpl"
)

type ByValue []exp.Addr

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func richListSingles(l []exp.Addr) (f []exp.Addr) {
	i := 1
	for _, r := range l {
		if len(r.Addr) > 20 {
			f = append(f, exp.Addr{r.Addr, r.Value, i})
		}
		i++
	}
	return f
}
func richListGrouped(l []exp.Addr) (f []exp.Addr) {
	i := 1
	for _, r := range l {
		if len(r.Addr) < 20 {
			f = append(f, exp.Addr{r.Addr, r.Value, i})
		}
		i++
	}
	return f
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var rich = make(map[string]float64)
	var richFullGet = make(map[string]float64)
	jdb.JDB.Read("addrs", "addrs", &richFullGet)
	for r, m := range richFullGet {
		if m < 10 {
			rich["Less than 10"] += m
		}
		if m < 100 {
			rich["Less than 100"] += m
		}
		if m < 1000 {
			rich["Less than 1000"] += m
		}
		if m < 3000 {
			rich["Less than 3000"] += m
		}
		if m < 5000 {
			rich["Less than 5000"] += m
		}
		if m < 10000 {
			rich["Less than 10000"] += m
		}
		if m > 10000 {
			rich["More than 10000"] += m
			rich[r] = m
		}
	}
	richListToSort := []exp.Addr{}
	i := 1
	for k, r := range rich {
		richListToSort = append(richListToSort, exp.Addr{k, r, i})
		i++
	}
	sort.Sort(sort.Reverse(ByValue(richListToSort)))
	richListSort := []exp.Addr{}
	ii := 1
	for _, r := range richListToSort {
		richListSort = append(richListSort, exp.Addr{r.Addr, r.Value, ii})
		ii++
	}
	richList := map[string][]exp.Addr{
		"full":    richListSort,
		"singles": richListSingles(richListToSort),
		"grouped": richListGrouped(richListToSort),
	}

	data := exp.Index{
		Coins: exp.Coins(),
		// Blocks:   []exp.Block{},
		AMP:      amp.AMPS(),
		RichList: richList,
	}
	tpl.TPLHandler().ExecuteTemplate(w, "base_gohtml", data)
}
