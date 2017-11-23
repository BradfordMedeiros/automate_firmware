package main

import mqtt "github.com/eclipse/paho.mqtt.golang"
import "fmt"
import "github.com/nu7hatch/gouuid"
import "errors"

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

func add_subscription(manager *mqtt_manager, topic string, file_path string, is_script bool) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return errors.New("Could not create uuid")

	}

	subscription_to_add := subscription{uuid: uuid.String(), path: file_path, is_script: is_script}

	subscriptions, ok := (*manager).topic_subscriptions[topic]
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

	return nil
}

func (manager mqtt_manager) add_script_subscription(topic string, file_path string) error {
	return add_subscription(&manager, topic, file_path, true)
}

func (manager mqtt_manager) add_file_subscription(topic string, file_path string) error {
	return add_subscription(&manager, topic, file_path, false)
}

func (manager mqtt_manager) remove_subscription(uuid string) error {

	topic := ""
	index_to_remove := -1

	for subscription_topic, subscriptions := range manager.topic_subscriptions {
		for index, subscription := range subscriptions {
			if subscription.uuid == uuid{
				topic  = subscription_topic
				index_to_remove = index
			}
		}
	}

	if index_to_remove == -1 {
		return errors.New("UUID does not exist")
	}else{
		manager.topic_subscriptions[topic] = append(manager.topic_subscriptions[topic][:index_to_remove], manager.topic_subscriptions[topic][index_to_remove+1:]...)
		if  len(manager.topic_subscriptions) == 0{
			delete(manager.topic_subscriptions, topic)
			manager.mqtt_client.Unsubscribe(topic)
		}

		return nil
	}
}

func (manager *mqtt_manager) reset() {
	for topic, _ := range manager.topic_subscriptions {
		manager.mqtt_client.Unsubscribe(topic)
	}
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

func (mqtt_manager) handle_mqtt_message(topic string, value string){
	fmt.Println("mqtt message: ", topic, " value: ", value)
}
