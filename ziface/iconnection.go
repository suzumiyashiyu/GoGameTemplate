package ziface

import "net"

type IConnection interface{
	//启动链接
	StartConnection()

	StopConnection()

	GetTcpConnection() *net.TCPConn

	GetId()uint32

	Send(chars []byte)

	RemoteAddr()  *net.Addr
}


type HandleFunc func( *net.TCPConn,[]byte,int )error