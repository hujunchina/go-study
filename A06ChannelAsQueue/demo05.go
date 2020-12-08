package main
//等待worker处理所有队列的job
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker3(n int, jobChan <-chan Job, wg *sync.WaitGroup)  {
	for job:= range jobChan {
		Process3(n, job, wg)
	}
}

func Process3(n int, job Job, wg *sync.WaitGroup)  {
	defer wg.Done()	//执行完调用，唤醒wg.Wait()

	size := rangTime()
	time.Sleep(size*time.Millisecond)
	fmt.Printf("worker %2d process job %d, time %dms\n", n, job, size)
}

func rangTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(400)+100)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	ch := make(chan struct{})	//定义一个临时的通道
	//异步执行
	go func() {
		wg.Wait()
		close(ch)	//等待wg被唤醒并关闭通道
	}()

	select {
	case <- ch:		//如果及时关闭了通道，就会触发通道信息，返回正常
		return true
	case <- time.After(timeout):	//如果超时就返回失败
		return false
	}
}

//使用等待队列
func workPool3(workerNum int, jobChan chan Job, wg *sync.WaitGroup)  {
	for i:=0; i<workerNum; i++{
		go func(i int) {
			worker3(i, jobChan, wg)
		}(i)
	}
}

func main()  {
	wg := &sync.WaitGroup{}
	jobChan := make(chan Job, 10)

	workPool3(5, jobChan, wg)

	for i:=0; i<20; i++{
		wg.Add(1)
		jobChan <- Job(i)
	}

	t := time.Now()

	wg.Wait()

	ok := WaitTimeout(wg, 300*time.Millisecond)
	if !ok {
		fmt.Printf("\n warning, process job timeout \n")
	}

	fmt.Printf("\n handle queue time %v\n", time.Now().Sub(t))
}