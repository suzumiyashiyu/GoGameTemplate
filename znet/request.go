package znet

import "GoGameServer/ziface"

type Request struct {
	//已经和客户端建立好的链接
	conn ziface.IConnection
	//客户端请求的数据
	data []byte
}

func (request *Request) GetConnection() *ziface.IConnection {
	return &request.conn
}

//得到请求的消息数据
func (request *Request) GetData() []byte {
	return request.data
}

func newRequest(conn ziface.IConnection, data []byte) Request {
	request := Request{
		conn: conn,
		data: data,
	}
	return request
}
