package database

import (
	"go-Redis/datastruct/dict"
	"go-Redis/interface/resp"
)

type DB struct {
	index int
	data dict.Dict
}

type ExecFunc func(db *DB, args [][]byute) resp.Reply{

}

type CmdLine = [][]byte

func makeDB() *DB{
	db := &DB{
		data: dict.MakeSyncDict(),
	}
	return db
}

