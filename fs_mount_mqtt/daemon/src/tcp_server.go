
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
	action string
	finish_client func(string)
}



// Handles incoming requests.
func handleRequest(conn net.Conn) command {
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


	fmt.Println("command is:  ", parsed_json)

	fmt.Println("action is: ", parsed_json.Action)
	fmt.Println("topic is :", parsed_json.Topic)


	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()

	return parsed_json
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
	// Close the listener when the application closes.
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
		handleRequest(conn)
	}
}
