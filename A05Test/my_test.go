package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

//16进制转10进制
func TestArr(t *testing.T) {
	arr := [3]string{"ea","ea", "1b"}
	res := 0
	if arr[0][1] > '9' {
		res += int(arr[0][1] - 'a')+10
	} else {
		res += int(arr[0][1] - '0')
	}
	if arr[0][0] > '9' {
		res += int(arr[0][0] - 'a')*16+10*16
	} else {
		res += int(arr[0][0] - '0')*16
	}
	fmt.Println(res)
}


func TestMutex(t *testing.T) {
	var lock sync.RWMutex
	m := map[string]string{}

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
}

//func TestNewReceiver(t *testing.T) {
//	r := NewReceiver()
//	r.data<-1
//	r.data<-2
//	close(r.data)
//	r.Wait()
//}

func TestTime(t *testing.T)  {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano()/1000000)
	fmt.Printf("%d%d%d\n", time.Now().Year(), time.Now().Month(), time.Now().Day())
	fmt.Println(strings.TrimLeft("Tac-12913293821093", "Tac-"))

	t1 := time.Now()
	time.Sleep(10*time.Second)
	fmt.Println(time.Since(t1))
	fmt.Println(10 * time.Second)
	if time.Since(t1) > 11*time.Second {
		fmt.Println("OK")
	}

}

func TestMap(t *testing.T) {
	m := make(map[string]string)
	m["a"] = "1"
	m["b"] = "2"
	for k, v := range m {
		fmt.Printf("%v, %v\n", k, v)
	}

	data := "\"dps\":{\"8\":0},\"cid\":\"1125620000002EC5\"},\"protocol\":5,\"s\":561153,\"t\":1604134468}"
	idx := strings.Index(data, "dps")
	fmt.Println(idx)
	fmt.Println(string(data[idx+7]))

	var light int
	var tmp int
	tmp = 0
	for{
		light = tmp * 100/4096
		fmt.Println(light)
		if tmp >= 4000 {
			tmp = 0
		}
		tmp += 500
	}

	fmt.Println(time.Now())
	time.Sleep(time.Second*time.Duration(20))
	fmt.Println(time.Now())

}


func TestChannel(t *testing.T) {
	var msgChan chan int
	var recvChan chan int
	groupCmd(100, recvChan)
	groupCmd(200, recvChan)
	for m := range msgChan {
		go groupCmd(m, recvChan)
		select {
		case ret := <-recvChan:
			fmt.Println("ret: ", ret)
		case <-time.After(2):
			fmt.Println("timeout")
		}
	}
}
func groupCmd(m int, recvChan chan int){
	fmt.Println("m: ",m)
	recvChan <- m+1
}

func TestLong2String(t *testing.T){
	var sn int64
	sn = 1329718369783648268
	str := strconv.FormatInt(sn, 10)
	fmt.Println(str)
}

type UpdateCmdDP struct {
	Location string `json:"location,omitempty"`
	Type     string `json:"type,omitempty"`
	IsCancel int    `json:"isCancel,omitempty"`
	Open     int    `json:"open,omitempty"`
	Delay    int    `json:"delay,omitempty"`
	SN       int64  `json:"sn,omitempty"`
	Gateway  bool   `json:"gateway,omitempty"`
}

func TestJsonEmpty(t *testing.T){
	var cmd UpdateCmdDP
	msg := "{\"type\":\"ode\",\"open\": 1}"
	_ = json.Unmarshal([]byte(msg), &cmd)
	if strings.Contains(msg, "type") {
		fmt.Printf("%v", cmd.Type)
	}
	if strings.Contains(msg, "open") {
		fmt.Printf("%v", cmd.Open)
	}
	//fmt.Printf("%v", cmd.Open)
	fmt.Printf("%v", cmd)
}

func TestString(t *testing.T) {
	str := "KS30012012020016_7"
	if strings.Contains(str, "_") {
		cid := strings.Split(str, "_")
		fmt.Printf(cid[0])
		fmt.Println()
	}
}