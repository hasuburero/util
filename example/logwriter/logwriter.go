package main

import (
	"fmt"
	"github.com/hasuburero/util/logwriter"
	"github.com/hasuburero/util/times"
	"os/exec"
	"time"
)

func main() {
	writer, err := logwriter.MakeWriter("output.txt", []string{"timestamp", "1", "2", "3"})
	if err != nil {
		fmt.Println(err)
		fmt.Println("logwriter.MakeWriter error")
		return
	}

	for i := range 100 {
		ts := time.Now().Format(times.Layout1)
		buf := []string{}
		buf = append(buf, ts)
		buf = append(buf, fmt.Sprintf("%d", i*1))
		buf = append(buf, fmt.Sprintf("%d", i*2))
		buf = append(buf, fmt.Sprintf("%d", i*3))
		writer.Chan <- buf
	}

	output, err := exec.Command("cat", "output.txt").CombinedOutput()
	if err != nil {
		fmt.Println(err)
		fmt.Println("exec.Command error")
		return
	}
	fmt.Print(string(output))

	return
}
