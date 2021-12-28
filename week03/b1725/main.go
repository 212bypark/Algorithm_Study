package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func getMin(arr []int) int {
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(tmp)
	// fmt.Println("checking:", arr)
	// fmt.Println("checking:", tmp)
	return tmp[0]
}

func getMaxVal(baselen int, arr []int) int {
	val := 0
	for i := 0; i+baselen <= len(arr); i++ {
		tmp := getMin(arr[i:i+baselen]) * baselen
		// fmt.Println("---> baselen=", baselen, ", i=", i, ", tmp=", tmp)
		if tmp > val {
			val = tmp
		}
	}
	return val
}

func getMaxSize() {
	n, val := scanInt(), 0
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = scanInt()
	}
	for baselen := 1; baselen <= n; baselen++ {
		tmp := getMaxVal(baselen, arr)
		if tmp > val {
			val = tmp
		}
		// fmt.Println("(baselen, val)=", baselen, ", ", val)
	}
	fmt.Println(val)
}

func main() {
	getMaxSize()
}
