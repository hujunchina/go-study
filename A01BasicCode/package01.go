package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var s string="123"

//常量不能用 := 语法声明
const PI = 314/100

//
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

//类型转换,都需要显式转换
func tranType(){
	var x,y int=3,4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x,y,z)
}


func main(){
	fmt.Println("产生一个随机数", rand.Intn(10) )
	fmt.Printf("小数输出 %g \n", math.Sqrt(7))
	fmt.Println("以大写字母开头", math.Pi)
	fmt.Printf("字符串用q %q\n", s)

	fmt.Printf("type %T, value %v \n", ToBe, ToBe)
	fmt.Printf("type %T, value %v \n", MaxInt, MaxInt)
	fmt.Printf("type %T, value %v \n", z, z)

	tranType()
}


/*
Go 的基本类型
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
*/