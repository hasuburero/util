package main

import (
	"github.com/hasuburero/util/panic"
	"github.com/hasuburero/util/resource"
	"time"
)

func main() {
	panic.Start()
	resource.StartCPU(1000, "cpuoutput.csv")
	resource.StartMEM(1000, "memoutput.csv")
	time.Sleep(time.Duration(10 * time.Second))
	return
}
