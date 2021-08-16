package main

import "GoGameServer/znet"

func main(){
	server:=znet.NewServer("gameServer")
	server.Run()


}
