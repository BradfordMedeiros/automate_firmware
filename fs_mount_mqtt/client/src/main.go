package main

import (
	"flag"
	"fmt"
)

type param struct {
	is_set bool
	value string
}

func (param *param) String() string {
	return "hello world"
}

func (param *param) Set(s string) error {
	param.is_set = true
	param.value = s
	return nil
}

func print_error() {
	fmt.Println("error invalid parameters")
}

func main() {
	topicFlag := param{ is_set: false }
	flag.Var(&topicFlag, "topic", "topic to subscribe to")

	pathFlag := param{ is_set: false }
	flag.Var(&pathFlag, "path", "file to write when a topic is published")

	scriptFlag := param{ is_set: false }
	flag.Var(&scriptFlag, "script", "script to call when a topic is published")

	deleteFlag := param{ is_set: false }
	flag.Var(&deleteFlag, "delete", "uuid of subscription to delete")

	resetFlag := flag.Bool("reset", false, "reset all subscriptions")
	listFlag := flag.Bool("list", false, "list all subscriptions")

	flag.Parse()

	if *listFlag {
		list()
	}else if *resetFlag {
		reset()
	}else if deleteFlag.is_set {
		delete_subscription(deleteFlag.value)
	}else if topicFlag.is_set {
		if scriptFlag.is_set && pathFlag.is_set {
			print_error()
		}else if scriptFlag.is_set {
			add_topic_with_script(topicFlag.value, scriptFlag.value)
		}else if pathFlag.is_set {
			add_topic_with_path(topicFlag.value, pathFlag.value)
		}else{
			print_error()
		}
	}else {
		fmt.Println("No parameters specified- see usage\n")
		flag.Usage()
	}

}
