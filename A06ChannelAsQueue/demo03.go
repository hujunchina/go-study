package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumSenders = 1000

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// senders
	for i := 0; i < NumSenders; i++ {
		go func() {
			for {
				select {
				case <- stopCh:		//只要触发通道就被select选择并执行分支
					return
				case dataCh <- rand.Intn(Max):	//入队
				}
			}
		}()
	}

	// the receiver
	go func() {
		for value := range dataCh {	//不停出队操作
			if value == Max-1 {
				fmt.Println("send stop signal to senders.")
				close(stopCh)	//触发通道的一种
				return
			}

			fmt.Println(value)
		}
	}()

	select {
	case <- time.After(time.Hour):
	}
}
