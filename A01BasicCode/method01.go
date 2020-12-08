package main

import (
	"fmt"
)

type hu struct{
	x,y float64
}

//方法就是一类带特殊的 接收者 参数的函数。
//只修改副本
func (h hu) getX() float64{
	h.x = 0
	return h.x
}

//指针接收者
//指针接收者的方法可以修改接收者指向的值
func (h *hu) scale(f float64)  {
	h.x = h.x*f
	h.y = h.y*f
}

func main() {
	v:=hu{1,2}
	fmt.Println(v.getX())
	fmt.Println(v)

	v.scale(1.2)
	fmt.Println(v)
}
