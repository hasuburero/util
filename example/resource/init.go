package main

import (
	"fmt"
	"github.com/hasuburero/util/resource"
	"time"
)

func main() {
	rs, err := resource.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("mem usage: %f, mem total:%d, mem available: %d\n", float32(rs.MEMStat.Available)/float32(rs.MEMStat.Total), rs.MEMStat.Total, rs.MEMStat.Available)
	fmt.Printf("cpu prev sys: %d, cpu prev sys: %d, cpucpu prev usage: %f, cpu current usage: %f\n", rs.Prev.Sys, rs.Current.Sys, rs.Prev.Usage, rs.Current.Usage)

	time.Sleep(1000 * time.Millisecond)
	err = rs.NewCPUStat()
	err = rs.NewMEMStat()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("mem usage: %f, mem total:%d, mem available: %d\n", float32(rs.MEMStat.Available)/float32(rs.MEMStat.Total), rs.MEMStat.Total, rs.MEMStat.Available)
	fmt.Printf("cpu prev sys: %d, cpu prev sys: %d, cpucpu prev usage: %f, cpu current usage: %f\n", rs.Prev.Sys, rs.Current.Sys, rs.Prev.Usage, rs.Current.Usage)

	fmt.Printf("Usage: %f\n", rs.Current.Usage)

	return
}
