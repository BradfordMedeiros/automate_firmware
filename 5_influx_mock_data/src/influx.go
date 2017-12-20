package main

import "fmt"
import "time"
import "strconv"
import  "github.com/influxdata/influxdb/client/v2"


func log_data_to_influx(time_to_log_data time.Time, topic string, value string) {
	//con, err := client.NewClient(client.Config{URL: *host})
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	// Create a new point batch
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "automate",
		Precision: "us",
	})

	// Create a point and add to batch
	value_int, err := strconv.Atoi(value)
	tags := map[string]string{"t_value": value}

	if err != nil {
		fmt.Println("error")
		fields := map[string]interface{}{"value": value}
		pt, err := client.NewPoint(topic, tags, fields, time_to_log_data)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		bp.AddPoint(pt)
	}else{
		fields := map[string]interface{}{"value": value_int}
		pt, err := client.NewPoint(topic, tags, fields, time_to_log_data)
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		bp.AddPoint(pt)
	}

	fmt.Println("time is: ", time_to_log_data)
	c.Write(bp)
}
