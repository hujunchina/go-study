package main

import "fmt"

/**
 * 已知f1函数存在，f2是新写的，如何让f1调用f2？
 * 这就需要闭包做函数适配，类似适配器模式
 * 设计一个可以作为f1参数的函数f3，其参数是f2即可
 * 闭包特点就是一个内部函数使用了外部变量
 */
func f1(ff func(int), a int) {
	fmt.Println("f1")
	ff(a)
}

func f2(x,y int) int {
	fmt.Println("f2")
	return x+y
}

func f3(ff func(int, int) int, x int, y int) (func(int), int) {
	//直接返回一个匿名函数和ff，但是没有用ff(x,y)()直接调用
	return func(int){
		fmt.Println(x+1)
	}, ff(x,y)
}

func main(){
	f1(f3(f2,10,10))
}