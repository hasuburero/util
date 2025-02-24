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
}

const (
	statfile = "/proc/stat"
)

var resource_instance *Resource_Struct
var column []string = []string{"Timestamp", "Usage", "User", "Nice", "Sys", "Idle", "IOwait", "Irq", "Softirq", "Steal"}

func (self *Resource_Struct) ErrorHandler() {
	self.Logwriter.Fd.Close()
}

func (self *Resource_Struct) LoopThread() {
	go func() {
		inter := time.Duration(self.Interval * int(time.Millisecond))
		var content []string = make([]string, 10, 10)
		var ts time.Time
		var Usage float32
		var User, Nice, Sys, Idle, IOwait, Irq, Softirq, Steal int
		for {
			fd, err := os.Open(statfile)
			if err != nil {
				panic.Error(err)
				return
			}
			defer fd.Close()

			ts = time.Now()
			reader := bufio.NewReader(fd)
			input, err := reader.ReadString('\n')
			if err != nil {
				panic.Error(err)
				return
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

			User, err = strconv.Atoi(slice[1])
			Nice, err = strconv.Atoi(slice[2])
			Sys, err = strconv.Atoi(slice[3])
			Idle, err = strconv.Atoi(slice[4])
			IOwait, err = strconv.Atoi(slice[5])
			Irq, err = strconv.Atoi(slice[6])
			Softirq, err = strconv.Atoi(slice[7])
			Steal, err = strconv.Atoi(slice[8])
			if err != nil {
				panic.Error(err)
				return
			}
			total := User + Nice + Sys + Idle + IOwait + Irq + Softirq + Steal
			Usage = float32(total-Idle-IOwait) / float32(total)

			content[0] = ts.Format(times.Layout1)
			content[1] = strconv.FormatFloat(float64(Usage), 'f', 3, 32)
			content[2] = strconv.Itoa(User)
			content[3] = strconv.Itoa(Nice)
			content[4] = strconv.Itoa(Sys)
			content[5] = strconv.Itoa(Idle)
			content[6] = strconv.Itoa(IOwait)
			content[7] = strconv.Itoa(Irq)
			content[8] = strconv.Itoa(Softirq)
			content[9] = strconv.Itoa(Steal)
			self.Logwriter.Write(content)
			time.Sleep(inter)
		}
	}()
}

func Start(interval int, filename string) {
	loginstance, err := logwriter.MakeWriterWithOverride(filename, column)
	if err != nil {
		panic.Error(err)
	}
	resource_instance = new(Resource_Struct)
	resource_instance.Logwriter = loginstance
	resource_instance.Interval = interval
	panic.Add(resource_instance.ErrorHandler)

	resource_instance.LoopThread()
}
