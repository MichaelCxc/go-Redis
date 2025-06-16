package cluster

import (
	"go-Redis/interface/resp"
	"go-Redis/resp/reply"
)

func Rename(cluster *ClusterDatabase, c resp.Connection, cmdArgs [][]byte) resp.Reply {
	if len(cmdArgs) != 3 {
		return reply.MakeErrReply("Err Wrong number args")
	}

	src := string(cmdArgs[1])
	dst := string(cmdArgs[2])
	peerSource := cluster.peerPicker.PickNode(src)
	peerTarget := cluster.peerPicker.PickNode(dst)
	//Just implement standalone DB now
	if peerSource != peerTarget {
		return reply.MakeErrReply("Err rename must within on peer")
	}

	return cluster.relay(peerSource, c, cmdArgs)
}
