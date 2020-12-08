package main

import (
	"fmt"
)

func point01()  {
	i,j:=42,122
	p:=&i // 指向 i
	fmt.Println(*p)
	fmt.Println(p) //地址
	*p = 21
	fmt.Println(*p)
	fmt.Println(p)

	p=&j
	*p = *p/2
	fmt.Println(j) //改变
}

func main()  {
	point01()
}