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

func getInt() int {
	scanner.Scan()
	x := 0
	for _, b := range scanner.Bytes() {
		x *= 10
		x += int(b - '0')
	}
	return x
}

// func getMin(arr []uint64) uint64 {
// 	min := arr[0]
// 	for _, val := range arr {
// 		if val < min {
// 			min = val
// 		}
// 	}
// 	return min
// }

// func getSum(arr []uint64) uint64 {
// 	var sum uint64
// 	sum = 0
// 	for _, val := range arr {
// 		sum += val
// 	}
// 	return sum
// }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxScore() {
	n := getInt()
	a, sum := make([]int, n+2), make([]int, n+2)
	for i := 1; i <= n; i++ {
		a[i] = getInt()
		sum[i] = sum[i-1] + a[i]
	}
	fmt.Println("a:", a, " , sum:", sum)
	type p struct {
		x, y int
	}
	st := make([]p, 0, n)
	ans := 0
	for i := 1; i <= n+1; i++ {
		fmt.Println("------------------------", i)
		if i != 1 {
			fmt.Println("===> ", len(st), ">0 && ", st[len(st)-1].x, " > ", a[i])
		}
		for len(st) > 0 && st[len(st)-1].x > a[i] {
			t := st[len(st)-1]
			st = st[:len(st)-1]
			pidx := 0
			if len(st) != 0 {
				pidx = st[len(st)-1].y
			}
			ans = max(ans, t.x*(sum[i-1]-sum[pidx]))
			fmt.Println("t:", t, ", st:", st, ", pidx:", pidx, ", ans:", ans)
		}
		st = append(st, p{a[i], i})
		fmt.Println("ST - ", st, ", ans -", ans)
	}
	fmt.Println(ans)
	// size := getInt()
	// scoreArr := make([]uint64, size)
	// for i := range scoreArr {
	// 	scoreArr[i] = getInt()
	// }
	// var max uint64
	// max = 0
	// for i := 0; i < len(scoreArr); i++ {
	// 	for j := i + 1; j < len(scoreArr)+1; j++ {
	// 		tmp := getSum(scoreArr[i:j]) * getMin(scoreArr[i:j])
	// 		// fmt.Println(getSum(scoreArr[i:j]), getMin(scoreArr[i:j]))
	// 		if max < tmp {
	// 			max = tmp
	// 		}
	// 	}
	// }
	// fmt.Println(max)
}

func main() {
	maxScore()
}

// test case
/*
6
3 1 6 4 5 2
*/
