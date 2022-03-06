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

func isValid(param int, limit int, arr []int) bool {
	for _, tmp := range arr {
		if tmp > param {
			return false
		} else {

		}
	}

}

func binarySearch(max int, limit int, arr []int) int {
	lo, hi := 1, max+1
	for lo < hi {
		mid := (lo + hi) / 2
		if isValid(mid, limit, arr) {
			lo = mid
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	defer wr.Flush()

	limit := 1000000000
	n, m := scanInt(), scanInt()
	plan := make([]int, n)
	for i := 0; i < n; i++ {
		plan[i] = scanInt()
	}
	// money := sort.Search(limit, func(i int) bool {
	// 	lo, hi := 1, limit+1
	// 	for _, tmp := range plan {
	// 		if tmp > i {
	// 			return false
	// 		}
	// 		if tmp > hi {
	// 			hi =
	// 		}
	// 	}
	// })

}
