package panic

import ()

import (
	"github.com/hasuburero/util/log"
)

const (
	Layout = "2006-01-02 15:04:05.000000"
)

func PrintError(err error) {
	log.PrintLog(err.Error(), "error")
}
