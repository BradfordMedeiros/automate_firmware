package main

import "fmt"
import "time"

func log_data_to_influx(time_to_log_data time.Time, topic string, value string) {
	fmt.Println("@: ", time_to_log_data, "topic: ", topic, " value: ", value)
}
