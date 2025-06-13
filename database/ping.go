package database

import (
	"go-Redis/interface/resp"
	"go-Redis/resp/reply"
)

func Ping(db *DB, args [][]byte) resp.Reply {
	return reply.PongReply{}
}

func init() {
	RegisterCommand("ping", Ping, 1)
}
