package main

import (
	"fmt"
	"github.com/hasuburero/util/resource"
	"time"
)

func main() {
	cpu, err := resource.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(1 * time.Second)

	for _ = range 10 {
		err = cpu.GetCPU()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("%s: %f%%\n", cpu.Current.Ts.String(), cpu.Current.Usage)
		//fmt.Printf("%f\n", cpu.Current.Usage)
		time.Sleep(1 * time.Second)
	}
}
