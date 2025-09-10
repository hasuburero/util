package panic

import (
	"fmt"
	"strings"
	"time"
)

const (
	Layout = "2006-01-02 15:04:05.000000"
)

func PrintError(err error) {
	now := time.Now()
	lf := ""
	if !strings.Contains(err.Error(), "\n") {
		lf = "\n"
	}

	fmt.Printf("[%v][error] %s%s", now.Format(Layout), err.Error(), lf)

	return
}
