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

func nextInt() int {
	scanner.Scan()
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

func cal(a, b, c int) int {
	if b == 1 {
		return a % c
	}
	if b%2 == 0 {
		temp := cal(a, b/2, c)
		return (temp * temp) % c
	}
	return (cal(a, b/2, c) * cal(a, b/2+1, c)) % c
}

func main() {
	a, b, c := nextInt(), nextInt(), nextInt()
	fmt.Println(cal(a, b, c))
}
