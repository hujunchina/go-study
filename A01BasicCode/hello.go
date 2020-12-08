package main

import "fmt"
// 加法
func add(x, y int) int{
	//简洁赋值语句 := 可在类型明确的地方代替 var 声明
	sum := x+y
	return sum
}

func add2(x int, y int) int{
	return x+y
}

// 任意数量的返回值，用小括号
func swap(x, y string) (string, string) {
	return y, x
}

//没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
func split(sum int) (x, y int) {
	x = sum*4/9
	y = sum-x
	return
}

//var 语句可以出现在包或函数级别
var c, python, java bool
//变量会从初始值中获得类型
var ii, jj int = 1,2

//函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
var sum=ii+jj

func main() {
	fmt.Printf("Hello, world\n")
	fmt.Printf("%v \n", add(1,2))

	fmt.Println("a", "b")
	a, b := swap("a", "b")
	fmt.Println(a, b)

	fmt.Println(split(7))

}