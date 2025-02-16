package panic

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

var Err chan error
var addchan chan func()
var handler []func()
var running bool = false
var mux sync.Mutex

func Add(fn func()) error {
	mux.Lock()
	if running {
		mux.Unlock()
		return errors.New("panic handler is not running")
	}
	mux.Unlock()
	addchan <- fn
	return nil
}

func Error(err error) {
	mux.Lock()
	if running {
		mux.Unlock()
		fmt.Println("panic handler is not running")
		os.Exit(2)
	}
	mux.Unlock()

	Err <- err
}

func execHandle() {
	for _, ctx := range handler {
		ctx()
	}
}

func Start() error {
	mux.Lock()
	if running {
		mux.Unlock()
		return errors.New("panic handler is already running")
	}
	mux.Unlock()
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt)
	Err = make(chan error)
	addchan = make(chan func())
	go func() {
		for {
			select {
			case newfunc := <-addchan:
				handler = append(handler, newfunc)
			case err := <-Err:
				fmt.Println(err)
				execHandle()
				os.Exit(2)
			case sig := <-sigch:
				fmt.Println("signal received!! ", sig)
				execHandle()
				os.Exit(2)
			}
		}
	}()
	return nil
}
