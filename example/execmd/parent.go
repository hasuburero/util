package main

import (
	"fmt"
	"github.com/hasuburero/util/execmd"
	"os/exec"
)

var (
	cmd  = "go"
	args = []string{"run", "echo.go"}
)

func main() {
	cmd := exec.Command(cmd, args...)
	//stdpipe, err := execmd.InitPipe(cmd, 1024)
	var stdpipe *execmd.StdPipe = new(execmd.StdPipe)
	var err error
	stdpipe.StdinPipe, err = cmd.StdinPipe()
	//stdpipe.StdinPipe, err = cmd.StdinPipe()
	if err != nil {
		return
	}
	stdpipe.StdoutPipe, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	stdpipe.StderrPipe, err = cmd.StderrPipe()
	if err != nil {
		return
	}

	stdpipe.Output_buf = make([]byte, 1024)

	if err != nil {
		fmt.Println(err)
		return
	}

	//outpipe, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//var bytebuf = make([]byte, 1024)
	//n, err := outpipe.Read(bytebuf)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(bytebuf[:n])

	go func() {
		fmt.Println("Writing")
		err := stdpipe.Write("Hello World\n")
		if err != nil {
			fmt.Println(err)
			fmt.Println("failed to write")
			fmt.Println("Killing Process")
			cmd.Process.Kill()
			return
		}

		fmt.Println("Reading")
		n, err := stdpipe.StdoutPipe.Read(stdpipe.Output_buf)
		if err != nil {
			fmt.Println(err)
			fmt.Println("failed to read")
			fmt.Println("Killing Process")
			cmd.Process.Kill()
			return
		}
		fmt.Println(string(stdpipe.Output_buf[:n]))
		//n, err := stdpipe.Read()
		//if err != nil {
		//	fmt.Println(err)
		//	fmt.Println("failed to read")
		//	fmt.Println("Killing Process")
		//	cmd.Process.Kill()
		//	return
		//}
		//fmt.Println(string(stdpipe.Output_buf[:n]))

		cmd.Process.Kill()
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Pid: %d\n", cmd.Process.Pid)
	cmd.Wait()
	fmt.Printf("%s\n", cmd.ProcessState.String())
}
