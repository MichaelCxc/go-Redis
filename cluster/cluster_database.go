package cluster

import (
	pool "github.com/jolestar/go-commons-pool/v2"
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
