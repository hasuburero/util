package panic

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type Handler struct {
	AddChan  chan chan os.Signal
	Channels []chan os.Signal
}

func (self *Handler) Add(ch chan os.Signal) {
	self.AddChan <- ch
}

func Init() Handler {
	var handler_instance Handler
	addchan := make(chan chan os.Signal)
	channels := make([]chan os.Signal, 0)
	handler_instance = Handler{addchan, channels}
	return handler_instance
}

func (self *Handler) Start() {
	go func() {
		select {
		case newChannel := <-self.AddChan:
			self.Channels = append(self.Channels, newChannel)
		default:
			for _, ctx := range self.Channels {
				select {
				case sig := <-ctx:
					fmt.Println("signal received")
					fmt.Println("sig: %d\n", sig)
					os.Exit(1)
				default:
				}
			}
		}
	}()
}
