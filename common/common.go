package common

import (
	"log"
	"os"
	"time"
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
func logF(prefix string, format string, args ...interface{}) {
	log.Printf("[%s] [%s] "+format, append([]interface{}{time.Now().Format(timeFormat), prefix}, args...)...)
}
