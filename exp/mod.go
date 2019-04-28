package exp

import (
	"encoding/json"
	"fmt"

	"git.parallelcoin.io/marcetin/explorer/amp"
	"git.parallelcoin.io/marcetin/explorer/jdb"
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
type Node struct {
	NodeID      string `json:"nodeid" form:"nodeid"`
	Coin        string `json:"coin" form:"coin"`
	Slug        string `json:"slug"`
	RPCUser     string `json:"rpcuser" form:"rpcuser"`
	RPCPassword string `json:"rpcpassword" form:"rpcpassword"`
	Address     string `json:"address" form:"address"`
	IP          string `json:"ip" form:"ip"`
	Port        int64  `json:"port" form:"port"`
}

type Block struct {
	BlockHeight int    `json:"blockheight"`
	BlockHash   string `json:"bhash"`
	Data        []byte `json:"b"`
}

type Blk struct {
	Block string `json:"b"`
}
type Tx struct {
	TxHash string `json:"txhash""`
	Data   []byte `json:"data"`
}
type Addr struct {
	Addr  string  `json:"nodeid" form:"addr"`
	Value float64 `json:"nodeid" form:"value"`
	Rank  int     `json:"nodeid" form:"rank"`
}
type Index struct {
	ID         string                   `json:"id"`
	LastBlocks []map[string]interface{} `json:"lb"`
	Coins      []Coin                   `json:"c"`
	Blocks     []Block                  `json:"b"`
	AMP        amp.AMP                  `json:"amp"`
	RichList   map[string][]Addr        `json:"rich"`
	// Node       *Node                    `json:"n"`
}

type BlVw struct {
	ID    string  `json:"id"`
	Block Block   `json:"block"`
	AMP   amp.AMP `json:"amp"`
}

type TxVw struct {
	ID  string  `json:"id"`
	Tx  Tx      `json:"tx"`
	AMP amp.AMP `json:"amp"`
}

type AdVw struct {
	ID   string  `json:"id"`
	Addr Addr    `json:"addr"`
	AMP  amp.AMP `json:"amp"`
}

func SrcNode(c string) (n *Node) {
	nodes, _ := jdb.JDB.ReadAll("data/" + c + "/nodes")
	for _, nodeItem := range nodes {
		fmt.Println("asassaaaaaaaa", c)
		if err := json.Unmarshal([]byte(nodeItem), &n); err != nil {
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
