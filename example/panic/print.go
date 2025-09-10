package main

import (
	"errors"
)

import (
	"github.com/hasuburero/util/panic"
)

func main() {
	err := errors.New("test1 error has occured\n")
	panic.PrintError(err)
	err = errors.New("test2 error has occured")
	panic.PrintError(err)

	return
}
