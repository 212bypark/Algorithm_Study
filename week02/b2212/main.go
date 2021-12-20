package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type myData struct {
	R       *bufio.Reader
	divide  int
	sensors []int
}

func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	n, num, index := 0, 0, 0
	fmt.Fscan(d.R, &n, &d.divide)
	var temp []int
	for i := 0; i < n; i++ {
		fmt.Fscan(d.R, &num)
		temp = append(temp, num)
	}
	sort.Ints(temp)
	d.sensors = append(d.sensors, temp[0])
	for i := 1; i < n; i++ {
		if temp[i] != d.sensors[index] {
			d.sensors = append(d.sensors, temp[i])
			index++
		}
	}
	return d
}

func getMinSum(d *myData) int {
	sum := 0
	var disArr []int
	for i := 1; i < len(d.sensors); i++ {
		disArr = append(disArr, d.sensors[i]-d.sensors[i-1])
	}
	sort.Ints(disArr)
	for i := 0; i < len(disArr)-d.divide+1; i++ {
		sum += disArr[i]
	}
	return sum
}

func main() {
	d := inputAction()
	fmt.Println(getMinSum(d))
}

// 1 3 6 7 9
// 1 3 6 6 7 9 (2) => [1~3][6~9] : 2+3 = 5
//  2 3   1 2

// 3 6 7 8 10 12 14 15 18 20 (5) => [3] [6~10] [12] [14~15] [18~20] : 0+4+0+1+2 = 7
//  3 1 1 2  2  2  1  3  2
// [3][6~8][10][12~15][18~20] => 0+2+0+3+2 = 7
