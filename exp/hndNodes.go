package exp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"git.parallelcoin.io/marcetin/explorer/jdb"
	"github.com/gorilla/mux"
	"github.com/ip2location/ip2location-go"
)

func GetNodeStatus(w http.ResponseWriter, r *http.Request) {
	// nd := cs.MainJDB.MJDBGetAll("nodes")
	// for nd.Next() {
	// 	var node Node
	// 	err := nd.Decode(&node)
	// 	if err != nil {
	// 	}
	// 	if node.BitNode {
	// 		c := tools.NewClient(node.RPCUser, node.RPCPassword, node.IP, node.Port)
	// 		//c := tools.NewClient("duo", "pass", "127.0.0.1", 11066)
	// 		if c == nil {
	// 			fmt.Println("Error node status write")
	// 		}
	// 		params := []string{}
	// 		getinfo, err := c.MakeRequest("getinfo", params)
	// 		if err != nil {
	// 			fmt.Println("Error node status getinfo", err)
	// 		}
	// 		getpeerinfo, err := c.MakeRequest("getpeerinfo", params)
	// 		if err != nil {
	// 			fmt.Println("Error node status getpeerinfo", err)
	// 		}
	// 		// fmt.Println("getinfo", getinfo)
	// 		var nodec Node
	// 		nodec = node
	// 		nodec.GetInfo = getinfo
	// 		nodec.GetPeerInfo = getpeerinfo
	// 		cs.MainJDB.MJDBSet("nodes", node.GPSID, nodec)
	// 	}
	// }
}

func GetNodes(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	rpcSource, err := jdb.JDB.ReadAll("data/" + v["coin"] + "/rpcsrc")
	if err != nil {
		fmt.Println("Error", err)
	}
	nodesList := make(map[string]Node)
	errr := jdb.JDB.Read("data/"+v["coin"], "nodes", &nodesList)
	if errr != nil {
		fmt.Println("Error", errr)
	}
	for _, nd := range rpcSource {
		var node Node
		if err := json.Unmarshal([]byte(nd), &node); err != nil {
			fmt.Println("Error", err)
		}
		var peers []interface{}
		fmt.Println("Load Nodes Direct Connect", node.IP)
		gpeers := RPCSRC(v["coin"]).GetPeerInfo()
		if gpeers != nil {
			switch gpeers.(type) {
			case []interface{}:
				peers = gpeers.([]interface{})
				var peer map[string]interface{}
				for _, gpeer := range peers {
					peer = gpeer.(map[string]interface{})
					peerAddress := peer["addr"].(string)
					if node.IP != peerAddress {
						splitted := strings.Split(peerAddress, ":")
						peerPort := splitted[len(splitted)-1]
						var peerIP string
						for i := range splitted {
							if i == len(splitted)-1 {
								break
							}
							peerIP += splitted[i]
							if i < len(splitted)-2 {
								peerIP += ":"
							}
							pP, err := strconv.ParseInt(peerPort, 10, 64)
							if err != nil {
							}

							RPCSRC(v["coin"]).AddNode(peerIP)
							var nodec Node
							nodec.Coin = node.Coin
							nodec.IP = peerIP
							nodec.Port = pP
							nodesList[peerIP] = nodec
							addNodesRaw := RPCSRC(v["coin"]).GetAddNodeInfo(nodec.IP)
							addNodes := addNodesRaw.([]interface{})
							for _, addNode := range addNodes {
								an := addNode.(map[string]interface{})
								nodesList[an["addednode"].(string)] = Node{
									Coin: v["coin"],
									IP:   an["addednode"].(string),
								}

							}

							// fmt.Println("Load NodesCCCCCCCCCCCCCCCCC::", nodesList)

							// cs.MainJDB.MJDBSet("nodes", GPSID, nodec)
							// jdb.JDB.Write("nodes", nodec.Slug, nodec)
							jdb.JDB.Write("data/"+v["coin"], "nodes", &nodesList)
						}
					}
				}
			}
		}
	}
}

// func GetNodesByCoin(coin string) (nodes []Node) {
// 	cnodes, err := jdb.JDB.Read("data/coin/nodes")
// 	if err != nil {
// 		fmt.Println("Error", err)
// 	}
// 	for _, nd := range cnodes {
// 		var node Node
// 		if err := json.Unmarshal([]byte(nd), &node); err != nil {
// 			fmt.Println("Error", err)
// 		}
// 		if node.Coin == coin {
// 			nodes = append(nodes, node)
// 		}
// 	}
// 	return
// }

func NodesMap(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	var nodesIPs []NodeMap
	nodesList := make(map[string]Node)
	errr := jdb.JDB.Read("data/"+v["coin"], "nodes", &nodesList)
	if errr != nil {
		fmt.Println("Error", errr)
	}
	for _, node := range nodesList {
		var NodeMap NodeMap
		ip2location.Open("./utl/IP2LOCATION-LITE-DB11.BIN")
		results := ip2location.Get_all(node.IP)
		NodeMap.Coin = v["coin"]
		NodeMap.IP = node.IP
		NodeMap.Country_short = results.Country_short
		NodeMap.Country_long = results.Country_long
		NodeMap.Region = results.Region
		NodeMap.City = results.City
		NodeMap.Latitude = results.Latitude
		NodeMap.Longitude = results.Longitude
		NodeMap.Zipcode = results.Zipcode
		NodeMap.Timezone = results.Timezone
		ip2location.Close()
		nodesIPs = append(nodesIPs, NodeMap)
	}
	mapNodes := map[string]interface{}{
		"nodes": nodesIPs,
	}
	nodes, err := json.Marshal(mapNodes)
	if err != nil {
		fmt.Println("Error encoding JSON")
		return
	}
	w.Write([]byte(nodes))
}

// map[
// 	addnode:false
// 	addr:13.127.94.172:8333
// 	addrbind:10.0.0.19:24854
// 	addrlocal:212.62.35.158:56449
// 	banscore:0
// 	bytesrecv:6.532931e+06
// 	bytesrecv_per_msg:
// 		map[
// 				addr:34892
// 				cmpctblock:17439
// 				feefilter:32
// 				getdata:13934
// 				getheaders:1053
// 				headers:2756
// 				inv:1.100541e+06
// 				notfound:12829
// 				ping:3552
// 				pong:3552
// 				reject:156
// 				sendcmpct:66
// 				sendheaders:24
// 				tx:5.341954e+06
// 				verack:24
// 				version:127]
// 				bytessent:2.076168e+06
// 				bytessent_per_msg:
// 					map[
// 							addr:8375
// 							feefilter:32
// 							getaddr:24
// 							getdata:528753
// 							getheaders:1053
// 							headers:954
// 							inv:1.468165e+06
// 							notfound:4092
// 							ping:3552
// 							pong:3552
// 							reject:471 sendcmpct:66 sendheaders:24 tx:56905
// 							verack:24
// 							version:126] conntime:1.556680757e+09 id:154
// 							inbound:false
// 							inflight:[]
// 							lastrecv:1.55669407e+09 lastsend:1.556694073e+09
// 							minping:0.141898 pingtime:0.146475
// 							relaytxes:true
// 							services:000000000000000d
// 							startingheight:574025 subver:/Satoshi:0.14.99/
// 							synced_blocks:574051 synced_headers:574051
// 							timeoffset:-18
// 							version:70015 whitelisted:false]
