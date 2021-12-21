package main

import (
	"bufio"
	"fmt"
	"os"
)

type myData struct {
	R                    *bufio.Reader
	L, N, jSpeed, bSpeed int
	restPoint            [][2]int
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	fmt.Fscanln(d.R, &d.L, &d.N, &d.jSpeed, &d.bSpeed)
	for i := 0; i < d.N; i++ {
		var x, c int
		fmt.Fscanln(d.R, &x, &c)
		temp := [2]int{x, c}
		d.restPoint = append(d.restPoint, temp)
	}
	return d
}

func getMax(d *myData, index int) int {
	max, pick := 0, 0
	for i := index; i < len(d.restPoint); i++ {
		if d.restPoint[i][1] > max {
			pick = i
			max = d.restPoint[i][1]
		}
	}
	return pick
}

func getMaxTastiness(d *myData) int {
	totalTaste, gap, pre := 0, 0, 0
	for i := 0; i < len(d.restPoint); i++ {
		index := getMax(d, i)
		gap = d.restPoint[index][0] - pre
		pre = d.restPoint[index][0]
		totalTaste += d.restPoint[index][1] * gap * (d.jSpeed - d.bSpeed)
		i = index
	}
	return totalTaste
}

func main() {
	d := inputAction()
	fmt.Println(getMaxTastiness(d))
}

/*
10 5 4 3
3 1
1 2
8 3
9 3
6 2

54+8+3 = 65

10 9 4 3
1 2
2 1
3 1
4 6
5 4
6 9
7 2
8 4
9 3

*/
