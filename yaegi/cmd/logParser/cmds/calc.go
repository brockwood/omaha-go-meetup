package script

import (
	"fmt"
	"strconv"
	"time"
)

var (
	Command     = `calculate`
	Description = `calculates the number of logs where the value > 50%`
)

func Cmd(times []time.Time, logLines []string) {
	counter := 0
	for _, logLine := range logLines {
		value, err := strconv.Atoi(logLine)
		if err != nil {
			continue
		}
		if value >= 500 {
			counter++
		}
	}
	percent := float64(counter) / float64(len(logLines))
	fmt.Printf("Percentage of reports > 500: %f%%\n", percent)
}
