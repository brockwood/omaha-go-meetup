package script

import (
	"fmt"
	"time"
)

var (
	Command     = `print-logs`
	Description = `prints out log entries`
)

func Cmd(times []time.Time, logLines []string) {
	for entry, logLine := range logLines {
		fmt.Printf("%s: %s\n", times[entry].Format(time.RFC3339), logLine)
	}
}
