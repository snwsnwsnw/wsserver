package server

import "net"

type IClient[USERDATA any, CONN net.Conn, MT MsgType] interface {
	GetId() ClientID
	GetUserData() USERDATA
	GetConn() CONN
	IsClosed() bool
	Close() error
	WriteMessage(MT, []byte) error
}
