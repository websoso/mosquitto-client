package mc

import (
	"encoding/json"
	"fmt"
	paho "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	pahoClient paho.Client
	converter  *Converter
}

func NewClient(option *Option, fallback Handler) *Client {
	pahoOption := pahoOpt(option)
	pahoClient := paho.NewClient(pahoOption)
	token := pahoClient.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	c := &Client{pahoClient: pahoClient, converter: NewConverter(option.capacity, fallback)}
	pahoClient.Subscribe(option.clientName, 0, func(client paho.Client, message paho.Message) {
		payload := string(message.Payload())
		var m *Message
		json.Unmarshal([]byte(payload), &m)
		c.converter.decode(m)
	})
	return c
}

func (c *Client) Send(topic, message string) {
	wrapper := c.converter.encode(message)
	fmt.Println(len(wrapper.messageList))
	for _, m := range wrapper.messageList {
		jsonMessage, _ := json.Marshal(m)
		token := c.pahoClient.Publish(topic, byte(0), false, string(jsonMessage))
		if token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}
}

func pahoOpt(opt *Option) *paho.ClientOptions {
	option := paho.NewClientOptions()
	option.AddBroker(opt.broker)
	option.SetClientID(opt.clientId)
	option.SetUsername(opt.username)
	option.SetPassword(opt.password)
	option.SetCleanSession(true)
	return option
}
