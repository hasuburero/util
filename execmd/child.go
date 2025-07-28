package execmd

import (
	"bufio"
	"io"
	"os"
)

func Read() ([]byte, error) {
	scanner := bufio.NewScanner(os.Stdin)
	available := scanner.Scan()
	if !available {
		return nil, io.EOF
	}
	return scanner.Bytes(), nil
}
