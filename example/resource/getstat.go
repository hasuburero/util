package main

import (
	"fmt"
	"github.com/hasuburero/util/resource"
)

func main() {
	ts, cpustat, err := resource.GetCPUStat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cpustat)
	fmt.Printf("%v %v\n", ts, cpustat)

	ts, memstat, err := resource.GetMem()
	if err != nil {
		return
	}

	fmt.Println(memstat)
	fmt.Printf("%v, %v\n", ts, memstat)

	return
}
