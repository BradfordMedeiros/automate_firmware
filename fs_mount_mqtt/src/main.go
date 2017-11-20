package main

import "fmt"

func main() {

	mqtt_messages := make(chan mqtt_message)

	go listen("temperature", mqtt_messages)

	//mess:= mqtt_message { topic: "something", message: "somethingelse"}

	//fmt.Println("value is ", mess)

	for {
		x := <- mqtt_messages
		fmt.Println("topic: ", x)
	}
	/*for {
		select {
		case b = <-freeList:
		// Got one; nothing more to do.
		default:
		// None free, so allocate a new one.
			b = new(Buffer)
		}
		load(b)              // Read next message from the net.
		serverChan <- b      // Send to server.
	}*/

}
