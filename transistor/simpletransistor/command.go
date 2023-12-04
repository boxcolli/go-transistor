package simpletransistor

import (
	"context"
	"errors"
)

const (
	CmdPing = "ping"
)

var (
	ErrNotFound = errors.New("the command doesn't exist")
)

func (c *simpleTransistor) command(ctx context.Context, args []string) (<-chan string, error) {
	switch args[0] {
	case CmdPing:		return c.cmdPing(ctx, args)
	default:			return nil, ErrNotFound
	}
}

func (c *simpleTransistor) cmdPing(ctx context.Context, args []string) (<-chan string, error) {
	ch := make(chan string, 1)
	defer close(ch)

	ch <- "pong"
	return ch, nil
}
