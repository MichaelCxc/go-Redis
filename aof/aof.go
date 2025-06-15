package aof

import (
	"go-Redis/config"
	"go-Redis/interface/database"
	"go-Redis/lib/logger"
	"go-Redis/lib/utils"
	"go-Redis/resp/reply"
	"os"
	"strconv"
)

type CmdLine = [][]byte

const aofBufferSize = 1 << 16

type payload struct {
	cmdLine CmdLine
	dbIndex int
}

type AofHandler struct {
	database    database.Database
	aofChan     chan *payload
	aofFile     *os.File
	aofFilename string
	currentDB   int
}

//New aof Handler

func NewAofHandler(db database.Database) (*AofHandler, error) {
	handler := &AofHandler{}
	handler.aofFilename = config.Properties.AppendFilename
	handler.database = db
	//load aof
	handler.LoadAof()
	aoFile, err := os.OpenFile(handler.aofFilename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	handler.aofFile = aoFile
	//channel
	handler.aofChan = make(chan *payload, aofBufferSize)
	go func() {
		handler.handleAof()
	}()
	return handler, nil
}

// Add payload(set k v) -> aofChan
func (handler *AofHandler) AddAof(dbIndex int, cmdLine CmdLine) {

	if config.Properties.AppendOnly && handler.aofChan != nil {
		handler.aofChan <- &payload{
			cmdLine: cmdLine,
			dbIndex: dbIndex,
		}
	}

}

// handle AOF payload(set k v) <- aofChan(flush)
func (handler *AofHandler) handleAof() {
	handler.currentDB = 0
	for p := range handler.aofChan {
		if p.dbIndex != handler.currentDB {
			data := reply.MakeMultiBulkReply(utils.ToCmdLine("SELECT", strconv.Itoa(p.dbIndex))).ToBytes()
			_, err := handler.aofFile.Write(data)

			if err != nil {
				logger.Warn(err)
				continue
			}
			handler.currentDB = p.dbIndex
		}

		data := reply.MakeMultiBulkReply(p.cmdLine).ToBytes()
		_, err := handler.aofFile.Write(data)
		if err != nil {
			logger.Warn(err)
		}
	}
}

// Load AOF
func (handler *AofHandler) LoadAof() {

}
