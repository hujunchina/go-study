package main

import(
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i:=0; i<5; i++{
		time.Sleep(100* time.Millisecond)
		fmt.Println(s)
	}
}

var set = make(map[int]bool, 0)

//应该打印一次的，但是对set访问冲突，没有加锁，导致多才输出
func printOnce(num int) {
	if _, exist := set[num]; !exist {
		fmt.Println(num)
	}
	set[num] = true
}

var m sync.Mutex
// 使用互斥量控制访问
func printOnce2(num int) {
	m.Lock()
	if _, exist := set[num]; !exist{
		fmt.Println(num)
	}
	set[num] = true
	m.Unlock()
}


//Go 程（goroutine）是由 Go 运行时管理的轻量级线程
func main() {
	go say("world")
	say("hello")

	for i:=0; i<10; i++ {
		go printOnce(100)
	}
	time.Sleep(time.Second)

	for i:=0; i<10; i++{
		go printOnce2(100)
	}
	time.Sleep(time.Second)
}
