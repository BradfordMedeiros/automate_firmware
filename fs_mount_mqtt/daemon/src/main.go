package main

import "fmt"

func main() {

	settings := get_command_line_options()
	client := connect_to_mqtt_broker()

	tcp_requests := make(chan tcp_request)
	go listen_tcp(settings.tcp_port,  tcp_requests)

	mqtt_messages := make(chan mqtt_message)
	mqtt_topic_manager := New_mqtt_manager(client, func(message mqtt_message) {
		mqtt_messages <- message
	})

	for {
		select {
		case mqtt := <-mqtt_messages:
			mqtt_topic_manager.handle_mqtt_message(mqtt.topic, mqtt.message)
		case request := <-tcp_requests:
			fmt.Println("received request: ", request.action)

			action_type := request.action.Action

			if action_type == "list" {
				new_subscriptions := mqtt_topic_manager.list_subscription()
				request.finish_client(new_subscriptions)
			} else if action_type == "reset" {
				mqtt_topic_manager.reset()
				request.finish_client("ok")
			} else if action_type == "topic_path" {
				mqtt_topic_manager.add_file_subscription(request.action.Topic, request.action.Path_or_script)
				request.finish_client("ok")
			} else if action_type == "topic_script" {
				mqtt_topic_manager.add_script_subscription(request.action.Topic, request.action.Path_or_script)
				request.finish_client("ok")
			} else if action_type == "delete" {
				err := mqtt_topic_manager.remove_subscription(request.action.Id)
				if err == nil {
					request.finish_client("ok")
				}else{
					request.finish_client(err.Error())
				}
			} else {
				request.finish_client("error - invalid command")
			}
		}

	}
}
