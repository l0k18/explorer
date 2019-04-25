package mod

import (
	"git.parallelcoin.io/marcetin/explorer/amp"
	"git.parallelcoin.io/marcetin/explorer/exp"
)

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
	Addr string `json:"addr"`
}

type Index struct {
	Node   *exp.Node `json:"n"`
	Blocks []Block   `json:"b"`
	AMP    amp.AMP   `json:"amp"`
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
