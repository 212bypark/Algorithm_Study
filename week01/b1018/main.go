package main

import (
	"bufio"
	"fmt"
	"os"
)

type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	n     int
	m     int
	chess []string
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.n, &d.m)
	var nat = make([]string, d.n)
	for i := 0; i < d.n; i++ {
		fmt.Fscan(d.R, &nat[i])
	}
	d.chess = nat
	return d
}

func myMin(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

func count(d *myData) int {
	min := d.n * d.m
	for i := 0; i < d.n-7; i++ {
		for j := 0; j < d.m-7; j++ {
			count1, count2 := 0, 0
			for a := i; a < i+8; a++ {
				for b := j; b < j+8; b++ {
					if (a+b)%2 == 0 {
						if string(d.chess[a][b]) == "B" {
							count1++
						} else {
							count2++
						}
					} else {
						if string(d.chess[a][b]) == "B" {
							count2++
						} else {
							count1++
						}
					}
				}
			}
			if min > myMin(count1, count2) {
				min = myMin(count1, count2)
			}
		}
	}
	return min
}

func main() {
	d := inputAction()
	fmt.Println(count(d))
}
