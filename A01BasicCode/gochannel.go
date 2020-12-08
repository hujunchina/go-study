package main

import(
	"fmt"
)

func sum01(s []int, c chan int)  {
	sum := 0
	for _, v := range s{
		sum+=v
	}
	c <- sum
}
//信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
//默认情况下，发送和接收操作在另一端准备好之前都会阻塞。
//这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步 。
//
func main() {
	s := []int{7,2,8,-9,4}
	c := make(chan int)
	go sum01(s[:len(s)/2], c)
	go sum01(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x,y)
}