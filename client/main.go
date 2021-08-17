package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	/**
	直接连接服务  得到一个connection
	writer写数据
	*/

	connection, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Printf("client start err ,exit  %s  \n", err)
	}

	for {
		connection.Write([]byte("hello  this is first  demo for your "))
		return
	}

	buf := make([]byte, 512)
	cnt, err := connection.Read(buf)
	if err != nil {
		fmt.Sprintf(" reader error %s   \n", err)
	}

	fmt.Printf("server call back %s ,cnt=%d  \n", buf, cnt)

	time.Sleep(1 * time.Second)
}
