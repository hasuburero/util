package main

import (
	"errors"
	"fmt"
	"github.com/hasuburero/util/panic"
	"sync"
	"time"
)

func test1() {
	fmt.Println("this is test1")
}
func test2() {
	fmt.Println("this is test2")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	err := panic.Start()
	if err != nil {
		fmt.Println(err)
		fmt.Println("panic.Start error")
		return
	}
	panic.Add(test1)
	panic.Add(test2)
	time.Sleep(5 * time.Second)
	panic.Error(errors.New("exitting with normal"))
	wg.Wait()
}
