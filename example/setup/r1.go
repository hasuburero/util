package main

import (
	"fmt"
	"os"
)

import "github.com/hasuburero/util/setup"

func main() {
	args := setup.ParseArgs(os.Args)
	value, err := args.GetArgs("worker_id")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(value)
	return
}
