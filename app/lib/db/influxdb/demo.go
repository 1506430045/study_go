package influxdb

import (
	"context"
	"fmt"
	influx "github.com/influxdata/influxdb-client-go"
	"time"
)

func Test() {
	client := influx.NewClient("http://localhost:9999", "my-token")

	writeApi := client.WriteApiBlocking("my-org", "my-bucket")

	p := influx.NewPoint("stat",
		map[string]string{},
		map[string]interface{}{},
		time.Now())

	_ = writeApi.WritePoint(context.Background(), p)

	p = influx.NewPointWithMeasurement("stat").AddTag("unit", "temperature").AddField("avg", 23.2).AddField("max", 23).SetTime(time.Now())
	_ = writeApi.WritePoint(context.Background(), p)

	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	_ = writeApi.WriteRecord(context.Background(), line)

	queryApi := client.QueryApi("my-org")
}
