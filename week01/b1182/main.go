package main

import (
	"bufio"
	"fmt"
	"os"
)

// myData 구조체에 표준입력을 통해 받은 n,s 값과 전체수열 저장
type myData struct {
	R *bufio.Reader
	W *bufio.Writer

	n             int
	s             int
	sequenceArray []int
}

// 표준입력을 통해 값을 받아옴
func inputAction() *myData {
	d := &myData{}
	d.R = bufio.NewReader(os.Stdin)
	d.W = bufio.NewWriter(os.Stdout)

	fmt.Fscanln(d.R, &d.n, &d.s)
	for i := 0; i < d.n; i++ {
		val := 0
		fmt.Fscan(d.R, &val)
		d.sequenceArray = append(d.sequenceArray, val)
	}
	return d
}

// 2^(num) 연산 수행
// math 패키지의 Pow(x, y float64) float64 함수로 대체 가능
func squaredPower(num int) int {
	val := 1
	for i := 0; i < num; i++ {
		val *= 2
	}
	return val
}

// 2^(전체수열크기)  만큼 부분수열을 만들 수 있는 경우의 수 존재
// 부분수열은 하나 이상의 값을 포함해야 함 -> for문 i=1부터 시작
// 2진법 로직을 이용하여 문제 해결
// (예)
// 5 0
// -7 -3 -2 5 8
// 부분수열을 만들 수 있는 경우의 수 2^5 - 1 (모두 포함하지 않는 경우 제외)
// 1 = 10000 -> [-7]
// 2 = 01000 -> [-3]
// 3 = 11000 -> [-7, -3]
// 4 = 00100 -> [-2]
// ...
// 29 = 10111 -> [-7, -2, 5, 8]
// 30 = 01111 -> [-3, -2, 5, 8]
// 31 = 11111 -> [-7, -3, -2, 5, 8]
// total변수에 부분수열 값들의 총합 저장
// count 변수에 S값을 만들 수 있는 부분수열의 경우의 수 저장
// Complexity - Time: O(2^n) Exponential | Space: O(1)
func checkAvail(d *myData) int {
	count := 0
	numCheck := squaredPower(d.n)
	for i := 1; i < numCheck; i++ {
		val := i
		total := 0
		for j := d.n - 1; j >= 0; j-- {
			if val/squaredPower(j) == 1 {
				total += d.sequenceArray[j]
				val -= squaredPower(j)
			}
		}
		if total == d.s {
			count++
		}
	}
	return count
}

func main() {
	d := inputAction()
	fmt.Println(checkAvail(d))
}
