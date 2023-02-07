package api

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

type APIConn struct {
	rateLimiter *rate.Limiter
}

func Open() *APIConn {
	return &APIConn{
		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 2),
	}
}

func (c *APIConn) Read(context context.Context) (string, error) {
	if err := c.rateLimiter.Wait(context); err != nil {
		return "", err
	}

	return "Data Result", nil
}

func (c *APIConn) Resolve(context context.Context) error {
	if err := c.rateLimiter.Wait(context); err != nil {
		return err
	}

	return nil
}
