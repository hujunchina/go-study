package main
// 既是发布者也是订阅者

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"time"
)

const BROKER_URL string = "tcp://localhost:1883"
const CLIENT_ID string = "112562F00000FEEC"
const UP_TOPIC string = "112562F00000FEEC/up"
const DOWN_TOPIC string = "112562F00000FEEC/cmddown"

type MqttService struct{
	mqttClient mqtt.Client
	broker string
	clientId string
	upTopic string
	downTopic string
}

var f mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	//消息主题
	fmt.Printf("TOPIC: %s\n", message.Topic())
	//消息内容
	fmt.Printf("MSG: %s\n", message.Payload())
}

//构造方法
func New(broker string, clientId string, upTopic string, downTopic string) *MqttService{
	//options := mqtt.ClientOptions{
	//	CleanSession:   true,
	//	Username:       "hujun",
	//	Password:       "hujun",
	//	ConnectTimeout: 10,
	//	KeepAlive:      20,
	//}
	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientId)
	opts.SetKeepAlive(20*time.Second)
	opts.SetPingTimeout(1*time.Second)
	opts.SetDefaultPublishHandler(f)
	return &MqttService{
		broker: broker,
		clientId: clientId,
		upTopic: upTopic,
		downTopic: downTopic,
		mqttClient: mqtt.NewClient(opts),
	}
}

//先连接
func (service *MqttService) connect()  {
	token := service.mqttClient.Connect()
	token.Wait()
	if token.Error() != nil {
		fmt.Printf("error %s\n", token.Error())
	}else{
		fmt.Printf("连接成功\n")
	}
}

//订阅主题
func (service *MqttService) subscribe() mqtt.Token{
	token := service.mqttClient.Subscribe(service.downTopic, 0, f)
	return token
}

//发布消息
func (service *MqttService) publish(msg string) mqtt.Token{
	token := service.mqttClient.Publish(service.upTopic, 0, false, msg)
	token.Wait()
	return token
}

//取消订阅
func (service *MqttService) unSubScribe() mqtt.Token{
	token := service.mqttClient.Unsubscribe(service.downTopic)
	token.Wait()
	return token
}

//关闭连接
func (service *MqttService) disconnect() {
	service.mqttClient.Disconnect(250)
}
