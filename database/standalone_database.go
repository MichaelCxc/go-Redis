package database

import (
	"go-Redis/aof"
	"go-Redis/config"
	"go-Redis/interface/resp"
	"go-Redis/lib/logger"
	"go-Redis/resp/reply"
	"strconv"
	"strings"
)

type Database struct {
	dbSet      []*DB
	aofHandler *aof.AofHandler
}

func NewDatabase() *Database {
	database := &Database{}
	if config.Properties.Databases == 0 {
		config.Properties.Databases = 16
	}
	database.dbSet = make([]*DB, config.Properties.Databases)
	for i := range database.dbSet {
		db := makeDB()
		db.index = i
		database.dbSet[i] = db
	}

	if config.Properties.AppendOnly {
		aofHandler, err := aof.NewAofHandler(database)
		if err != nil {
			panic(err)
		}
		database.aofHandler = aofHandler
		for _, db := range database.dbSet {
			sdb := db
			sdb.addAof = func(line CmdLine) {
				database.aofHandler.AddAof(sdb.index, line)
			}
		}
	}
	return database
}

func (db *Database) Exec(client resp.Connection, args [][]byte) resp.Reply {
	defer func() {
		err := recover()
		if err != nil {
			logger.Error(err)
		}
	}()
	cmdName := strings.ToLower(string(args[0]))
	if cmdName == "select" {
		if len(args) != 2 {
			return reply.MakeArgNumErrReply("select")
		}
		return execSelect(client, db, args[1:])
	}

	dbIndex := client.GetDBIndex()
	database := db.dbSet[dbIndex]
	return database.Exec(client, args)
}

func (db *Database) Close() {
}

func (db *Database) AfterClientClose(c resp.Connection) {
}

// select 1
func execSelect(c resp.Connection, database *Database, args [][]byte) resp.Reply {
	dbIndex, err := strconv.Atoi(string(args[0]))
	if err != nil {
		return reply.MakeErrReply("Err invalid DB index")
	}

	if dbIndex < 0 || dbIndex >= len(database.dbSet) {
		return reply.MakeErrReply("Err DB index is out of range")
	}

	c.SelectDB(dbIndex)
	return reply.MakeOkReply()
}
