package main

import(
	"fmt"
	"strings"
)

func slice01()  {
	primes:=[4]int{2,3,5,7}

	var s []int=primes[1:3]
	x:=s[1:]
	fmt.Println(s)
	fmt.Println(x)
}

func slice02()  {
	names:=[4]string{
		"A","B","C", "D",
	}
	fmt.Println(names)
	a:=names[0:2]
	b:=names[1:3]
	fmt.Println(a,b)

	b[0] = "xxx"
	fmt.Println(a,b)
	fmt.Println(names)
}

func slice03()  {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func slice04() {
	//切片可以用内建函数 make 来创建，这也是你创建动态数组的方式
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println(b)
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	fmt.Println(b)
	b = b[1:]      // len(b)=4, cap(b)=4
	a := make([]int, 5)
	printSlice(a)
}

//切片的切片
func slice05() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func slice06() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


func main()  {
	slice01()
	slice02()
	slice03()
	slice04()
	slice05()
}
