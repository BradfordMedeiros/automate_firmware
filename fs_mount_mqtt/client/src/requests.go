package main

import "fmt"
import "encoding/json"

type basic_command struct {
	Action string `json:"action"`
}

type topic_command struct {
	Action string `json:"action"`
	Topic string `json:"topic"`
	Path_or_script string `json:"path_or_script"`
}

type delete_command struct {
	Action string `json:"action"`
	Id string `json:"id"`
}

func list() {
	command := basic_command { Action: "list" }
	json_string, _ := json.Marshal(command)

	fmt.Println("list placeholder: ", string(json_string))
}

func reset() {
	command := basic_command { Action: "reset" }
	json_string, _ := json.Marshal(command)

	fmt.Println("reset placeholder: ", string(json_string))
}

func add_topic_with_path(topic string, file_path string) {
	command := topic_command { Action: "topic_path"}
	json_string, _ := json.Marshal(command)

	fmt.Println("add topic  w/ path placeholder", string(json_string))
}

func add_topic_with_script(topic string, script_path string) {
	command := topic_command { Action: "topic_script"}
	json_string, _ := json.Marshal(command)

	fmt.Println("add topic w/ script placeholder: ", string(json_string))
}

func delete_subscription(uuid string) {
	command :=  delete_command { Action: "delete", Id: uuid }
	json_string, _ := json.Marshal(command)

	fmt.Println("delete topic placeholder: ", string(json_string))
}