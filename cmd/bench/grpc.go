package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

var logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

var logOpts = []logging.Option{
	logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
}

func InterceptorLogger(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			msg = fmt.Sprintf("DEBUG :%v", msg)
		case logging.LevelInfo:
			msg = fmt.Sprintf("INFO :%v", msg)
		case logging.LevelWarn:
			msg = fmt.Sprintf("WARN :%v", msg)
		case logging.LevelError:
			msg = fmt.Sprintf("ERROR :%v", msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
		l.Println(append([]any{"msg", msg}, fields...))
	})
}
