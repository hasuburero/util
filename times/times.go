package times

import (
  "fmt"
  "strconv"
  "strings"
  "time"
)

type TimeStamp struct{
  Timer time.Time
  Timestamp string
  Hour int
  Min int
  Sec int
  Millis int
}

func (self *TimeStamp)getTime(){
  self.Timer = time.Now()
  self.Timestamp = self.Timer.Format(time.StampMilli)
  buf := strings.
}
