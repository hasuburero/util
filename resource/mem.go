// in this file, i'm using MemAvailable
package resource

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadMem() (time.Time, []int, error) {
	ts := time.Now()
	fd, err := os.Open(meminfofile)
	if err != nil {
		return time.Time{}, nil, err
	}
	defer fd.Close()

	var int_buf []int = make([]int, 3)
	reader := bufio.NewReader(fd)
	for i := 0; i < 3; i++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			return time.Time{}, nil, err
		}
		slice := strings.Split(input, " ")
		for j := 0; j < len(slice); {
			ctx := slice[j]
			if ctx == "" {
				if j == len(slice)-1 {
					slice = slice[:j]
				} else {
					slice = append(slice[:j], slice[j+1:]...)
				}
				continue
			}
			j++
		}

		int_buf[i], err = strconv.Atoi(slice[1])
		if err != nil {
			return time.Time{}, nil, err
		}
	}

	return ts, int_buf, nil
}

func GetMem() (time.Time, float32, error) {
	ts, int_buf, err := ReadMem()
	if err != nil {
		return ts, 0, err
	}

	total := float32(int_buf[0])
	available := float32(int_buf[2])
	return ts, (total - available) / total * 100, nil
}
