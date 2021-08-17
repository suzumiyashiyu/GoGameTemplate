package znet

import "GoGameServer/ziface"

/***
	默认实现
这里之所以BaseRouter的方法都为空
是因为由Router不希望由PreHandle PostHandle这俩个业务员
所以router 全部继承
 */
type BaseRouter struct{



}

func (router BaseRouter ) PreHandle(request ziface.IRequest){

}

func (router BaseRouter )  Handle(request ziface.IRequest){

}

func (router BaseRouter )  PostHandle(request ziface.IRequest){

}