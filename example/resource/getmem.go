package main

import (
	"fmt"
	"github.com/hasuburero/util/resource"
)

func main() {
	ts, mem, err := resource.GetMem()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s: %f%%\n", ts.String(), mem)

	return
}
