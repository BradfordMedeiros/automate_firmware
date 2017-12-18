package main

import "fmt"

type settings struct {
	topic      string
	min_value  int
	max_value  int
	data_type  string
	fromYear   int
	untilYear  int
	fromMonth  int
	untilMonth int
	fromDay    int
	untilDay   int
	interval   string
}

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

func print_error() {
	fmt.Println("error invalid parameters")
}
