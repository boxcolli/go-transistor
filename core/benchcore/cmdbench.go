package benchcore

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/boxcolli/go-transistor/middleware/base/spoofmiddleware"
	"github.com/boxcolli/go-transistor/tools"
	"github.com/boxcolli/go-transistor/types"
	"google.golang.org/protobuf/proto"
)

func (c *benchCore) cmdBench(ctx context.Context, args []string) (<-chan string, error) {
	fs := flag.NewFlagSet("bench", flag.ContinueOnError)
	var (
		forTime	= fs.Duration("for", time.Second * 10, "watch count duration for")
		rate	= fs.Duration("rate", time.Second * 1, "report rate")
	)
	{
		err := fs.Parse(args)
		if err != nil { return nil, ErrInvalidArgument }
	}

	var (
		in	= make(chan *types.Message, 10)
		mw	= spoofmiddleware.NewSpoofMiddlware(in)
		out	= make(chan string, 10)
	)
	go func() {
		var (
			done		= ctx.Done()
			endTimer	= tools.NewTimer(*forTime)
			rateTimer	= tools.NewTimer(*rate)
			end		= endTimer.End()
			rateEnd	= rateTimer.End()
			count	= tools.NewCounter(0)
			sum		= tools.NewCounter(0)
			rateSec	= rate.Seconds()
		)
		defer endTimer.Stop()
		defer rateTimer.Stop()
		defer close(out)
		defer c.mid.Del(mw)

		out <- "bench, ms/s, b/s"

		c.mid.Add(mw)
		endTimer.Set()
		rateTimer.Set()
		for {
			select {
			case <- done:
				return

			case m := <- in:
				count.AddOne()

				b, err := proto.Marshal(m.Marshal())
				if err != nil { return }
				sum.Add(int64(len(b)))

			case <- end:
				return

			case <- rateEnd:
				rateTimer.Set()
				out <- fmt.Sprintf(
					"bench, %s, %s",
					count.Quo(rateSec).Text('g', 1),
					sum.Quo(rateSec).Text('g', 1),
				)
				count.Reset()
				sum.Reset()
			}
		}
	} ()

	return out, nil
}
