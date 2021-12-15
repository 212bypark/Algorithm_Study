package main

import (
	"bufio"
	"fmt"
	"os"
)

type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	n             int
	s             int
	sequenceArray []int
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.n, &d.s)
	for i := 0; i < d.n; i++ {
		val := 0
		fmt.Fscan(d.R, &val)
		d.sequenceArray = append(d.sequenceArray, val)
	}
	return d
}

func squaredPower(num int) int {
	val := 1
	for i := 0; i < num; i++ {
		val *= 2
	}
	return val
}

func checkAvail(d *myData) int {
	count := 0
	numCheck := squaredPower(d.n)
	for i := 1; i < numCheck; i++ {
		val := i
		total := 0
		for j := d.n - 1; j >= 0; j-- {
			if val/squaredPower(j) == 1 {
				total += d.sequenceArray[j]
				val -= squaredPower(j)
			}
		}
		if total == d.s {
			count++
		}
	}
	return count
}

func main() {
	d := inputAction()
	fmt.Println(checkAvail(d))
}
