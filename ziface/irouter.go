package ziface

/***
路由抽象接口
路由里的数据都是IRequest
*/

type IRouter interface {
	//在处理conn 业务之前的钩子方法 hook
	PreHandle(request IRequest)

	Handle(request IRequest)

	PostHandle(request IRequest)
}
