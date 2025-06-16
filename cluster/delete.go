package cluster

import (
	"go-Redis/interface/resp"
	"go-Redis/resp/reply"
)

//del k1 k2 k3 k4 k5

func Del(cluster *ClusterDatabase, c resp.Connection, cmdArgs [][]byte) resp.Reply {
	replies := cluster.broadcast(c, cmdArgs)
	var errReply reply.ErrorReply
	var deleted int64 = 0
	for _, r := range replies {
		if reply.IsErrorReply(r) {
			errReply = r.(reply.ErrorReply)
			break
		}
		intReply, OK := r.(*reply.IntReply)
		if !OK {
			errReply = reply.MakeErrReply("error")
		}
		deleted += intReply.Code
	}
	if errReply == nil {
		return reply.MakeIntReply(deleted)
	}
	return reply.MakeErrReply("error : " + errReply.Error())
}
