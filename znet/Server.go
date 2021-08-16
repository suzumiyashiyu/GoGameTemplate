package znet

import (
	"GoGameServer/ziface"
	"fmt"
	"net"
)

//IServer的实现
type Server struct {
	//名称
	Name string
	//版本
	IPVersion string
	//ip
	IP string
	//端口
	Port int

}

func (s *Server) Start() {
	fmt.Printf("starting  server   IP:%s port:%d  isStarting   \n", s.IP, s.Port)
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Start Server Error:", err)
			return
		}
		listen, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listening  server error:", err)
		}
		fmt.Println("Start Zins server succ,", s.Name, "succ listenning ...")

		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Connection errpr", err)
				continue
			}
			go func() {
				for {
					var result = make([]byte, 512)
					cnt, err := conn.Read(result)
					if err != nil {
						fmt.Println("Recived Message Rrror:", err)
						continue
					}
					if _, err := conn.Write(result[:cnt]); err != nil {
						if err != nil {
							fmt.Println("Write Mssage Errpr:", err)
						}

					}
				}
			}()
		}




}

func (s *Server) Stop() {

}

func (s *Server) Run() {
	 s.Start()
	//TODO  启动之后的一些额外工作
	//阻塞

}

/**
初始化 server模块的方法
*/

func NewServer(name string) ziface.IServer {
	s := &Server{
		name,
		"tcp4",
		"0.0.0.0",
		8999,
	}

	return s
}
