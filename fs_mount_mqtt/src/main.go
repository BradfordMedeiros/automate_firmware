package main

import "time"

func main() {
	hello()

	go listen("temperature")

	client := connect("pub")
	timer := time.NewTicker(1 * time.Second)
	for range timer.C {
		client.Publish(topic, 0, false, "hello")
	}
}
