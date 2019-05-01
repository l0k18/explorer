package exp

import (
	"fmt"

	"git.parallelcoin.io/marcetin/explorer/utl"
)

func (rpc *RPCSource) GetRawMemPool() interface{} {
	jrc := utl.NewClient(rpc.RPCUser, rpc.RPCPassword, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getrawmempool", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Raw Mem Pool Error", err)
	}
	return get
}

func (rpc *RPCSource) GetMiningInfo() interface{} {
	jrc := utl.NewClient(rpc.RPCUser, rpc.RPCPassword, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getmininginfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Mining Info Error", err)
	}
	return get
}

func (rpc *RPCSource) GetInfo() interface{} {
	jrc := utl.NewClient(rpc.RPCUser, rpc.RPCPassword, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Info Error", err)
	}
	return get
}
func (rpc *RPCSource) GetPeerInfo() interface{} {
	jrc := utl.NewClient(rpc.RPCUser, rpc.RPCPassword, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getpeerinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}

func (rpc *RPCSource) AddNode(ip string) interface{} {
	jrc := utl.NewClient(rpc.RPCUser, rpc.RPCPassword, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}

	bparams := []string{ip, "add"}
	get, err := jrc.MakeRequest("addnode", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}

func (rpc *RPCSource) GetAddNodeInfo(ip string) interface{} {
	jrc := utl.NewClient(rpc.RPCUser, rpc.RPCPassword, rpc.IP, rpc.Port)
	if jrc == nil {
		fmt.Println("Error n status write")
	}
	bparams := []int{}
	get, err := jrc.MakeRequest("getaddednodeinfo", bparams)
	if err != nil {
		fmt.Println("Jorm Node Get Peer Info Error", err)
	}
	return get
}
