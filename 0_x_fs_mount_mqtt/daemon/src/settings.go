package main

import "flag"
import "strconv"
import "fmt"

type param struct {
	is_set bool
	value  string
}

func (param *param) String() string {
	return param.value
}

func (param *param) Set(s string) error {
	param.is_set = true
	param.value = s
	return nil
}

type settings struct {
	tcp_port   int
	broker_url string
}

func get_command_line_options() settings {
	brokerPort := param{is_set: false}
	flag.Var(&brokerPort, "broker", "port of the broker")
	port := param{is_set: false}

	flag.Var(&port, "port", "port to host the daemon on (default = 9002)")
	flag.Parse()

	tcp_port := 3333
	if port.is_set {
		portConverted, err := strconv.Atoi(port.value)
		tcp_port = portConverted
		if err != nil {
			fmt.Println("error")
		}
	}

	return settings{tcp_port: tcp_port}

}
