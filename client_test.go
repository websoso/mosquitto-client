package mc

import (
	"fmt"
	"testing"
	"time"
)

const (
	Broker   = "tcp://39.107.254.231:2883"
	ClientId = "SERVER_GO"
	Username = "go_client"
	Password = "000000"
	Topic    = "WEBRTC/SERVER_GO_01"
	Capacity = 10
)

func TestNewClient(t *testing.T) {
	client := NewClient(opt(), fallback)
	client.Send(Topic, "ABCDEFG梁召峰啊ABCDEFG梁召峰啊ABCDEFG梁召峰啊ABCDEFG梁召峰啊ABCDEFG梁召峰啊")
	client.Send(Topic, "这是一个消息吗")
	time.Sleep(time.Minute)
}

func opt() *Option {
	option := NewOption()
	option.SetBroker(Broker)
	option.SetClientId(ClientId)
	option.SetUsername(Username)
	option.SetPassword(Password)
	option.SetClientName(Topic)
	return option
}

func fallback(payload string) {
	fmt.Println("Received message", payload)
}
