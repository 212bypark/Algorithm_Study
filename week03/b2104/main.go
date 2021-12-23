package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func getInt() uint64 {
	scanner.Scan()
	var x uint64
	x = 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += uint64(b - '0')
	}
	return x
}

func getMin(arr []uint64) uint64 {
	min := arr[0]
	for _, val := range arr {
		if val < min {
			min = val
		}
	}
	return min
}

func getSum(arr []uint64) uint64 {
	var sum uint64
	sum = 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func maxScore() {
	size := getInt()
	scoreArr := make([]uint64, size)
	for i := range scoreArr {
		scoreArr[i] = getInt()
	}
	var max uint64
	max = 0
	for i := 0; i < len(scoreArr); i++ {
		for j := i + 1; j < len(scoreArr)+1; j++ {
			tmp := getSum(scoreArr[i:j]) * getMin(scoreArr[i:j])
			// fmt.Println(getSum(scoreArr[i:j]), getMin(scoreArr[i:j]))
			if max < tmp {
				max = tmp
			}
		}
	}
	fmt.Println(max)
}

func main() {
	maxScore()
}

// test case
/*
6
3 1 6 4 5 2
*/
