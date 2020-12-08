package main

import (
	"fmt"
	"time"
)

func Subscriber() {
	service := New("tcp://127.0.0.1:1883",
		"mqtt_go_clientA",
		"down/deviceId",
		"up/deviceId")
	service.connect()
	for {
		token := service.subscribe()
		token.Wait()
		if token.Error() != nil{
			fmt.Printf("error: %s", token.Error())
		}
	}

}

func main()  {
	go Subscriber()

	client := New("tcp://127.0.0.1:1883",
		"mqtt_go_clientB",
		"up/deviceId",
		"down/deviceId")

	client.connect()

	for i:=0; i<10; i++ {
		msg := "hujun"
		token := client.publish(msg)
		if token.Error() != nil{
			fmt.Printf("error: %s", token.Error())
		} else {
			fmt.Printf("发送 %s 成功\n", msg)
		}
		time.Sleep(2 * time.Second)
	}
}
