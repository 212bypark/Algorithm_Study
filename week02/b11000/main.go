package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func getInt() int {
	scanner.Scan()
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}
*/

type classInfo struct {
	S, T int
}

type byEnd []classInfo

func (a byEnd) Len() int      { return len(a) }
func (a byEnd) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// 시작시간 오름차순 정렬
func (a byEnd) Less(i, j int) bool {
	if a[i].S == a[j].S {
		return a[i].T < a[j].T
	}
	return a[i].S < a[j].S
}

func countNumClass() {
	// N := getInt()
	// classes := make([]classInfo, N)
	// for i := range classes {
	// 	s := getInt()
	// 	t := getInt()
	// 	classes[i] = classInfo{s, t}
	// }
	var num int
	R := bufio.NewReader(os.Stdin)
	fmt.Fscanln(R, &num)
	classes := make([]classInfo, num)
	for i := 0; i < num; i++ {
		var s, t int
		fmt.Fscanln(R, &s, &t)
		classes[i] = classInfo{s, t}
	}

	sort.Sort(byEnd(classes))

	var endTime []int
	endTime = append(endTime, classes[0].T)
	for i := 1; i < num; i++ {
		s, t := classes[i].S, classes[i].T
		for j := range endTime {
			if endTime[j] <= s {
				endTime[j] = t
				break
			}
			if j == len(endTime)-1 {
				endTime = append(endTime, t)
			}
		}
	}
	// fmt.Println(endTime)
	fmt.Println(len(endTime))
}

func main() {
	countNumClass()
}
