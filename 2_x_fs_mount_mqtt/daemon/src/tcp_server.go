package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strconv"
)

type command struct {
	Action         string `json:"action"`
	Topic          string `json:"topic"`
	Path_or_script string `json:"path_or_script"`
	Id             string `json:"id"`
}

type tcp_request struct {
	action        command
	finish_client func(string)
}

// Handles incoming requests.
func handleRequest(conn net.Conn, tcp_message chan<- tcp_request) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	parsed_json := command{}
	trimmed_bytes := bytes.Trim(buf, "\x00")
	json.Unmarshal(trimmed_bytes, &parsed_json)

	finish_client := func(message string) {
		// Send a response back to person contacting us.
		conn.Write([]byte(message))
		// Close the connection when you're done with it.
		conn.Close()
	}

	tcp_message <- tcp_request{action: parsed_json, finish_client: finish_client}

}

func listen_tcp(port int, tcp_message chan<- tcp_request) {
	CONN_HOST := "localhost"
	CONN_PORT := strconv.Itoa(port)
	CONN_TYPE := "tcp"

	server, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		handleRequest(conn, tcp_message)
	}
}
