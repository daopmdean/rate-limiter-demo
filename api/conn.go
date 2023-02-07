package api

import "context"

type APIConn struct{}

func Open() *APIConn {
	return &APIConn{}
}

func (c *APIConn) Read(context context.Context) (string, error) {
	return "Data Result", nil
}

func (c *APIConn) Resolve(context context.Context) error {
	return nil
}
