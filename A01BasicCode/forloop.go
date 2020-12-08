package main

import(
	"fmt"
)

//原始形式
//Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的。
func for01(){
	sum:=0
	for i:=0; i<10; i++{
		sum+=i
	}
	fmt.Println(sum)
}

func for02(){
	sum:=1
	for ;sum<100;{
		sum+=sum
	}
	fmt.Println(sum)
}

func for03()  {
	sum:=1
	for sum<1000{
		sum+=sum
	}
	fmt.Println(sum)
}
//range 形式可遍历切片或映射
func for04() {
	var pow = []int{1,2,4,8}
	for i,v:=range pow{
		fmt.Println(i,v)
	}
}

func main()  {
	for01()
	for02()
	for03()
	for04()
}