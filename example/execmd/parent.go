package main

import (
	"github.com/hasuburero/util/execmd"
	"os/exec"
)

var (
	cmd  = "go"
	args = []string{"run", "echo.go"}
)

func main() {
	cmd := exec.Command(cmd, args...)
	stdpipe, err := execmd.InitPipe(cmd)
	if err != nil {
		return
	}

	cmd.Start()

	go stdpipe.Read()
	go stdpipe.Write()

	cmd.Wait()

}
