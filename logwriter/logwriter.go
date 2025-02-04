package logwriter

import (
	"errors"
	"fmt"
	"os"
)

type LogWriter struct {
	Filename string
	Fd       *os.File
	Column   []string
	Chan     chan []string
}

func (self *LogWriter) WritingThread() {
	go func() {
		for {
			select {
			case values := <-self.Chan:
				length := len(values)
				if length != len(self.Column) {
					continue
				}
				row := ""
				for i := 0; ; i++ {
					if values[i] == "" {
						break
					}
					row += values[i]
					if i != length-1 {
						row += ","
					} else {
						row += "\n"
						break
					}
				}
				self.Fd.WriteString(row)
			}
		}
	}()
}

func MakeWriter(filename string, column []string) (*LogWriter, error) {
	if len(column) == 0 {
		return nil, errors.New("column has zero length")
	}
	col := ""
	for i, length := 0, len(column); ; i++ {
		if column[i] == "" {
			buf := fmt.Sprintf("empty column name, index:%d", i)
			return nil, errors.New(buf)
		}
		col += column[i]
		if i != length-1 {
			col += ","
		} else {
			col += "\n"
			break
		}
	}

	_, err := os.Stat(filename)
	if err == nil {
		return nil, errors.New("another file is exists")
	}
	fd, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	writer := &LogWriter{Filename: filename, Fd: fd, Column: column}

	_, err = writer.Fd.WriteString(col)
	if err != nil {
		return nil, err
	}

	writer.Chan = make(chan []string)
	writer.WritingThread()

	return writer, nil
}
