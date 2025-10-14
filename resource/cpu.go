package resource

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

type CPUStat struct {
	Prev    CPU
	Current CPU
}

type CPU struct {
	Usage   float32
	Ts      time.Time
	Total   int
	User    int
	Nice    int
	Sys     int
	Idle    int
	IOwait  int
	Irq     int
	Softirq int
	Steal   int
}

const (
// statfile    = "/proc/stat"
// meminfofile = "/proc/meminfo"
)

func ReadStat() (time.Time, []string, error) {
	fd, err := os.Open(statfile)
	if err != nil {
		return time.Time{}, nil, err
	}
	defer fd.Close()

	ts := time.Now()
	reader := bufio.NewReader(fd)
	input, err := reader.ReadString('\n')
	if err != nil {
		return time.Time{}, []string{}, err
	}

	slice := strings.Split(input, " ")
	for i := 0; i < len(slice); {
		ctx := slice[i]
		if ctx == "" {
			if i == len(slice)-1 {
				slice = slice[:i]
			} else {
				slice = append(slice[:i], slice[i+1:]...)
			}
			continue
		}
		i++
	}

	return ts, slice, nil
}

func calcTotal(cpu *CPU, slice []string, ts time.Time) error {
	var err error
	cpu.Ts = ts
	cpu.User, err = strconv.Atoi(slice[1])
	cpu.Nice, err = strconv.Atoi(slice[2])
	cpu.Sys, err = strconv.Atoi(slice[3])
	cpu.Idle, err = strconv.Atoi(slice[4])
	cpu.IOwait, err = strconv.Atoi(slice[5])
	cpu.Irq, err = strconv.Atoi(slice[6])
	cpu.Softirq, err = strconv.Atoi(slice[7])
	cpu.Steal, err = strconv.Atoi(slice[8])
	if err != nil {
		return err
	}

	cpu.Total = cpu.User + cpu.Nice + cpu.Sys + cpu.Idle + cpu.IOwait + cpu.Irq + cpu.Softirq + cpu.Steal

	return nil
}

func (self *CPUStat) GetCPU() error {
	ts, slice, err := ReadStat()
	if err != nil {
		return err
	}

	err = calcTotal(&self.Current, slice, ts)
	if err != nil {
		return err
	}

	if self.Current.Total == self.Prev.Total {
		self.Current.Usage = 0.0
	} else {
		self.Current.Usage = float32((self.Current.Total-self.Prev.Total)-(self.Current.Idle+self.Current.IOwait-self.Prev.Idle-self.Prev.IOwait)) / float32(self.Current.Total-self.Prev.Total) * 100
	}
	self.Prev = self.Current

	return nil
}

func Init() (CPUStat, error) {
	var cpu CPUStat
	ts, slice, err := ReadStat()
	if err != nil {
		return CPUStat{}, err
	}

	err = calcTotal(&cpu.Prev, slice, ts)

	return cpu, nil
}
