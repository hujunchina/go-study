package main

import (
	"fmt"
	"time"
)

//线程池
func workPool7(workNum int, jobChan chan Job, ch chan struct{}){
	for i:= 0; i<workNum; i++{
		go func(i int) {
			worker7(i, jobChan, ch)
		}(i)
	}
}

func worker7(n int, jobChan <-chan Job, ch chan struct{}){
	for{
		select {
		case job:=<-jobChan :
			Process7(n, job)
		case <-ch :
			fmt.Printf("cancel worker %d\n", n)
			return
		}
	}
}

func Process7(n int, job Job)  {
	size := randTime()
	time.Sleep(size * time.Millisecond)
	fmt.Printf("worker %v processing %v job\n", n, job)
}

func main(){
	jobChan := make(chan Job, 10)
	ch := make(chan struct{})
	workPool7(5, jobChan, ch)
	for i:=0; i<20; i++{
		jobChan <- Job(i)	//入队
	}
	close(ch)
	time.Sleep(5 * time.Second)
}