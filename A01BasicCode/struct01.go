package main

import (
	"fmt"
)

type vertex struct{
	x int
	y int
}

var (
	v1 = vertex{1,2}
	v2 = vertex{x:1}
	v3 = vertex{}
	p = &vertex{1,2}
)

func struct01()  {
	v:=vertex{2,3}
	fmt.Println(v) //输出有大括号
	v.x = 1
	fmt.Println(v)

	//如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。
	//不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以
	p:=&v
	p.x = 9
	fmt.Println(v)
}

func main()  {
	struct01()

}
