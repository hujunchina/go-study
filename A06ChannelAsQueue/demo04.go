package main

/*
 * 带线程池的消息队列
 * 通过对worker多线程，间接对process多线程
 */

import (
	"fmt"
	"time"
)

//<-chan 声明一个只能用于接收的通道
func worker2(i int, jobChan <-chan Job) {
	for job:= range jobChan {
		Process2(i, job)
	}
}

func Process2(i int, job Job)  {
	if job==3 {
		time.Sleep(time.Second)
	}
	fmt.Printf("worker %2d process job %d\n", i, job)
}

// chan 声明一个双向通道
func workPool(workerNum int, jobChan chan Job)  {
	for i:=0; i<workerNum; i++{
		go func(i int) {
			worker2(i, jobChan)
		}(i)
	}
}

func main()  {

	jobChan := make(chan Job, 10)
	//启动多个worker池并发处理队列的job，
	//多个worker去抢队列job来处，
	//只有空闲的worker才能从队列中获取job
	workPool(5, jobChan)

	for i:=0; i<20; i++{
		jobChan <- Job(i)
	}

	time.Sleep(2*time.Second)
}