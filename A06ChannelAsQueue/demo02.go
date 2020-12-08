package main

import "time"

var done = make(chan bool)
var msg string

func aGoroutine() {
	msg = "hello, world"
	println(msg)
	time.Sleep(time.Second*3)
	done <- true
}

func main() {
	go aGoroutine()
	<-done	//阻塞操作
	println(msg)
}

