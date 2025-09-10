package log

import (
	"fmt"
	"strings"
	"time"
)

const (
	Layout = "2006-01-02 15:04:05.000000"
)

func PrintLog(msg, tag string) {
	now := time.Now()
	lf := ""
	if !strings.Contains(msg, "\n") {
		lf = "\n"
	}

	fmt.Printf("[%v][%s] %s%s", now.Format(Layout), tag, msg, lf)
}
