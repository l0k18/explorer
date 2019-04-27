package exp

import (
	"fmt"
	"strconv"

	"git.parallelcoin.io/marcetin/explorer/utl"
)

// var _ JormNode = &Node{}

// type JormNode interface {
// 	GetBlockCount() int
// 	// GetBlockByHash(blockhash string) interface{}
// 	GetBlock(blockhash string) interface{}
// 	GetBlockTxAddr(blockheight int) interface{}
// 	GetBlockByHeight(blockheight int) interface{}
// 	GetTx(txid string) interface{}
// 	GetAddr(cs *CSystem, addr string) interface{}
// 	GetRawMemPool() interface{}
// 	GetMiningInfo() interface{}
// 	GetInfo() interface{}
// }

func (node *Node) GetBlockCount() (b int) {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	gbc, err := jrc.MakeRequest("getblockcount", bparams)
	if err != nil {
		fmt.Println("Error node call: ", err)

	}
	switch gbc.(type) {
	case float64:
		return int(gbc.(float64))
	case string:
		b, _ := strconv.Atoi(gbc.(string))
		return b
	default:
		b, _ := strconv.Atoi(gbc.(string))
		return b
	}
	return
}

func (node *Node) GetLastBlocks() *map[int]map[string]interface{} {
	lb := make(map[int]map[string]interface{})
	blockcount := SrcNode().GetBlockCount()
	minusblockcount := int(blockcount - 100)
	for ibh := minusblockcount; ibh <= blockcount; {
		var blk map[string]interface{}
		blk = (SrcNode().GetBlockByHeight(ibh)).(map[string]interface{})
		lb[ibh] = blk
		ibh++
	}
	return &lb
}

func (node *Node) GetBlock(blockhash string) interface{} {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []string{blockhash}
	block, err := jrc.MakeRequest("getblock", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Error", err)
	}
	return block
}

func (node *Node) GetBlockTxAddr(blockheight int) interface{} {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	//bh, err := strconv.Atoi(blockheight)
	//bparams := []int{bh}
	bparams := []int{blockheight}
	blockHash, err := jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block Tx Addr Error", err)
	}
	var block interface{}
	var txs []interface{}
	if blockHash != nil {
		block = node.GetBlock((blockHash).(string))
	}
	iblock := make(map[string]interface{})
	iblock = block.(map[string]interface{})

	itxs := iblock["tx"].([]interface{})
	//txs := itxs.([]string)

	for _, itx := range itxs {
		var txid string
		txid = itx.(string)

		verbose := int(1)
		var grtx []interface{}
		grtx = append(grtx, txid)
		grtx = append(grtx, verbose)
		rtx, err := jrc.MakeRequest("getrawtransaction", grtx)
		if err != nil {
			fmt.Println("Jorm Node Get Block Tx Addr Tx Error", err)
		}
		txs = append(txs, rtx)

	}

	//fmt.Println("blockblockblockblockblock", txs)

	blocktxaddr := map[string]interface{}{
		"b": block,
		"t": txs,
	}
	return blocktxaddr
}
func (node *Node) GetBlockByHeight(blockheight int) interface{} {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	//bh, err := strconv.Atoi(blockheight)
	//bparams := []int{bh}
	bparams := []int{blockheight}
	blockHash, err := jrc.MakeRequest("getblockhash", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Block By Height Error", err)
	}
	var block interface{}
	if blockHash != nil {
		block = node.GetBlock((blockHash).(string))
	}
	return block
}
func (node *Node) GetTx(txid string) interface{} {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	verbose := int(1)
	var grtx []interface{}
	grtx = append(grtx, txid)
	grtx = append(grtx, verbose)
	rtx, err := jrc.MakeRequest("getrawtransaction", grtx)
	if err != nil {
		fmt.Println("Jorm Node Get Tx Error", err)
	}
	// if rtx != nil {
	// 	rawtx = rtx.(Tx)
	// }
	return rtx
}

// func (node *Node) GetAddr(addr string) interface{} {

// aD := exJDB.EJDBGetAddr(addr)
// if aD.Addr == "" {
// 	return nil
// }
// 	 return aD
// }
func (node *Node) GetRawMemPool() interface{} {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getrawmempool", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Raw Mem Pool Error", err)
	}
	return get
}

func (node *Node) GetMiningInfo() interface{} {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getmininginfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Mining Info Error", err)
	}
	return get
}

func (node *Node) GetInfo() interface{} {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Info Error", err)
	}
	return get
}
func (node *Node) GetPeerInfo() interface{} {
	jrc := utl.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	if jrc == nil {
		fmt.Println("Error node status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getpeerinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}
