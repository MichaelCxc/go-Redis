package handler

import (
	"context"
	"errors"
	databaseface "go-Redis/interface/database"
	"go-Redis/lib/logger"
	"go-Redis/lib/sync/atomic"
	"go-Redis/resp/connection"
	"go-Redis/resp/parser"
	"go-Redis/resp/reply"
	"io"
	"net"
	"strings"
	"sync"
)

type RespHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
	db         databaseface.Database
}

func (r *RespHandler) closeClient(client *connection.Connection) {
	_ = client.Close()
	r.db.AfterClientClose(client)
	r.activeConn.Delete(client)
}

func (r *RespHandler) Handle(ctx context.Context, conn net.Conn) {
	if r.closing.Get() {
		_ = conn.Close()
	}
	client := connection.NewConn(conn)
	r.activeConn.Store(client, struct{}{})
	ch := parser.ParseStream(conn)

	for payload := range ch {
		//1,error from channel
		if payload.Err != nil {
			if payload.Err == io.EOF || errors.Is(payload.Err, io.ErrUnexpectedEOF) || strings.Contains(payload.Err.Error(), "use of closed network connection") {
				r.closeClient(client)
				logger.Info("Connection closed: " + client.RemoteAddr().String())
				return
			}
		}
		//2.protocol error
		errReply := reply.MakeErrReply(payload.Err.Error())
		err := client.Write(errReply.ToBytes())
		if err != nil {
			r.closeClient(client)
			logger.Error("Connection closed: " + client.RemoteAddr().String())
			return
		}
		continue
	}
}

func (r *RespHandler) Close() error {
	logger.Info("Handler shutting down.")
	r.closing.Set(true)
	r.activeConn.Range(
		func(key interface{}, value interface{}) bool {
			client := key.(*connection.Connection)
			_ = client.Close()
			return true
		},
	)
	r.db.Close()
	return nil
}
