package database

import "go-Redis/interface/resp"

type CmdLine = [][]byte

type DataEntity struct {
	Data interface{}
}

type Database interface {
	Exec(client resp.Connection, args [][]byte) resp.Reply
	Close()
	AfterClientClose(c resp.Connection)
}
