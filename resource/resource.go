package resource

import (
	"bufio"
	"github.com/hasuburero/util/logwriter"
	"github.com/hasuburero/util/panic"
	"github.com/hasuburero/util/times"
	"os"
	"strconv"
	"strings"
	"time"
)

type Resource_Struct struct {
	Interval  int
	Logwriter *logwriter.LogWriter
	Current   CPUTime
	Prev      CPUTime
}

type CPUTime struct {
	Usage   float32
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
	statfile    = "/proc/stat"
	meminfofile = "/proc/meminfo"
)

var cpu_resource_instance *Resource_Struct
var mem_resource_instance *Resource_Struct
var cpu_column []string = []string{"Timestamp", "Usage%", "User", "Nice", "Sys", "Idle", "IOwait", "Irq", "Softirq", "Steal"}
var mem_column []string = []string{"Timestamp", "Usage%", "Total", "Free", "Available", "Buffers", "Cached"}

func (self *Resource_Struct) ErrorHandler() {
	self.Logwriter.Fd.Close()
}

func GetCPUStat() (time.Time, []string, error) {
	fd, err := os.Open(statfile)
	if err != nil {
		panic.Error(err)
		return time.Time{}, []string{}, err
	}
	defer fd.Close()

	ts := time.Now()
	reader := bufio.NewReader(fd)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic.Error(err)
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

func (self *Resource_Struct) CPULoopThread() {
	go func() {
		ts, slice, err := GetCPUStat()
		if err != nil {
			panic.Error(err)
			return
		}
		self.Prev.User, err = strconv.Atoi(slice[1])
		self.Prev.Nice, err = strconv.Atoi(slice[2])
		self.Prev.Sys, err = strconv.Atoi(slice[3])
		self.Prev.Idle, err = strconv.Atoi(slice[4])
		self.Prev.IOwait, err = strconv.Atoi(slice[5])
		self.Prev.Irq, err = strconv.Atoi(slice[6])
		self.Prev.Softirq, err = strconv.Atoi(slice[7])
		self.Prev.Steal, err = strconv.Atoi(slice[8])
		if err != nil {
			panic.Error(err)
			return
		}
		self.Prev.Total = self.Prev.User + self.Prev.Nice + self.Prev.Sys + self.Prev.Idle + self.Prev.IOwait + self.Prev.Irq + self.Prev.Softirq + self.Prev.Steal

		inter := time.Duration(self.Interval * int(time.Millisecond))
		var content []string = make([]string, 10, 10)
		for {
			ts, slice, err = GetCPUStat()
			self.Current.User, err = strconv.Atoi(slice[1])
			self.Current.Nice, err = strconv.Atoi(slice[2])
			self.Current.Sys, err = strconv.Atoi(slice[3])
			self.Current.Idle, err = strconv.Atoi(slice[4])
			self.Current.IOwait, err = strconv.Atoi(slice[5])
			self.Current.Irq, err = strconv.Atoi(slice[6])
			self.Current.Softirq, err = strconv.Atoi(slice[7])
			self.Current.Steal, err = strconv.Atoi(slice[8])
			if err != nil {
				panic.Error(err)
				return
			}
			self.Current.Total = self.Current.User + self.Current.Nice + self.Current.Sys + self.Current.Idle + self.Current.IOwait + self.Current.Irq + self.Current.Softirq + self.Current.Steal
			if self.Current.Total-self.Prev.Total == 0 {
				self.Current.Usage = 0.0
			} else {
				self.Current.Usage = float32((self.Current.Total-self.Prev.Total)-(self.Current.Idle+self.Current.IOwait-self.Prev.Idle-self.Prev.IOwait)) / float32(self.Current.Total-self.Prev.Total) * 100
			}

			content[0] = ts.Format(times.Layout1)
			content[1] = strconv.FormatFloat(float64(self.Current.Usage), 'f', 4, 32)
			content[2] = strconv.Itoa(self.Current.User)
			content[3] = strconv.Itoa(self.Current.Nice)
			content[4] = strconv.Itoa(self.Current.Sys)
			content[5] = strconv.Itoa(self.Current.Idle)
			content[6] = strconv.Itoa(self.Current.IOwait)
			content[7] = strconv.Itoa(self.Current.Irq)
			content[8] = strconv.Itoa(self.Current.Softirq)
			content[9] = strconv.Itoa(self.Current.Steal)
			self.Logwriter.Write(content)
			self.Prev = self.Current
			time.Sleep(inter)
		}
	}()
}

func (self *Resource_Struct) MEMLoopThread() {
	go func() {
		inter := time.Duration(self.Interval * int(time.Millisecond))
		var content []string = make([]string, 7, 7)
		var ts time.Time
		var Usage float32
		var int_buf []int = make([]int, 5, 5)
		for {
			fd, err := os.Open(meminfofile)
			if err != nil {
				panic.Error(err)
				return
			}
			defer fd.Close()

			ts = time.Now()
			reader := bufio.NewReader(fd)
			for i := 0; i < 5; i++ {
				input, err := reader.ReadString('\n')
				if err != nil {
					panic.Error(err)
					return
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
					panic.Error(err)
					return
				}
			}
			Usage = float32(int_buf[0]-int_buf[1]-int_buf[3]-int_buf[4]) / float32(int_buf[0]) * 100
			content[0] = ts.Format(times.Layout1)
			content[1] = strconv.FormatFloat(float64(Usage), 'f', 3, 32)
			content[2] = strconv.Itoa(int_buf[0])
			content[3] = strconv.Itoa(int_buf[1])
			content[4] = strconv.Itoa(int_buf[2])
			content[5] = strconv.Itoa(int_buf[3])
			content[6] = strconv.Itoa(int_buf[4])
			self.Logwriter.Write(content)
			time.Sleep(inter)
		}
	}()
}

func StartCPU(interval int, filename string) {
	loginstance, err := logwriter.MakeWriterWithOverride(filename, cpu_column)
	if err != nil {
		panic.Error(err)
	}
	cpu_resource_instance = new(Resource_Struct)
	cpu_resource_instance.Logwriter = loginstance
	cpu_resource_instance.Interval = interval
	panic.Add(cpu_resource_instance.ErrorHandler)

	cpu_resource_instance.CPULoopThread()
}

func StartMEM(interval int, filename string) {
	loginstance, err := logwriter.MakeWriterWithOverride(filename, mem_column)
	if err != nil {
		panic.Error(err)
	}
	mem_resource_instance = new(Resource_Struct)
	mem_resource_instance.Logwriter = loginstance
	mem_resource_instance.Interval = interval
	panic.Add(cpu_resource_instance.ErrorHandler)

	mem_resource_instance.MEMLoopThread()
}
