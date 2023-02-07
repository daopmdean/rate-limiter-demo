package apiv2

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

type APIConn struct {
	apiLimit,
	dbLimit RateLimiter
}

func Open() *APIConn {
	return &APIConn{
		apiLimit: MultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 1),
			rate.NewLimiter(Per(5, time.Minute), 5),
		),
		dbLimit: MultiLimiter(
			rate.NewLimiter(rate.Every(time.Second*5), 1),
		),
	}
}

func Per(eventCount int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCount))
}

func (c *APIConn) Read(context context.Context) (string, error) {
	if err := c.apiLimit.Wait(context); err != nil {
		return "", err
	}

	return "Data Result", nil
}

func (c *APIConn) Resolve(context context.Context) error {
	if err := c.dbLimit.Wait(context); err != nil {
		return err
	}

	return nil
}
