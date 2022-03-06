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

func binarySearch(lo int, hi int, arr []int, limit int) int {
	for lo+1 < hi {
		mid := (lo + hi) / 2
		if isValid(mid, limit, arr) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return hi
}

func isValid(param int, limit int, arr []int) bool {
	storage, num := 0, 1
	for _, tmp := range arr {
		if storage+tmp > param {
			storage = 0
			num++
			if num > limit {
				return false
			}
		}
		storage += tmp
	}
	if num <= limit {
		return true
	} else {
		return false
	}

}

func main() {
	defer wr.Flush()

	n, m := scanInt(), scanInt()
	arr := make([]int, n)
	max, sum := 0, 0
	for i := 0; i < n; i++ {
		arr[i] = scanInt()
		if arr[i] > max {
			max = arr[i]
		}
		sum += arr[i]
	}
	long := binarySearch(max-1, sum+1, arr, m)
	wr.WriteString(strconv.Itoa(long) + "\n")
}
