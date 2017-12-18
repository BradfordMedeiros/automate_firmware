/*


data types: numeric, categorical

if categorical: specify list, and then uniformly populate the data
if number, specify the range, and uniformly populate the data

specify time range
 */

package main

import (
	"flag"
	"fmt"
	"time"
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
	min_value float64
	max_value float64
	data_type string

	fromYear int
	untilYear int

	fromMonth int
	untilMonth int

	fromDay int
	untilDay int
}

func run_program_with_settings(sett *settings ){
  fmt.Println("settings topic is: ", sett)
}

func main() {
	topicFlag := param{is_set: false}
	flag.Var(&topicFlag, "topic", "topic to create mock data")

	typeFlag := param{is_set: false}
	flag.Var(&typeFlag, "type", "<categorical/numeric>")

	minValue := flag.Float64("min", 1, "minimum value to write for numeric data")
	maxValue := flag.Float64("max", 100, "maxmum value to write for numeric data")

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
	}

	fromTime := time.Date(
		program_settings.fromYear,  // year
		time.Month(12),		    // month
		program_settings.fromDay,   // day
		0,			    // hour
		0,			    // min
		0,			    // s
		0,			    // ns
		time.UTC,
	)
	untilTime := time.Date(
		program_settings.fromYear,  // year
		time.Month(12),		    // month
		program_settings.fromDay,   // day
		12,			    // hour
		0,			    // min
		0,			    // s
		0,			    // ns
		time.UTC,
	)

	fmt.Println("begin time: ", fromTime)
	fmt.Println("end time: ", untilTime)

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
