package main

import (
	"fmt"
)
//defer到最后执行
func defer01()  {
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("1")
}

//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
func defer02()  {
	for i:=0; i<10; i++{
		defer fmt.Printf("%v,", i)
	}
	fmt.Println("done")
}

func main()  {
	defer01()
	defer02()
	arr01()
}
