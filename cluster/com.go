package cluster

import (
	"context"
	"errors"
	"go-Redis/resp/client"
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
