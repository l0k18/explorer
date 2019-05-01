package exp

import (
	"fmt"
	"hash/fnv"

	"git.parallelcoin.io/marcetin/explorer/jdb"
)

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func CreateAddrsList(coin string) map[uint32]Addr {
	a := make(map[uint32]Addr)
	last := RPCSRC(coin).GetBlockCount()
	// last := 1111
	for block := 0; block <= last; {
		b := (RPCSRC(coin).GetBlockByHeight(block))
		txs := b["tx"].([]string)
		for _, t := range txs {
			if block > 0 {
				txm := RPCSRC(coin).GetTx(string(t))
				if txm != nil {
					vout := txm["vout"].([]interface{})
					for _, v := range vout {
						spk := v.(map[string]interface{})
						spkmap := spk["scriptPubKey"].(map[string]interface{})
						if spkmap["addresses"] != nil {
							addrs := spkmap["addresses"].([]interface{})
							var addr Addr
							for _, ad := range addrs {
								id := hash(ad.(string))
								a[id] = Addr{
									Addr:  ad.(string),
									Value: (a[id].Value + spk["value"].(float64)),
									Txs:   append(addr.Txs, t),
								}
								fmt.Println("Address:", a[id])
								fmt.Println("Value Sum:", a[id].Value)
								fmt.Println("Block:", block)
							}
						}
					}
				}
			}
		}
		block++
	}
	jdb.JDB.Write("data/"+coin+"/addrs", "full", a)
	return a
}

// func CreateRichlist(a []Addr) map[uint32]Addr {
// 	var rich = make(map[string]Addr)
// 	for _, m := range a {
// 		switch {
// 		case m.Value <= 10:
// 			rich["Less than 10"] = Addr{
// 				Value: rich["Less than 10"].Value + m.Value,
// 			}
// 		case m.Value <= 100:
// 			rich["Less than 100"] = Addr{
// 				Value: rich["Less than 100"].Value + m.Value,
// 			}
// 		case m.Value <= 1000:
// 			rich["Less than 1000"] = Addr{
// 				Value: rich["Less than 1000"].Value + m.Value,
// 			}
// 		case m.Value <= 3000:
// 			rich["Less than 3000"] = Addr{
// 				Value: rich["Less than 3000"].Value + m.Value,
// 			}
// 		case m.Value <= 5000:
// 			rich["Less than 5000"] = Addr{
// 				Value: rich["Less than 5000"].Value + m.Value,
// 			}
// 		case m.Value <= 10000:
// 			rich["Less than 10000"] = Addr{
// 				Value: rich["Less than 10000"].Value + m.Value,
// 			}
// 		case m.Value >= 10000:
// 			rich[m.Addr] = Addr{
// 				Addr:  m.Addr,
// 				Value: m.Value,
// 				Txs:   m.Txs,
// 			}
// 		}
// 	}
// 	richListToSort := []Addr{}
// 	i := 1
// 	for k, r := range rich {
// 		richListToSort = append(richListToSort, Addr{k, r.Value, r.Txs})
// 		i++
// 	}
// 	sort.Sort(sort.Reverse(ByValue(richListToSort)))
// 	richListSort := []Addr{}
// 	ii := 1
// 	for _, r := range richListToSort {
// 		richListSort = append(richListSort, Addr{r.Addr, r.Value, ii, r.Txs})
// 		ii++
// 	}
// 	// jdb.JDB.Write("data/"+coin+"/richlist", "full", richListSort)
// 	// jdb.JDB.Write("data/"+coin+"/richlist", "singles", richListSingles(richListToSort))
// 	// jdb.JDB.Write("data/"+coin+"/richlist", "grouped", richListGrouped(richListToSort))
// 	return a
// }

type ByValue []Addr

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Less(i, j int) bool { return a[i].Value < a[j].Value }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func richListSingles(l []Addr) (f []Addr) {
	i := 1
	for _, r := range l {
		if len(r.Addr) > 20 {
			f = append(f, Addr{r.Addr, r.Value, r.Txs})
		}
		i++
	}
	return f
}

func richListGrouped(l []Addr) (f []Addr) {
	i := 1
	for _, r := range l {
		if len(r.Addr) < 20 {
			f = append(f, Addr{r.Addr, r.Value, r.Txs})
		}
		i++
	}
	return f
}
