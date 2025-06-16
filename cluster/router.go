package cluster

import "go-Redis/interface/resp"

func makeRouter() map[string]CmdFunc {
	router := make(map[string]CmdFunc)
	router["exists"] = defaultFunc
	router["type"] = defaultFunc
	router["set"] = defaultFunc
	router["setnx"] = defaultFunc
	router["get"] = defaultFunc
	router["getset"] = defaultFunc
	//router["rename"]

	return router
}

// GET key // SET k1 v1
func defaultFunc(cluster *ClusterDatabase, c resp.Connection, cmdArgs [][]byte) resp.Reply {
	key := string(cmdArgs[1])
	peer := cluster.peerPicker.PickNode(key)
	return cluster.relay(peer, c, cmdArgs)
}
