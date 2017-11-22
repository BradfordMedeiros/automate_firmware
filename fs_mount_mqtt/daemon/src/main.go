package main

import "fmt"

func main() {

	client := listen()
	mqtt_messages := make(chan mqtt_message)
	tcp_requests := make(chan tcp_request)

	go listen_tcp(tcp_requests)

	mqtt_topic_manager := New_mqtt_manager(client, func(message mqtt_message) {
		mqtt_messages <- message
	})

	for {
		select {
		case mqtt := <-mqtt_messages:
			fmt.Println("received mqtt: ", mqtt.topic)
			fmt.Println("value is : ", mqtt.message)
			mqtt_topic_manager.handle_mqtt_message(mqtt.topic, mqtt.message)
		case request := <-tcp_requests:
			fmt.Println("received request: ", request.action)

			action_type := request.action.Action

			if action_type == "list" {
				size := len(mqtt_topic_manager.topic_subscriptions)
				fmt.Println("before: size of list: ", size)
				new_subscriptions := mqtt_topic_manager.list_subscription()
				fmt.Println("new subscription: ", new_subscriptions)
				size2 := len(mqtt_topic_manager.topic_subscriptions)
				fmt.Println("after: size2 of list: ", size2)
				request.finish_client(new_subscriptions)
			} else if action_type == "reset" {
				//mqtt_topic_manager.topic_subscriptions
				size := len(mqtt_topic_manager.topic_subscriptions)
				fmt.Println("before: size of list: ", size)

				mqtt_topic_manager.reset()
				fmt.Println("reset called")
				size2 := len(mqtt_topic_manager.topic_subscriptions)
				fmt.Println("after: size of list: ", size2)
				request.finish_client("ok")
			} else if action_type == "topic_path" {
				mqtt_topic_manager.add_file_subscription(request.action.Topic, request.action.Path_or_script)
				request.finish_client("ok")
			} else if action_type == "topic_script" {
				mqtt_topic_manager.add_script_subscription(request.action.Topic, request.action.Path_or_script)
				request.finish_client("ok")
			} else if action_type == "delete" {
				mqtt_topic_manager.remove_subscription("placeholder id")
			} else {
				request.finish_client("error - invalid command")
			}
		}

	}
}
