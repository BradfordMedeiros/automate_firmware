package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"time"
)

type mqtt_message struct {
	topic   string
	message string
}

func createClientOptions(clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:1883")
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

func connect_to_mqtt_broker() mqtt.Client {
	client := connect("fs_mount_mqtt")
	return client
}
