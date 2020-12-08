package main

import "fmt"

func arr01()  {
	//数组的长度是其类型的一部分，因此数组不能改变大小
	var a [2]string
	a[0] = "h"
	a[1] = "w"
	fmt.Println(a[0], a[1])
	fmt.Println(a)	//输出有中括号

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
}

func main()  {
	arr01()
}
