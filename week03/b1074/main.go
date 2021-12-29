package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func scanInt() int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func findOrder(i, j int) int {
	if i < 2 && j < 2 {
		return 2*i + j
	}
	e := 2
	for {
		if i/e == 0 && j/e == 0 {
			e = e / 2
			break
		} else {
			e = e * 2
		}
	}
	return (2*e*e*(i/e) + e*e*(j/e) + findOrder(i%e, j%e))
}

func main() {
	scanInt()
	r, c := scanInt(), scanInt()
	fmt.Println(findOrder(r, c))
}

// 00 01 10 11
// 02 03 12 13
// 20 21 30 31
// 22 32 32 33
