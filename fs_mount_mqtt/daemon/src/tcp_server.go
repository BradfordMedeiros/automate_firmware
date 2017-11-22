
package main

import (
	"encoding/json"
	"net"
	"fmt"
	"os"
	"bytes"
)



type command struct {
	Action string `json:"action"`
	Topic string `json:"topic"`
	Path_or_script string `json:"path_or_script"`
}


const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

type tcp_request struct {
	action command
	finish_client func(string)
}



// Handles incoming requests.
func handleRequest(conn net.Conn, tcp_message  chan <- tcp_request) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}


	parsed_json := command{ }
	trimmed_bytes := bytes.Trim(buf, "\x00")
	json.Unmarshal(trimmed_bytes, &parsed_json)

	finish_client :=  func (message string) {
		// Send a response back to person contacting us.
		conn.Write([]byte(message))
		// Close the connection when you're done with it.
		conn.Close()
	}

	tcp_message <- tcp_request{ action: parsed_json, finish_client: finish_client}



}

func listen_tcp(tcp_message  chan <- tcp_request) {
	CONN_HOST := "localhost"
	CONN_PORT := "3333"
	CONN_TYPE := "tcp"

	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		handleRequest(conn, tcp_message)
	}
}
