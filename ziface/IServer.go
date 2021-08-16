package ziface
//服务器接口


type IServer interface{
	//启动服务器
	Start()
	//stop
	Stop()
	//run
	Run()


}