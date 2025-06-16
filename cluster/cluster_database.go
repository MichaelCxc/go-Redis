package cluster

import (
	"context"
	pool "github.com/jolestar/go-commons-pool/v2"
	"go-Redis/config"
	database2 "go-Redis/database"
	"go-Redis/interface/database"
	"go-Redis/interface/resp"
	"go-Redis/lib/consistenthash"
)

type ClusterDatabase struct {
	self           string
	nodes          []string
	peerPicker     *consistenthash.NodeMap
	peerconnection map[string]*pool.ObjectPool
	db             database.Database
}

func MakeClusterDatabase() *ClusterDatabase {
	cluster := &ClusterDatabase{
		self:           config.Properties.Self,
		db:             database2.NewStandaloneDatabase(),
		peerPicker:     consistenthash.NewNodeMap(nil),
		peerconnection: make(map[string]*pool.ObjectPool),
	}

	nodes := make([]string, 0, len(config.Properties.Peers)+1)
	cluster.nodes = nodes
	for _, peer := range config.Properties.Peers {
		nodes = append(nodes, peer)
	}
	nodes = append(nodes, config.Properties.Self)
	cluster.peerPicker.AddNode(nodes...)
	ctx := context.Background()
	for _, peer := range config.Properties.Peers {
		pool.NewObjectPoolWithDefaultConfig(ctx, &connectionFactory{
			Peer: peer,
		})
	}
	return cluster
}

func (c *ClusterDatabase) Exec(client resp.Connection, args [][]byte) resp.Reply {
	//TODO implement me
	panic("implement me")
}

func (c *ClusterDatabase) Close() {
	//TODO implement me
	panic("implement me")
}

func (c *ClusterDatabase) AfterClientClose(c resp.Connection) {
	//TODO implement me
	panic("implement me")
}
