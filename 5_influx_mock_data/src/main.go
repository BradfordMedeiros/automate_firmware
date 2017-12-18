/*


data types: numeric, categorical

if categorical: specify list, and then uniformly populate the data
if number, specify the range, and uniformly populate the data

specify time range
--interval m --ic 4
 */

package main

import (
	"flag"
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

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

type settings struct {
	topic string
	min_value int
	max_value int
	data_type string

	fromYear int
	untilYear int

	fromMonth int
	untilMonth int

	fromDay int
	untilDay int

	interval string
}

func log_data_to_influx(time_to_log_data time.Time, topic string, value string){
	fmt.Println("@: ", time_to_log_data)
	fmt.Println("log: ", topic, " with value: ", value)
}

func get_interval_duration(interval string) time.Duration {
	duration := time.Second;
	if interval == "s" {
		duration = time.Second
		fmt.Println("second interval")
	}else if interval == "m" {
		duration = time.Minute
		fmt.Println("minute interval")
	}else if interval == "h" {
		duration = time.Hour
		fmt.Println("hour interval")
	}else if interval == "d" {
		duration = time.Hour * 24
		fmt.Println("day interval")
	}else {
		fmt.Println("invalid interval")
	}
	return duration
}

func get_random_value (min_value int, max_value int) int {
	return (rand.Intn((max_value - min_value)) + min_value)
}

func run_program_with_settings(settings *settings ) {
	fmt.Println("settings topic is: ", settings)

	fromTime := time.Date(
		settings.fromYear, // year
		time.Month(12), // month
		settings.fromDay, // day
		12, // hour
		0, // min
		0, // s
		0, // ns
		time.UTC,
	)
	untilTime := time.Date(
		settings.untilYear, // year
		time.Month(12), // month
		settings.untilDay, // day
		12, // hour
		0, // min
		0, // s
		0, // ns
		time.UTC,
	)

	fmt.Println("begin time: ", fromTime)
	fmt.Println("end time: ", untilTime)

	duration := get_interval_duration(settings.interval)

	currentTime := fromTime

	count := 0

	for currentTime.Before(untilTime){
		currentTime = currentTime.Add(duration)
		log_data_to_influx(currentTime, "test topic", strconv.Itoa(get_random_value(settings.min_value, settings.max_value)))

		count ++
	}

	fmt.Println("total logged data: ", count)

}


func main() {
	topicFlag := param{is_set: false}
	flag.Var(&topicFlag, "topic", "topic to create mock data")

	typeFlag := param{is_set: false}
	flag.Var(&typeFlag, "type", "<categorical/numeric>")

	intervalFlag := param{is_set: false}
	flag.Var(&intervalFlag, "interval", "<month, day, minute, second>")

	minValue := flag.Int("min", 1, "minimum value to write for numeric data")
	maxValue := flag.Int("max", 100, "maxmum value to write for numeric data")

	today := time.Now()
	fromYear := flag.Int("fromYear", today.Year(), "start year")
	untilYear := flag.Int("untilYear", today.Year(), "end year")
	fromMonth := flag.Int("fromMonth",  1, "start month")
	untilMonth := flag.Int("untilMonth", 1, "end month")
	fromDay := flag.Int("fromDay", today.Day(), "start day")
	untilDay := flag.Int("untilDay", today.Day(), "end day")

	flag.Parse()

	program_settings := settings{
		topic: topicFlag.value,
		data_type: typeFlag.value,
		min_value: *minValue,
		max_value: *maxValue,
		fromYear: *fromYear,
		untilYear: *untilYear,
		fromMonth: *fromMonth,
		untilMonth: *untilMonth,
		fromDay: *fromDay,
		untilDay: *untilDay,
		interval: intervalFlag.value,
	}

	run_program_with_settings(&program_settings)



	/*newDate := fromTime.AddDate(0,0,1)

	oneSecond := time.Duration(time.Second)
	oneMinute := time.Duration(time.Minute)


	fmt.Println("new time: ", newDate)
	fmt.Println("duration: ", oneSecond)
	fmt.Println("minute: ", oneMinute)
	fmt.Println("time diff is: ", newDate.After(fromTime))*/

	/*if topicFlag.is_set {
		fmt.Println("topic is: ", topicFlag.value)
		fmt.Println("min value is: ", *minValue)
		fmt.Println("max value is: ", *maxValue)
	} else {
		fmt.Println("No parameters specified- see usage\n")
		flag.Usage()
	}

	run_program_with_settings(&program_settings)*/


}
