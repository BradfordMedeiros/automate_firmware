package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func get_random_value(min_value int, max_value int) int {
	return (rand.Intn((max_value - min_value)) + min_value)
}

func run_numeric(influxHost string, influxDatabase string, topic string, min_value int, max_value int, duration time.Duration, fromTime time.Time, untilTime time.Time) {
	currentTime := fromTime
	count := 0
	for currentTime.Before(untilTime) {
		currentTime = currentTime.Add(duration)
		log_data_to_influx(currentTime, topic, strconv.Itoa(get_random_value(min_value, max_value)))
		count++
	}
	fmt.Println("total logged data: ", count)
}
func run_category(influxHost string, influxDatabase string, topic string, categoryData string, duration time.Duration, fromTime time.Time, untilTime time.Time) {
	fmt.Println("category placeholder")
	currentTime := fromTime
	count := 0
	categories := strings.Split(categoryData, ",")
	fmt.Println(categories)
	for currentTime.Before(untilTime) {
		currentTime = currentTime.Add(duration)
		category := categories[get_random_value(0, len(categories))]
		log_data_to_influx(currentTime, topic, category)
		count++
	}
	fmt.Println("total logged data: ", count)
}
