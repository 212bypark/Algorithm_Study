// Memory: 2436KB
// Time : 56ms
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func nextInt() int {
	scanner.Scan()
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

type meeing struct {
	s, e int
}

type byEnd []meeing

func (a byEnd) Len() int      { return len(a) }
func (a byEnd) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byEnd) Less(i, j int) bool {
	if a[i].e == a[j].e {
		return a[i].s < a[j].s
	}
	return a[i].e < a[j].e
}

func getNumMaxMeet() {
	N := nextInt()
	meetings := make([]meeing, N)
	for i := range meetings {
		s := nextInt()
		e := nextInt()
		meetings[i] = meeing{s, e}
	}

	sort.Sort(byEnd(meetings))

	last := 0
	count := 0
	for _, m := range meetings {
		if m.s >= last {
			last = m.e
			count++
		}
	}
	fmt.Println(count)
}
