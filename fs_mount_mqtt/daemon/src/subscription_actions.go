package main

import "io/ioutil"
import "os/exec"
import "fmt"


func execute_file(path_to_file string, topic string, message string) {
	fmt.Println("executing script: ", path_to_file)
	cmd := exec.Command("bash", path_to_file, topic, message)
	cmd.Start()
}

func write_file(path_to_file string, message string) {
	_ = ioutil.WriteFile(path_to_file, []byte(message), 0644)
}

