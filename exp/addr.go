package exp

import (
	"fmt"

	"git.parallelcoin.io/marcetin/explorer/jdb"
)

func (n *Node) GetAddrs() {
	var AVS = make(map[string]float64)
	last := SrcNode(n.Coin).GetBlockCount()
	for block := 0; block <= last; {
		b := (SrcNode(n.Coin).GetBlockByHeight(block)).(map[string]interface{})
		txs := (b["tx"]).([]interface{})
		for _, tx := range txs {
			t := tx.(string)
			xt := SrcNode(n.Coin).GetTx(string(t))
			if xt != nil {
				txm := xt.(map[string]interface{})
				vout := txm["vout"].([]interface{})
				for _, v := range vout {
					spk := v.(map[string]interface{})
					spkmap := spk["scriptPubKey"].(map[string]interface{})
					if spkmap["addresses"] != nil {
						addrs := spkmap["addresses"].([]interface{})
						for _, addr := range addrs {
							AVS[addr.(string)] += spk["value"].(float64)
							fmt.Println("Address:", addr)
							fmt.Println("Value Sum:", AVS[addr.(string)])
							fmt.Println("Block:", block)
						}
					}

				}
			}
		}
		block++
	}
	jdb.JDB.Write("data/"+n.Coin+"/addrs", "addrs", AVS)
}
