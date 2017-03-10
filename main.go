package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"time"
)

var (
	HelloTopic = "/espresso/device/hello"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	o := mqtt.NewClientOptions()
	o.SetClientID("fake-device-generator")
	o.AddBroker("tcp://localhost:1883")

	c := mqtt.NewClient(o)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		//fmt.Println("Failed to connect", token.Error())
		panic(token.Error())
	}

	fmt.Println("Connected to broker!")

	for {
		if token := c.Publish(HelloTopic, 1, false, fmt.Sprintf("%d", rand.Int())); token.Wait() && token.Error() != nil {
			panic(token)
		}
		time.Sleep(5 * time.Second)
	}

}
