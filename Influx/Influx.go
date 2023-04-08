package Influx

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type Client struct {
	FluxClient influxdb2.Client
}
type InfluxChan struct {
	Ctx     context.Context
	Data    *chan map[string]interface{}
	ErrChan *chan error
}

func New(address, token string) *Client {
	return &Client{
		FluxClient: influxdb2.NewClient(address, token),
	}
}
func (c *Client) Write(org, bucket string, influx InfluxChan) {
	writer := c.FluxClient.WriteAPIBlocking(org, bucket)
	for {
		select {
		case d := <-*influx.Data:
			fmt.Println(d)
			point := influxdb2.NewPoint("Api-1", nil, d, time.Now())
			err := writer.WritePoint(influx.Ctx, point)
			if err != nil {
				fmt.Println(err)
				*influx.ErrChan <- err
			}
		case <-influx.Ctx.Done():
			fmt.Println("test")
			return
		}
	}

}
