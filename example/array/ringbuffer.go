package main

import (
	"fmt"
	"github.com/hasuburero/util/array"
)

type Test struct {
	A int
	B float32
}

func main() {
	rb := array.MakeRingBuffer(30)
	contents := rb.Get()
	fmt.Println(contents)
	for i := range 32 {
		terminated := rb.Push(Test{A: i, B: float32(i)})
		fmt.Println("terminated: ", terminated)
		contents = rb.Get()
		fmt.Println(contents)
		fmt.Println(rb.Length)
	}

	return
}
