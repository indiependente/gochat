package common

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

const timeFormat = "03:04:05 PM"

func Debugf(format string, args ...interface{}) {
	logF("DEBUG", format, args)
}
func Errorf(format string, args ...interface{}) {
	log.SetOutput(os.Stderr)
	logF("ERROR", format, args)
	log.SetOutput(os.Stdout)
}
func Warningf(format string, args ...interface{}) {
	logF("WARNING", format, args)
}
func ClientLogf(format string, args ...interface{}) {
	logF("CLIENT", format, args)
}
func ServerLogf(format string, args ...interface{}) {
	logF("SERVER", format, args)
}
func MessageLogf(format string, args ...interface{}) {
	log.Printf(format, args...)
}
func logF(prefix string, format string, args ...interface{}) {
	log.Printf("[%s] [%s] "+format, append([]interface{}{time.Now().Format(timeFormat), prefix}, args...)...)
}

func TsToTime(ts *timestamp.Timestamp) time.Time {
	t, err := ptypes.Timestamp(ts)
	if err != nil {
		return time.Now()
	}
	return t.In(time.Local)
}

func SignalContext(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func(c context.CancelFunc) {
		Debugf("%s", "listening for shutdown signal")
		<-sigs
		Debugf("%s", "shutdown signal received")
		signal.Stop(sigs)
		close(sigs)
		cancel()
	}(cancel)

	return ctx
}
