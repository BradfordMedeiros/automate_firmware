package main

import mqtt "github.com/eclipse/paho.mqtt.golang"
import "fmt"

type subscription struct {
	uuid      string
	path      string
	is_script bool // if not script is file path, how to do enums?
}

type mqtt_manager struct {
	topic_subscriptions map[string][]subscription
	mqtt_client         mqtt.Client
	on_mqtt_message     func(mqtt_message)
}

func contains_subscription(topicToCheck string, topic_subscriptions *map[string][]subscription) bool {
	for topic, _ := range *topic_subscriptions {
		if topic == topicToCheck {
			return true
		}
	}
	return false
}

func New_mqtt_manager(mqtt_client mqtt.Client, on_mqtt_message func(mqtt_message)) mqtt_manager {
	subscriptions := make(map[string][]subscription)
	return mqtt_manager{topic_subscriptions: subscriptions, mqtt_client: mqtt_client, on_mqtt_message: on_mqtt_message}
}

func add_subscription(manager *mqtt_manager, topic string, file_path string, is_script bool) {
	subscriptions, ok := (*manager).topic_subscriptions[topic]
	subscription_to_add := subscription{uuid: "testuuid", path: file_path, is_script: is_script}

	if !contains_subscription(topic, &((*manager).topic_subscriptions)) {
		(*manager).mqtt_client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
			message := mqtt_message{topic: msg.Topic(), message: string(msg.Payload())}
			(*manager).on_mqtt_message(message)
		})
	}

	if ok {
		(*manager).topic_subscriptions[topic] = append(subscriptions, subscription_to_add)
	} else {
		(*manager).topic_subscriptions[topic] = make([]subscription, 0)
		(*manager).topic_subscriptions[topic] = append((*manager).topic_subscriptions[topic], subscription_to_add)
	}

}

func (manager mqtt_manager) add_script_subscription(topic string, file_path string) {
	add_subscription(&manager, topic, file_path, true)
}

func (manager mqtt_manager) add_file_subscription(topic string, file_path string) {
	add_subscription(&manager, topic, file_path, false)
}

func (manager mqtt_manager) remove_subscription(uuid string) {

}

func (manager *mqtt_manager) reset() {
	for topic, _ := range manager.topic_subscriptions {
		manager.mqtt_client.Unsubscribe(topic)
	}
	fmt.Println("in reset in mqttmanager")
	manager.topic_subscriptions =  make(map[string][]subscription)

}

func (manager mqtt_manager) list_subscription() string {
	list_value := ""
	for topic, subscriptions := range manager.topic_subscriptions {
		for _, subscription := range subscriptions {
			list_value = list_value + "\ntopic: " + topic
			list_value = list_value + "    file: " + subscription.path
			list_value = list_value + "    uuid: " + subscription.uuid

			if subscription.is_script {
				list_value = list_value + "    type: " + subscription.path

			} else {
				list_value = list_value + "    type: " + subscription.path
			}

		}
	}

	if list_value == "" {
		return "No Subscriptions"
	} else {
		return list_value
	}
}
