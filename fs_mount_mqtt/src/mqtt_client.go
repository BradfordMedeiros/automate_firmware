package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type mqtt_message struct {
	topic string
	message string
}

func createClientOptions(clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", "127.0.0.1:1883"))
	opts.SetClientID("fs_mount_mqtt")
	return opts
}

func connect(clientId string) mqtt.Client {
	opts := createClientOptions(clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {

	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func listen(topic string, mqtt_messages  chan <- mqtt_message) {
	client := connect("sub")
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		mqtt_messages <- mqtt_message{ topic: msg.Topic(), message: string(msg.Payload())}
	})
}


