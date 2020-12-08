package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

//使用context或channel停止worker

func workPool6(workerNum int, jobChan chan Job, ctx context.Context) {
	for i:=0; i<workerNum; i++ {
		go func(i int) {
			worker6(i, jobChan, ctx)
		}(i)
	}
}

func worker6(n int, jobChan <-chan Job, ctx context.Context) {
	for {
		select {
		case job := <-jobChan:
			Process6(n, job)
		case <-ctx.Done():
			fmt.Printf("cancel worker %d\n", n)
			return
		}
	}
}

func Process6(n int, job Job)  {
	size := randTime()
	time.Sleep(size * time.Millisecond)
	fmt.Printf("worker %2d process job %2d, time %dms\n", n, job, size)
}

//随机时间100～500
func randTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(400)+100)
}

func main(){
	jobChan := make(chan Job, 10)
	// 使用context控制worker是否停止，适合多级函数传递和控制，并且有超时取消
	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	workPool6(5, jobChan, ctx)
	for i:=0; i<20; i++{
		jobChan <- Job(i)
	}
	cancel()	//？
	time.Sleep(5 * time.Second)
}