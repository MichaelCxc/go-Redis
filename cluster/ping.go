package cluster

import "go-Redis/interface/resp"

func ping(cluster *ClusterDatabase, c resp.Connection, cmdArgs [][]byte) resp.Reply {
	return cluster.db.Exec(c, cmdArgs)
}
