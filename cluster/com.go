package cluster

import (
	"context"
	"errors"
	"go-Redis/interface/resp"
	"go-Redis/lib/utils"
	"go-Redis/resp/client"
	"go-Redis/resp/reply"
	"strconv"
)

func (cluster *ClusterDatabase) getPeerClient(peer string) (*client.Client, error) {
	pool, ok := cluster.peerConnection[peer]
	if !ok {
		return nil, errors.New("connection not exist")
	}
	object, err := pool.BorrowObject(context.Background())
	if err != nil {
		return nil, err
	}
	c, ok := object.(*client.Client)
	if !ok {
		return nil, errors.New("Wrong type")
	}
	return c, err
}

func (cluster *ClusterDatabase) returnPeerClient(peer string, c *client.Client) error {
	pool, ok := cluster.peerConnection[peer]
	if !ok {
		return errors.New("connection not exist")
	}
	return pool.ReturnObject(context.Background(), c)
}

func (cluster *ClusterDatabase) relay(peer string, c resp.Connection, args [][]byte) resp.Reply {
	if peer == cluster.self {
		return cluster.db.Exec(c, args)
	}
	peerClient, err := cluster.getPeerClient(peer)
	if err != nil {
		return reply.MakeErrReply(err.Error())
	}
	defer func() {
		_ = cluster.returnPeerClient(peer, peerClient)
	}()
	peerClient.Send(utils.ToCmdLine("SELECT", strconv.Itoa(c.GetDBIndex())))
	// other nodes have no idea which db was selected
	return peerClient.Send(args)
}

func (cluster *ClusterDatabase) broadcast(c resp.Connection, args [][]byte) map[string]resp.Reply {
	res := make(map[string]resp.Reply)
	for _, node := range cluster.nodes {
		rep := cluster.relay(node, c, args)
		res[node] = rep
	}
	return res
}
