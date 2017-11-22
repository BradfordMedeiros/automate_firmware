package main

import "fmt"

func main() {

	mqtt_messages := make(chan mqtt_message)
	tcp_requests := make(chan tcp_request)

	go listen("temperature", mqtt_messages)
	go listen_tcp(tcp_requests)

	mqtt_topic_manager := mqtt_manager{ }

	for {
	  select {
		case mqtt := <-mqtt_messages:
			fmt.Println("received mqtt: ", mqtt.topic)
		case request := <-tcp_requests:
			fmt.Println("received request: ", request.action)

			action_type := request.action.Action

		  	if action_type  ==  "list" {
				request.finish_client(mqtt_topic_manager.list_subscription())
			}else if action_type == "reset" {
				request.finish_client("ok - placeholder")
			}else if action_type == "topic_path" {
				mqtt_topic_manager.add_file_subscription(request.action.Topic, request.action.Path_or_script)
			}else if action_type == "topic_script" {
				mqtt_topic_manager.add_script_subscription(request.action.Topic, request.action.Path_or_script)
			}else if action_type == "delete" {
				mqtt_topic_manager.remove_subscription("placeholder id")
			}else {
				request.finish_client("error - invalid command")
			}


		}

	}
}
