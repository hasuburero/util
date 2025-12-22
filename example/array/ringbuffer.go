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
	for i := range 32 {
		rb.Push(Test{A: i, B: float32(i)})
		contents := rb.Get()
		fmt.Println(contents)
		fmt.Println(rb.Length)
	}

	return
}
