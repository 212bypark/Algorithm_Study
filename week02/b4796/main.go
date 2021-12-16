package main

import (
	"bufio"
	"fmt"
	"os"
)

type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	testCases []testCase
}

type testCase struct {
	L int
	P int
	V int
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	for {
		var c testCase
		fmt.Fscanln(d.R, &c.L, &c.P, &c.V)
		if c.L == 0 && c.P == 0 && c.V == 0 {
			break
		}
		d.testCases = append(d.testCases, c)
	}
	return d
}

func solution() {
	d := inputAction()
	defer d.W.Flush()
	index := 1
	for _, value := range d.testCases {
		result := value.V / value.P * value.L
		if value.V%value.P > value.L {
			result += value.L
		} else {
			result += value.V % value.P
		}
		fmt.Printf("Case %d: %d\n", index, result)
		index++
	}
}

func main() {
	solution()
}
