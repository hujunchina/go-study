package main
/*
 * 最简单的消息队列，通过设置一个jobChan通道当作队列
 * 入队操作是 【对列<-元素】
 * 出队操作是  处理就出队
 */
import (
	"fmt"
	"time"
)

type Job int

func worker(jobChan <- chan Job)  {
	for job:= range jobChan {
		//顺序执行，缺点：阻塞整个队列
		Process(job)

		//并发执行，没有协程池维护，资源不可控
		//go Process(job)
	}
}

func Process(job Job)  {
	if job==3 {
		time.Sleep(10*time.Second)
	}
	fmt.Printf("job: %v\n", job)
}

func main(){
	//设置一个10个大小的队列
	jobChan := make(chan Job, 2)
	//启动协程
	go worker(jobChan)
	//入队
	for i:=0; i<20; i++{
		jobChan <- Job(i)	//这样就入队了
		fmt.Printf("len: %v\n", len(jobChan))
	}
	time.Sleep(2*time.Second)
}


