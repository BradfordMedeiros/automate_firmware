package main

import (
	"encoding/json"
	"log"
	"net"
	"fmt"
	"strconv"
	"strings"
)

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

	send_message("127.0.0.1", 3333, string(json_string))
}

func reset() {
	command := basic_command { Action: "reset" }
	json_string, _ := json.Marshal(command)

	send_message("127.0.0.1", 3333, string(json_string))
}

func add_topic_with_path(topic string, file_path string) {
	command := topic_command { Action: "topic_path", Topic: topic, Path_or_script: file_path}
	json_string, _ := json.Marshal(command)

	send_message("127.0.0.1", 3333, string(json_string))
}

func add_topic_with_script(topic string, script_path string) {
	command := topic_command { Action: "topic_script", Topic: topic, Path_or_script: script_path}
	json_string, _ := json.Marshal(command)

	send_message("127.0.0.1", 3333, string(json_string))
}

func delete_subscription(uuid string) {
	command :=  delete_command { Action: "delete", Id: uuid }
	json_string, _ := json.Marshal(command)

	send_message("127.0.0.1", 3333, string(json_string))
}

func send_message(ip string, port int, payload string) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("message is: ", payload)
	conn.Write([]byte(payload + "\n"))
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])
}
