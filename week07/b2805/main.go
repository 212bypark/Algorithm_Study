package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	scanner *bufio.Scanner
	wr      = bufio.NewWriter(os.Stdout)
	height  []int
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

func main() {
	defer wr.Flush()

	n, m := scanInt(), scanInt()
	for i := 0; i < n; i++ {
		tmp := scanInt()
		height = append(height, tmp)
	}
	lo, hi := 0, 1000000000
	for lo+1 < hi {
		mid := (lo + hi) / 2
		sum := 0
		for _, tree := range height {
			if tree > mid {
				sum += tree - mid
			}
		}
		if sum >= m {
			lo = mid
		} else {
			hi = mid
		}
	}
	wr.Write([]byte(strconv.Itoa(lo)))
	wr.WriteByte('\n')
}
