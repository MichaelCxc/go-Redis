package tcp

import (
	"context"
	"go-Redis/lib/sync/atomic"
	"go-Redis/lib/sync/wait"
	"net"
	"sync"
	"time"
)

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

func (e *EchoClient) Close() error {
	e.Waiting.WaitWithTimeout(10 * time.Second)
	_ = e.Conn.Close()
	return nil
}

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

func (handler *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if handler.closing.Get() {
		_ = conn.Close()
	}

	client := &EchoClient{
		Conn: conn,
	}

	handler.activeConn.Store(client, struct{}{})

}

func (handler *EchoHandler) Close() error {
	//TODO implement me
	panic("implement me")
}
