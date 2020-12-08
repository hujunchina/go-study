package main

import (
	"fmt"
	"sync"
)

//通常使用工厂方法将goroutine和通道绑定。
type Receiver struct{
	sync.WaitGroup
	data chan int
}

func NewReceiver() *Receiver {
	r:= &Receiver{
		data:make(chan int),
	}
	r.Add(1)
	go func() {
		defer r.Done()
		for x:=range r.data{
			fmt.Println("recv:", x)
		}
	}()
	return r
}

func main() {

	r := NewReceiver()
	r.data<-1
	r.data<-2
	close(r.data)
	r.Wait()

	var wg sync.WaitGroup
	wg.Add(2)

	c:=make(chan int)

	go func() {	//接收端
		defer wg.Done()

		for{
			var v int
			var ok bool
			select {
			case v,ok = <-c:
				fmt.Printf("a1:%v\n", v)
			case v,ok = <-c:
				fmt.Printf("a2:%v\n", v)
			}
			if !ok {	//避开select阻塞，注意退出循环，以避免CPU空转
				return
			}
		}
	}()

	go func() {  //发送者
		defer wg.Done()
		defer close(c)

		for i:=0; i<10; i++{
			select{
			case c<-i:
			case c<-i*10:
			}
		}
	}()

	wg.Wait()
}