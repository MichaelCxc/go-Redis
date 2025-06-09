package reply

type PongReply struct {
}

var pongbytes = []byte("+PONG\r\n")

func (r PongReply) ToBytes() []byte {
	return pongbytes
}

func MakePngReply() *PongReply {
	return &PongReply{}
}

type OkReply struct{}

var okbytes = []byte("+OK\r\n")

func (o OkReply) ToBytes() []byte {
	return okbytes
}

var theOkReply = new(OkReply)

func MakeOkReply() *OkReply {
	return theOkReply
}

type NullBulkReply struct{}

var nullBulkBytes = []byte("$-1\r\n")

func (n NullBulkReply) ToBytes() []byte {
	return nullBulkBytes
}

func MakeNullBulkReply() *NullBulkReply {
	return &NullBulkReply{}
}

var emptyMultiBulkBytes = []byte("$0\r\n")

type EmptyMultiBulkReply struct{}

func (e EmptyMultiBulkReply) ToBytes() []byte {
	return emptyMultiBulkBytes
}

func MakeEmptyMultiBulkReply() *EmptyMultiBulkReply {
	return &EmptyMultiBulkReply{}
}

type NoReply struct{}

var noReplyBytes = []byte("")

func (n NoReply) ToBytes() []byte {
	return noReplyBytes
}

func MakeNoReply() *NoReply {
	return &NoReply{}
}
