// Memory: 944KB
// Time : 212ms
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type myData struct {
	R *bufio.Reader

	count  int
	total  int
	values []int
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)

	fmt.Fscanln(d.R, &d.count, &d.total)
	for i := 0; i < d.count; i++ {
		var temp int
		fmt.Fscan(d.R, &temp)
		d.values = append(d.values, temp)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(d.values)))
	return d
}

func getMinCount() {
	d := inputAction()
	minCount := 0
	for _, cur := range d.values {
	CHECK:
		if d.total-cur >= 0 {
			d.total -= cur
			minCount++
			if d.total == 0 {
				break
			}
			goto CHECK
		}
	}
	fmt.Println(minCount)
}

func main() {
	getMinCount()
}
