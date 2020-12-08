package main

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	service := New("tcp://127.0.0.1:1883",
		"mqtt_go_clientA",
		"down/deviceId",
		"up/deviceId")

	client := New("tcp://127.0.0.1:1883",
		"mqtt_go_clientB",
		"up/deviceId",
		"down/deviceId")

	service.subscribe()

	msg := "hujun"
	token := client.publish(msg)
	if token.Error() != nil{
		fmt.Printf("error: %s\b", token.Error())
	} else {
		fmt.Printf("发送 %s 成功\n", msg)
	}
}
