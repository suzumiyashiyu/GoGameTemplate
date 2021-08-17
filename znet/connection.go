package znet

import (
	"GoGameServer/ziface"
	"fmt"
	"net"
)

type Connection struct {
	ConnId    uint32
	Conn      *net.TCPConn
	isClosed  bool
	ExitChan  chan bool
	router    ziface.IRouter
}

func (c *Connection) StartReader() {
	fmt.Printf("connection Id:%d will start Reader \n", c.ConnId)
	defer c.StopConnection()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Printf(" connectionId :%d   reader error \n", c.ConnId)
			continue
		}
		//得到当前conn数据的Request请求数据
		var req = newRequest(c, buf[cnt:])
		//从路由中，找到注册绑定的conn对应的router调用
		go func(request *Request) {
			c.router.PreHandle(request)
			c.router.Handle(request)
			c.router.PostHandle(request)
		}(&req)
	}
}

func (conn *Connection) StartConnection() {
	fmt.Printf("connection Id:%d will start \n", conn.ConnId)
	go conn.StartReader()
}

func (conn *Connection) StopConnection() {
	fmt.Printf("connection Id:%d will stop  \n", conn.ConnId)
	conn.Conn.Close()
	conn.isClosed = true
}

func (conn *Connection) GetTcpConnection() *net.TCPConn {
	return conn.Conn
}

func (conn *Connection) GetId() uint32 {

	return conn.ConnId
}

func (conn *Connection) Send(chars []byte) {

}

func (conn *Connection) RemoteAddr() *net.Addr {
	remote := conn.Conn.RemoteAddr()
	return &remote
}

func newConnection(conn *net.TCPConn, connId uint32 ,router ziface.IRouter) *Connection {

	s := Connection{
		ConnId:    connId,
		Conn:      conn,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
		router: router,

	}

	return &s
}
