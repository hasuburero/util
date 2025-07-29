package main

import (
	"fmt"
	"github.com/hasuburero/util/execmd"
)

func main() {
	for {
		input, err := execmd.Read()
		if err != nil {
			fmt.Errorf("%v", err)
		}
		fmt.Println(string(input))
	}
}
