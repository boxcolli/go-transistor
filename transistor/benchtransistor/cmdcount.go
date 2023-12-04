package benchtransistor

import (
	"flag"
	"fmt"
	"time"
	"context"

	"github.com/boxcolli/go-transistor/middleware/base/countmiddleware"
	"github.com/boxcolli/go-transistor/tools"
)

func (c *benchTransistor) cmdCount(ctx context.Context, args []string) (<-chan string, error) {
	if len(args) == 0 { return nil, ErrNotFound }
	switch args[0] {
	case CmdCountMod:	return c.cmdCountMod(ctx, args[1:])
	case CmdCountRate: 	return c.cmdCountRate(ctx, args[1:])
	default: 			return nil, ErrNotFound
	}
}

func (c *benchTransistor) cmdCountMod(ctx context.Context, args []string) (<-chan string, error) {
	// Parse args
	fmt.Println("cmdCountMod(): args:", args)
	fs := flag.NewFlagSet("count", flag.ContinueOnError)
	var (
		forTime	= fs.Duration("for", time.Second * 10, "total count duration")
		mod		= fs.Int64("mod", 1000000, "make an intermediate report if count % mod == 0")
	)
	{
		err := fs.Parse(args)
		if err != nil { return nil, ErrInvalidArgument }
	}

	var (
		in		= make(chan bool, 10)
		mw		= countmiddleware.NewCountMiddlware(in)
		out		= make(chan string, 10)
	)

	// Run middleware
	go func() {
		var (
			done 	= ctx.Done()
			timer	= tools.NewTimer(*forTime)
			end 	= timer.End()
			count 	= tools.NewCounter(*mod)
		)
		defer timer.Stop()
		defer close(out)
		defer c.mid.Del(mw)

		out <- "mod, count, elapsed"

		c.mid.Add(mw)
		timer.Set()
		start := time.Now()
		for {
			select {
			case <- done:
				return

			case <- end:
				elapsedSec := time.Since(start).Seconds()
				out <- fmt.Sprintf("mod, %s, %.1f", count.String(), elapsedSec)
				out <- fmt.Sprintf("end, ms/s, %s\n", count.Quo(elapsedSec).Text('f', -1))
				return

			case <- in:
				count.AddOne()
				if count.IsModZero() {
					out <- fmt.Sprintf("mod, %s, %.1f", count.String(), time.Since(start).Seconds())
				}
			}
		}
	} ()

	return out, nil
}

func (c *benchTransistor) cmdCountRate(ctx context.Context, args []string) (<-chan string, error) {
	// Parse args
	fs := flag.NewFlagSet("rate", flag.ContinueOnError)
	var (
		forTime	= fs.Duration("for", time.Second * 10, "watch count duration for")
		rate	= fs.Duration("rate", time.Second * 1, "report rate")
	)
	{
		err := fs.Parse(args)
		if err != nil {
			return nil, ErrInvalidArgument
		}
	}

	var (
		in	= make(chan bool, 10)
		mw	= countmiddleware.NewCountMiddlware(in)
		out	= make(chan string, 10)
	)
	go func() {
		var (
			done		= ctx.Done()
			endTimer	= tools.NewTimer(*forTime)
			rateTimer	= tools.NewTimer(*rate)
			end		= endTimer.End()
			rate	= rateTimer.End()
			count	= tools.NewCounter(0)
		)
		defer endTimer.Stop()
		defer rateTimer.Stop()
		defer close(out)
		defer c.mid.Del(mw)

		out <- "rate, count, ms/s"

		c.mid.Add(mw)
		endTimer.Set()
		rateTimer.Set()
		start := time.Now()
		for {
			select {
			case <- done:
				return
				
			case <- end:
				elapsedSec := time.Since(start).Seconds()
				out <- fmt.Sprintf("rate, %s, %.1f\n", count.String(), elapsedSec)
				out <- fmt.Sprintf("end, ms/s, %s\n", count.Quo(elapsedSec).Text('f', 1))
				return

			case <- rate:
				rateTimer.Set()
				out <- fmt.Sprintf("rate, %s, %.1f\n", count.String(), time.Since(start).Seconds())

			case <- in:
				count.AddOne()
			}
		}
	} ()

	return out, nil
}
