package exp

import (
	"encoding/json"
	"fmt"

	"git.parallelcoin.io/marcetin/explorer/jdb"
)

type Node struct {
	NodeID      string `json:"nodeid" form:"nodeid"`
	Slug        string `json:"slug"`
	RPCUser     string `json:"rpcuser" form:"rpcuser"`
	RPCPassword string `json:"rpcpassword" form:"rpcpassword"`
	Address     string `json:"address" form:"address"`
	IP          string `json:"ip" form:"ip"`
	Port        int64  `json:"port" form:"port"`
}

func SrcNode() (n *Node) {
	nodes, _ := jdb.JDB.ReadAll("nodes")
	for _, nodeItem := range nodes {
		if err := json.Unmarshal([]byte(nodeItem), &n); err != nil {
			fmt.Println("Error", err)
		}
	}
	return
}
