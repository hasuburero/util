package execmd

import (
	"bytes"
	"io"
	"os/exec"
)

type StdPipe struct {
	StdinPipe  io.WriteCloser
	StdoutPipe io.ReadCloser
	StderrPipe io.ReadCloser
	Output_buf []byte
}

func (self *StdPipe) Close() error {
	err := self.StdinPipe.Close()
	if err != nil {
		return err
	}
	err = self.StdoutPipe.Close()
	if err != nil {
		return err
	}
	err = self.StderrPipe.Close()
	if err != nil {
		return err
	}

	return nil
}

func (self *StdPipe) Read() (int, error) {
	n, err := self.StdoutPipe.Read(self.Output_buf)
	if err == io.EOF {
		return -1, err
	} else if err != nil {
		return -1, err
	}

	return n, nil
}

func (self *StdPipe) Write(arg string) error {
	_, err := io.Copy(self.StdinPipe, bytes.NewBuffer([]byte(arg)))
	return err
}

func InitPipe(cmd *exec.Cmd, bufsize int) (StdPipe, error) {
	var stdpipe StdPipe
	var err error
	stdpipe.StdinPipe, err = cmd.StdinPipe()
	if err != nil {
		return StdPipe{}, err
	}
	stdpipe.StdoutPipe, err = cmd.StdoutPipe()
	if err != nil {
		return StdPipe{}, err
	}
	stdpipe.StderrPipe, err = cmd.StderrPipe()
	if err != nil {
		return StdPipe{}, err
	}

	stdpipe.Output_buf = make([]byte, bufsize)

	return stdpipe, nil
}
