package conf

var WSConfig WSConf = WSConf{
	MaxConnection: 0,
	PingTimeOut:   180, //second
	WsPortRange:   "38001-38050",
}

type WSConf struct {
	MaxConnection uint64
	PingTimeOut   int
	WsPortRange   string
}
