package basiccore

import (
	"context"
	"errors"
)

const (
	CmdPing = "ping"
)

var (
	ErrNotFound = errors.New("the command doesn't exist")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrUnavailable = errors.New("unavailable")
)

func (c *basicCore) command(ctx context.Context, args []string) (<-chan string, error) {
	switch args[0] {
	case CmdPing:	return c.cmdPing(ctx, args[1:])
	default:		return nil, ErrNotFound
	}
}

func (c *basicCore) cmdPing(ctx context.Context, args []string) (<-chan string, error) {
	ch := make(chan string, 1)
	defer close(ch)

	ch <- "pong"
	return ch, nil
}
