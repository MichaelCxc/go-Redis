package database

import (
	"go-Redis/interface/database"
	"go-Redis/interface/resp"
	"go-Redis/lib/utils"
	"go-Redis/resp/reply"
)

// GET k1
func execGet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	entity, exists := db.GetEntity(key)
	if !exists {
		return reply.MakeNullBulkReply()
	}

	bytes := entity.Data.([]byte)
	return reply.MakeBulkReply(bytes)
}

// SET k1 v
func execSet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := args[1]
	entity := &database.DataEntity{
		Data: value,
	}
	db.PutEntity(key, entity)
	db.addAof(utils.ToCmdLine2("set", args...))
	return reply.MakeOkReply()
}

// SETNX k1 v1
func execSetnx(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := args[1]
	entity := &database.DataEntity{
		Data: value,
	}

	res := db.PutIfAbsent(key, entity)
	db.addAof(utils.ToCmdLine2("setnx", args...))
	return reply.MakeIntReply(int64(res))
}

// GETSET k1 v1
func execGetSet(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	value := args[1]
	entity, exists := db.GetEntity(key)
	db.PutEntity(key, &database.DataEntity{Data: value})
	if !exists {
		return reply.MakeNullBulkReply()
	}
	db.addAof(utils.ToCmdLine2("getset", args...))
	return reply.MakeBulkReply(entity.Data.([]byte))
}

// STRLEN
func execStrLen(db *DB, args [][]byte) resp.Reply {
	key := string(args[0])
	entity, exists := db.GetEntity(key)
	if !exists {
		return reply.MakeNullBulkReply()
	}

	bytes := entity.Data.([]byte)
	return reply.MakeIntReply(int64(len(bytes)))
}

func init() {
	RegisterCommand("Get", execGet, 2)
	RegisterCommand("Set", execSet, 3)
	RegisterCommand("Setnx", execSetnx, 3)
	RegisterCommand("GetSet", execGetSet, 3)
	RegisterCommand("StrLen", execStrLen, 2)
}
