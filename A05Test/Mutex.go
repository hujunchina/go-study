package main

import (
	"fmt"
	"sync"
	"time"
)

//接口有明确的关键字
type myInterface interface {
	test()
	string() string
}

//声明一个类
type data struct{}
//实现
func (d *data) test(){
	fmt.Println("test")
}
//实现
func (d *data) string() string{
	return "data class"
}

func main()  {
	//接口实例化
	var d data
	var t myInterface = &d
	t.test()

	var lock sync.RWMutex
	m := make(map[string]string)

	go func() {
		for {
			lock.Lock()
			m["a"] = "apple"
			lock.Unlock()
			time.Sleep(1*time.Second)
		}
	}()

	go func() {
		for {
			lock.RLock()
			fmt.Println(m["a"])
			lock.RUnlock()
			time.Sleep(1*time.Second)
		}
	}()

	select{}
}
