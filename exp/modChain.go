package exp

import (

)


// type Block struct {
// 	Bits              string   `json:"bits"`
// 	Chainwork         string   `json:"chainwork"`
// 	Confirmations     int64    `json:"confirmations"`
// 	Difficulty        float64  `json:"difficulty"`
// 	Hash              string   `json:"hash"`
// 	Height            int64    `json:"height"`
// 	Mediantime        int64    `json:"mediantime"`
// 	Merkleroot        string   `json:"merkleroot"`
// 	NTx               int      `json:"ntx"`
// 	NextBlockHash     string   `json:"nextblockhash"`
// 	Nonce             uint32   `json:"nonce"`
// 	PreviousBlockHash string   `json:"previousblockhash"`
// 	Size              int64    `json:"size"`
// 	StrippedSize      int64    `json:"strippedsize"`
// 	Time              int64    `json:"time"`
// 	TX                []string `json:"tx"`
// 	Version           int64    `json:"version"`
// 	VersionHex        string   `json:"versionhex"`
// 	Weight            int64    `json:"weight"`
// }

// type BlockShort struct {
// 	Bits          string  `json:"bits"`
// 	Confirmations int64   `json:"confirmations"`
// 	Difficulty    float64 `json:"difficulty"`
// 	Hash          string  `json:"hash"`
// 	Height        int64   `json:"height"`
// 	NTx           int     `json:"ntx"`
// 	Size          int64   `json:"size"`
// 	Time          int64   `json:"time"`
// }

type Addr struct {
	Addr  string   `json:"addr" form:"addr"`
	Value float64  `json:"value" form:"value"`
 	Txs   []string `json:"txs" form:"txs"`
}


