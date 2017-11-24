package main

import mqtt "github.com/eclipse/paho.mqtt.golang"
import "fmt"
import "os"
import "github.com/nu7hatch/gouuid"
import "errors"
import "encoding/json"
import "io/ioutil"

type subscription struct {
	Uuid      string
	Path      string
	Is_script bool // if not script is file path, how to do enums?
}

type mqtt_manager struct {
	Topic_subscriptions map[string][]subscription
	mqtt_client         mqtt.Client
	on_mqtt_message     func(mqtt_message)
	serialization_file  string
}

func contains_subscription(topicToCheck string, topic_subscriptions *map[string][]subscription) bool {
	for topic, _ := range *topic_subscriptions {
		if topic == topicToCheck {
			return true
		}
	}
	return false
}

func New_mqtt_manager(mqtt_client mqtt.Client, on_mqtt_message func(mqtt_message), serialization_file string) mqtt_manager {
	subscriptions := make(map[string][]subscription)
	manager := mqtt_manager{
		Topic_subscriptions: subscriptions,
		mqtt_client:         mqtt_client,
		on_mqtt_message:     on_mqtt_message,
		serialization_file:  serialization_file,
	}
	deserialized_subscriptions := manager.deserialize_state()
	for topic, subscriptions := range deserialized_subscriptions {
		for _, subscription := range subscriptions {
			if subscription.Is_script {
				manager.add_script_subscription(topic, subscription.Path)
			} else {
				manager.add_file_subscription(topic, subscription.Path)
			}
		}
	}
	return manager
}

func add_subscription(manager *mqtt_manager, topic string, file_path string, is_script bool) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return errors.New("Could not create uuid")

	}

	subscription_to_add := subscription{Uuid: uuid.String(), Path: file_path, Is_script: is_script}

	subscriptions, ok := (*manager).Topic_subscriptions[topic]
	if !contains_subscription(topic, &((*manager).Topic_subscriptions)) {
		(*manager).mqtt_client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
			message := mqtt_message{topic: msg.Topic(), message: string(msg.Payload())}
			(*manager).on_mqtt_message(message)
		})
	}

	if ok {
		(*manager).Topic_subscriptions[topic] = append(subscriptions, subscription_to_add)
	} else {
		(*manager).Topic_subscriptions[topic] = make([]subscription, 0)
		(*manager).Topic_subscriptions[topic] = append((*manager).Topic_subscriptions[topic], subscription_to_add)
	}

	manager.on_subscription_change()
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

	for subscription_topic, subscriptions := range manager.Topic_subscriptions {
		for index, subscription := range subscriptions {
			if subscription.Uuid == uuid {
				topic = subscription_topic
				index_to_remove = index
			}
		}
	}

	if index_to_remove == -1 {
		manager.on_subscription_change()
		return errors.New("UUID does not exist")
	} else {
		manager.Topic_subscriptions[topic] = append(manager.Topic_subscriptions[topic][:index_to_remove], manager.Topic_subscriptions[topic][index_to_remove+1:]...)
		if len(manager.Topic_subscriptions) == 0 {
			delete(manager.Topic_subscriptions, topic)
			manager.mqtt_client.Unsubscribe(topic)
		}

		manager.on_subscription_change()
		return nil
	}
}

func (manager *mqtt_manager) reset() {
	for topic, _ := range manager.Topic_subscriptions {
		manager.mqtt_client.Unsubscribe(topic)
	}
	manager.Topic_subscriptions = make(map[string][]subscription)
	manager.on_subscription_change()
}

func (manager mqtt_manager) list_subscription() string {
	list_value := ""
	for topic, subscriptions := range manager.Topic_subscriptions {
		for _, subscription := range subscriptions {
			list_value = list_value + "\ntopic: " + topic
			list_value = list_value + "    file: " + subscription.Path
			list_value = list_value + "    uuid: " + subscription.Uuid

			if subscription.Is_script {
				list_value = list_value + "    type: " + subscription.Path

			} else {
				list_value = list_value + "    type: " + subscription.Path
			}

		}
	}

	if list_value == "" {
		return "No Subscriptions"
	} else {
		return list_value
	}
}

func (manager mqtt_manager) handle_mqtt_message(topic string, value string) {
	subscriptions, ok := manager.Topic_subscriptions[topic]
	if !ok {
		fmt.Fprintln(os.Stderr, "fs_mount_mqtt error: handle_mqtt_message: got topic ", topic, " when not subscribed to topic")
	} else {
		for _, subscription := range subscriptions {
			if subscription.Is_script {
				execute_file(subscription.Path, topic, value)
			} else {
				write_file(subscription.Path, value)
			}
		}
	}

}

func (manager mqtt_manager) on_subscription_change() {
	manager.serialize_state()
}

func (manager *mqtt_manager) serialize_state() []byte {
	json_string, _ := json.Marshal(manager)
	_ = ioutil.WriteFile(manager.serialization_file, json_string, 0644)
	return json_string
}

func (manager *mqtt_manager) deserialize_state() map[string][]subscription {
	new_manager := mqtt_manager{}
	serialized_data, _ := ioutil.ReadFile(manager.serialization_file)
	json.Unmarshal(serialized_data, &new_manager)
	return new_manager.Topic_subscriptions
}
