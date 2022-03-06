package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	wr      = bufio.NewWriter(os.Stdout)
)

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func scanInt() int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func binarySearch(max int, arr []int, limit int) int {
	lo, hi := 0, max+1
	for lo+1 < hi {
		mid := (lo + hi) / 2
		if isValid(mid, limit, arr) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

func isValid(param int, limit int, state []int) bool {
	sum := 0
	for _, tmp := range state {
		sum += getMin(param, tmp)
	}
	if sum <= limit {
		return true
	} else {
		return false
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	defer wr.Flush()

	n := scanInt()
	state := make([]int, n)
	for i := 0; i < n; i++ {
		state[i] = scanInt()
	}
	m := scanInt()
	max := 0
	for _, tmp := range state {
		if tmp > max {
			max = tmp
		}
	}
	cost := binarySearch(max, state, m)
	wr.Write([]byte(strconv.Itoa(cost)))
	wr.WriteByte('\n')
}
