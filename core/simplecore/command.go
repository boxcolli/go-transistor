package simplecore

const (
	ping = "ping"
)

func (c *simpleCore) command(args []string) <-chan string {
	switch args[0] {
	case ping:	return c.ping(args)
	default:	return nil
	}
}

func (c *simpleCore) ping(args []string) <-chan string {
	ch := make(chan string, 1)

	ch <- "pong"
	defer close(ch)
	return ch
}
