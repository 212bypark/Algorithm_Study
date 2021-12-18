package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type myData struct {
	R *bufio.Reader
}

type meeting struct {
	startTime uint32
	endTime   uint32
}

type room struct {
	endTime uint32
	count   int
}

// 회의가 일찍 끝나는 순서 (같은면 일찍 시작하는 순서)
func solution() {
	d := &myData{}

	d.R = bufio.NewReader(os.Stdin)
	num := 0
	fmt.Fscanln(d.R, &num)
	var meetingInfo []meeting
	for i := 0; i < num; i++ {
		var info meeting
		fmt.Fscanln(d.R, &info.startTime, &info.endTime)
		meetingInfo = append(meetingInfo, info)
	}
	sort.Slice(meetingInfo, func(i, j int) bool {
		if meetingInfo[i].endTime == meetingInfo[j].endTime {
			return meetingInfo[i].startTime < meetingInfo[j].startTime
		}
		return meetingInfo[i].endTime < meetingInfo[j].endTime
	})
	var data []room
	for i := 0; i < num; i++ {
		var sT, eT uint32
		sT = meetingInfo[i].startTime
		eT = meetingInfo[i].endTime
		var rm room
		if i == 0 {
			rm.endTime = eT
			rm.count = 1
			data = append(data, rm)
		} else {
			for j := 0; j < len(data); j++ {
				if data[j].endTime <= sT {
					data[j].endTime = eT
					data[j].count++
					break
				}
				if j == len(data)-1 {
					rm.endTime = eT
					rm.count = 1
					data = append(data, rm)
					break
				}
			}
		}
	}
	fmt.Println(data[0].count)
}

func main() {
	solution()
}
