package benchtransistor

import (
	"context"
	"errors"
	"fmt"
)

const (
	CmdPing = "ping"
	CmdCount = "count"
	CmdCountMod = "mod"
	CmdCountRate = "rate"
	CmdBench = "bench"
	CmdCommand = "command"
)

var (
	ErrNotFound = errors.New("the command doesn't exist")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrUnavailable = errors.New("unavailable")
)

func (c *benchTransistor) command(ctx context.Context, args []string) (<-chan string, error) {
	fmt.Println("Command(): args:", args)
	if len(args) == 0 { return nil, ErrNotFound }
	switch args[0] {
	case CmdPing:	return c.cmdPing(ctx, args[1:])
	case CmdCount:	return c.cmdCount(ctx, args[1:])
	case CmdBench:	return c.cmdBench(ctx, args[1:])
	case CmdCommand:return c.cmdCommand(ctx, args[1:])
	default:		return nil, ErrNotFound
	}
}

func (c *benchTransistor) cmdPing(ctx context.Context, args []string) (<-chan string, error) {
	ch := make(chan string, 1)
	defer close(ch)

	ch <- "pong"
	return ch, nil
}

