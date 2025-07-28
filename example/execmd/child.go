package main

import (
	"fmt"
	"github.com/hasuburero/util/execmd"
)

func main() {
	input, err := execmd.Read()
	if err != nil {
		fmt.Println(string(input))
	}
}
