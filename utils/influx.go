package utils

import (
	"time"
	"github.com/influxdata/influxdb/client/v2"
	"log"
)

var (
	clnt client.Client
)

func init() {
	// Influx connection:
	if influxEnabled() {
		clnt = influxDBClient()
		clnt.Query(client.Query{
			Database: Database(),
			Command: "CREATE DATABASE " + Database(),
		})
	}
}

func WriteToInflux(value int, hand string) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  Database(),
		Precision: "us",
	})
	if err != nil {
		log.Fatal(err)
	}

	hands := map[string]string{
		"left": "left",
		"right": "right",
	}
	tags := map[string]string{
		"hand": hands[hand],
	}

	fields := map[string]interface{}{
		"capable": value,
	}

	pt, err := client.NewPoint(
		"hand",
		tags,
		fields,
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)

	if err := clnt.Write(bp); err != nil {
		log.Fatal(err)
	}
}

func influxDBClient() client.Client {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: InfluxAddr(),
	})
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func influxEnabled() bool{
	if InfluxAddr() != "" {
		return true
	} else {
		return false
	}
}