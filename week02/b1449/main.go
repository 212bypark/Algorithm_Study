// Memory: 928KB
// Time : 4ms
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	N                uint16
	L                uint16
	location         []uint16
	tapeLastLocation uint16
}

func inputAction() *myData {
	d := &myData{}

	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.N, &d.L)
	for i := 0; i < int(d.N); i++ {
		var val uint16
		fmt.Fscan(d.R, &val)
		d.location = append(d.location, val)
	}
	sort.Slice(d.location, func(i, j int) bool {
		return d.location[i] < d.location[j]
	})
	d.tapeLastLocation = 0
	return d
}

func solution() {
	d := inputAction()
	count := 0
	for _, val := range d.location {
		if d.tapeLastLocation < val {
			count++
			d.tapeLastLocation = val + d.L - 1
		}
	}
	fmt.Println(count)
}

func main() {
	solution()
}
