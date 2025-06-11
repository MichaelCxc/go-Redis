package reply

type UnknownErrReply struct{}

var unknownErrBytes = []byte("-Err unknown\r\n")

func (u UnknownErrReply) Error() string {
	return "Err unknown"
}

func (u UnknownErrReply) ToBytes() []byte {
	return unknownErrBytes
}

type ArgNumErrReply struct {
	Cmd string
}

//var argNumErrBytes = []byte("-Arg num err\r\n")

func (r *ArgNumErrReply) Error() string {
	return "-ERR wrong number of arguments for '" + r.Cmd + "' command\r\n"
}

func (r *ArgNumErrReply) ToBytes() []byte {
	return []byte("-ERR wrong number of arguments for '" + r.Cmd + "' command\r\n")
}

func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{Cmd: cmd}
}

type SyntaxErrReply struct{}

func (s *SyntaxErrReply) Error() string {
	//TODO implement me
	return "Err syntax error"
}

func (s *SyntaxErrReply) ToBytes() []byte {
	return syntaxErrBytes
}

var syntaxErrBytes = []byte("-Err syntax error\r\n")
var theSyntaxErrReply = &SyntaxErrReply{}

func MakeSyntaxErrReply() *SyntaxErrReply {
	return theSyntaxErrReply
}

type WrongTypeErrReply struct{}

func (w WrongTypeErrReply) Error() string {
	return "WRONGTYPE Operation against a key holding the wrong kind of value"
}

func (w WrongTypeErrReply) ToBytes() []byte {
	return wrongTypeErrBytes
}

var wrongTypeErrBytes = []byte("-WRONGTYPE Operation against a key holding the wrong kind of value\\r\\n")

//ProtocolErr

type ProtocolErrReply struct {
	Msg string
}

func (r *ProtocolErrReply) Error() string {
	return "ERR Protocol error: '" + r.Msg
}

func (r *ProtocolErrReply) ToBytes() []byte {
	//TODO implement me
	return []byte("-ERR Protocol error: '" + r.Msg + "'\r\n")
}
