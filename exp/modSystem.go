package exp

import (
	"git.parallelcoin.io/marcetin/explorer/amp"
)

type Config struct {
	Title string          `json:"title"`
	JDB   string          `json:"jdb"`
	Coins map[string]Coin `json:"coins"`
}
type Coin struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Abbr string `json:"abbr"`
	Icon string `json:"icon"`
	X    int64  `json:"x"`
}
type RPCSource struct {
	RPCSourceID string `json:"rpcsourceid" form:"rpcsourceid"`
	Coin        string `json:"coin" form:"coin"`
	Slug        string `json:"slug"`
	RPCUser     string `json:"rpcuser" form:"rpcuser"`
	RPCPassword string `json:"rpcpassword" form:"rpcpassword"`
	IP          string `json:"ip" form:"ip"`
	Port        int64  `json:"port" form:"port"`
}


type Index struct {
	ID         string                   `json:"id"`
	LastBlocks []map[string]interface{} `json:"lb"`
	Coins      []Coin                   `json:"c"`
	// Blocks     []Block                  `json:"b"`
	AMP      amp.AMP           `json:"amp"`
	RichList map[string][]Addr `json:"rich"`
	// Node       *Node                    `json:"n"`
}


type Chart struct {
	Coin     string      `json:"coin"`
	Chart    string      `json:"chart"`
	Category string      `json:"category"`
	Data     string      `json:"data"`
	CData    interface{} `json:"data"`
}
