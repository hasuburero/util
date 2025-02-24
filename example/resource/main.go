package main

import (
	"github.com/hasuburero/util/panic"
	"github.com/hasuburero/util/resource"
	"time"
)

func main() {
	panic.Start()
	resource.Start(1000, "output.csv")
	time.Sleep(time.Duration(10 * time.Second))
	return
}
